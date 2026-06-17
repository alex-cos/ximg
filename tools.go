package ximg

import (
	"math"
	"math/bits"
)

func clampU8(v int) uint8 {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return uint8(v)
}

func rgb2hsl(r, g, b uint8) (float64, float64, float64) {
	var s, h float64

	rf := float64(r) / 255
	gf := float64(g) / 255
	bf := float64(b) / 255

	mx := max(rf, max(gf, bf))
	mn := min(rf, min(gf, bf))
	delta := mx - mn

	l := (mx + mn) / 2

	if delta == 0 {
		return 0, 0, l
	}

	if l > 0.5 {
		s = delta / (2 - mx - mn)
	} else {
		s = delta / (mx + mn)
	}

	switch mx {
	case rf:
		h = (gf - bf) / delta
		if gf < bf {
			h += 6
		}
	case gf:
		h = (bf-rf)/delta + 2
	case bf:
		h = (rf-gf)/delta + 4
	}
	h /= 6

	return h, s, l
}

func hsl2rgb(h, s, l float64) (uint8, uint8, uint8) {
	var q float64

	if s == 0 {
		v := uint8(l * 255)
		return v, v, v
	}

	if l < 0.5 {
		q = l * (1 + s)
	} else {
		q = l + s - l*s
	}
	p := 2*l - q
	h = math.Mod(h, 1)
	if h < 0 {
		h += 1
	}

	toRGB := func(t float64) float64 {
		t = math.Mod(t, 1)
		if t < 0 {
			t += 1
		}
		if t < 1.0/6 {
			return p + (q-p)*6*t
		}
		if t < 0.5 {
			return q
		}
		if t < 2.0/3 {
			return p + (q-p)*6*(2.0/3-t)
		}
		return p
	}

	return uint8(toRGB(h+1.0/3)*255 + 0.5),
		uint8(toRGB(h)*255 + 0.5),
		uint8(toRGB(h-1.0/3)*255 + 0.5)
}

func nextPow2[I Integer](n I) I {
	if n <= 0 {
		return 1
	}
	if n&(n-1) == 0 {
		return n
	}
	return 1 << bits.Len(uint(n-1))
}

func bitReverse[C Complex](a []C) []C {
	n := len(a)
	logN := bits.Len(uint(n)) - 1
	for i := range a {
		j := int(bits.Reverse(uint(i)) >> (bits.UintSize - logN)) // nolint: gosec
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	return a
}

func minMaxChannel(w, h int, pixelAt func(int, int) uint8) (uint8, uint8) {
	mn := uint8(255)
	mx := uint8(0)
	for x := range w {
		for y := range h {
			v := pixelAt(x, y)
			if v < mn {
				mn = v
			}
			if v > mx {
				mx = v
			}
		}
	}
	return mn, mx
}

func rescaleU8(v, srcMin, srcMax, dstMin, dstMax uint8) uint8 {
	rng := srcMax - srcMin
	if rng == 0 {
		rng = 1
	}
	return uint8(float64(v-srcMin)*float64(dstMax-dstMin)/float64(rng) + float64(dstMin))
}

func complexFromPolar(angle float64) complex64 {
	sin, cos := math.Sincos(angle)
	return complex(float32(cos), float32(sin))
}

func lerp(a, b, t float64) float64 {
	return a*(1-t) + b*t
}

func bilinearSample(img *Ximg, sx, sy float64) (uint8, uint8, uint8, uint8) {
	w, h := img.Size()
	x0 := int(math.Floor(sx))
	y0 := int(math.Floor(sy))
	x1 := x0 + 1
	y1 := y0 + 1

	if x0 < 0 {
		x0 = 0
	}
	if x1 >= w {
		x1 = w - 1
	}
	if y0 < 0 {
		y0 = 0
	}
	if y1 >= h {
		y1 = h - 1
	}

	fx := sx - float64(x0)
	fy := sy - float64(y0)

	r00, g00, b00, a00 := img.RGBAAt(x0, y0)
	r10, g10, b10, a10 := img.RGBAAt(x1, y0)
	r01, g01, b01, a01 := img.RGBAAt(x0, y1)
	r11, g11, b11, a11 := img.RGBAAt(x1, y1)

	r := uint8(lerp(lerp(float64(r00), float64(r10), fx), lerp(float64(r01), float64(r11), fx), fy))
	g := uint8(lerp(lerp(float64(g00), float64(g10), fx), lerp(float64(g01), float64(g11), fx), fy))
	b := uint8(lerp(lerp(float64(b00), float64(b10), fx), lerp(float64(b01), float64(b11), fx), fy))
	a := uint8(lerp(lerp(float64(a00), float64(a10), fx), lerp(float64(a01), float64(a11), fx), fy))

	return r, g, b, a
}
