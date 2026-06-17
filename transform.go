package ximg

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/image/draw"
)

// Resize resizes the image. Set width or height to 0 to preserve aspect ratio.
func (img *Ximg) Resize(width, height int) *Ximg {
	w, h := img.Size()
	if width == 0 {
		width = w * height / h
	} else if height == 0 {
		height = h * width / w
	}
	ximg := NewRGBA(width, height)

	draw.CatmullRom.Scale(
		ximg,
		ximg.Bounds(),
		&img.RGBA,
		img.Bounds(),
		draw.Src,
		nil,
	)

	return ximg
}

// Split extracts square sub-images of the given width with the given stride.
func (img *Ximg) Split(width, shift int) []*Ximg {
	images := []*Ximg{}
	w, h := img.Size()
	if ((w == width) && (h == shift)) || ((w == shift) && (h == width)) {
		images = append(images, img)
		return images
	}
	for x := 0; x <= (w - width); x += shift {
		for y := 0; y <= (h - width); y += shift {
			r := image.Rectangle{
				Min: image.Point{X: x, Y: y},
				Max: image.Point{X: x + width, Y: y + width},
			}
			ximg, _ := New(img.SubImage(r))
			ximg.isGray = img.isGray
			images = append(images, ximg)
		}
	}

	return images
}

// Red extracts the red channel as a grayscale image.
func (img *Ximg) Red() *Ximg {
	f := func(src *Ximg, dst *Ximg, x int, y int) {
		r, _, _ := src.RGBAt(x, y)
		dst.Set(x, y, color.RGBA{r, r, r, 255})
	}
	ximg := img.Apply(1, f, nil)
	ximg.isGray = true

	return ximg
}

// Green extracts the green channel as a grayscale image.
func (img *Ximg) Green() *Ximg {
	f := func(src *Ximg, dst *Ximg, x int, y int) {
		_, g, _ := src.RGBAt(x, y)
		dst.Set(x, y, color.RGBA{g, g, g, 255})
	}
	ximg := img.Apply(1, f, nil)
	ximg.isGray = true

	return ximg
}

// Blue extracts the blue channel as a grayscale image.
func (img *Ximg) Blue() *Ximg {
	f := func(src *Ximg, dst *Ximg, x int, y int) {
		_, _, b := src.RGBAt(x, y)
		dst.Set(x, y, color.RGBA{b, b, b, 255})
	}
	ximg := img.Apply(1, f, nil)
	ximg.isGray = true

	return ximg
}

// Grey converts to grayscale using equal weights (R+G+B)/3.
func (img *Ximg) Grey() *Ximg {
	f := func(src *Ximg, dst *Ximg, x int, y int) {
		r, g, b, a := src.RGBAAt(x, y)
		grey := uint8((float64(r) + float64(g) + float64(b)) / 3.0)
		dst.Set(x, y, color.RGBA{grey, grey, grey, a})
	}
	ximg := img.Apply(1, f, nil)
	ximg.isGray = true

	return ximg
}

// Gray converts to grayscale using ITU-R BT.601 luminance weights.
func (img *Ximg) Gray() *Ximg {
	f := func(src *Ximg, dst *Ximg, x int, y int) {
		r, g, b, a := src.RGBAAt(x, y)
		grey := uint8(0.2989*float64(r) + 0.5870*float64(g) + 0.1140*float64(b))
		dst.Set(x, y, color.RGBA{grey, grey, grey, a})
	}
	ximg := img.Apply(1, f, nil)
	ximg.isGray = true

	return ximg
}

// Invert inverts all color channels while preserving alpha.
func (img *Ximg) Invert() *Ximg {
	f := func(src *Ximg, dst *Ximg, x int, y int) {
		r, g, b, a := src.RGBAAt(x, y)
		color := color.RGBA{
			255 - r,
			255 - g,
			255 - b,
			a,
		}
		dst.Set(x, y, color)
	}
	ximg := img.Apply(1, f, nil)
	ximg.isGray = img.isGray

	return ximg
}

// FlipH flips the image horizontally.
func (img *Ximg) FlipH() *Ximg {
	w, _ := img.Size()
	f := func(src *Ximg, dst *Ximg, x int, y int) {
		r, g, b, a := src.RGBAAt((w-1)-x, y)
		color := color.RGBA{r, g, b, a}
		dst.Set(x, y, color)
	}
	ximg := img.Apply(1, f, nil)
	ximg.isGray = img.isGray

	return ximg
}

// FlipV flips the image vertically.
func (img *Ximg) FlipV() *Ximg {
	_, h := img.Size()
	f := func(src *Ximg, dst *Ximg, x int, y int) {
		r, g, b, a := src.RGBAAt(x, (h-1)-y)
		color := color.RGBA{r, g, b, a}
		dst.Set(x, y, color)
	}
	ximg := img.Apply(1, f, nil)
	ximg.isGray = img.isGray

	return ximg
}

// MaxPool applies max pooling with the given increment.
func (img *Ximg) MaxPool(incr int) *Ximg {
	w, h := img.Size()
	f := func(src *Ximg, dst *Ximg, x int, y int) {
		fr := float64(0.0)
		fg := float64(0.0)
		fb := float64(0.0)
		for i := range incr {
			for j := range incr {
				k := x + i
				l := y + j
				if k >= 0 && k < w && l >= 0 && l < h {
					r, g, b := src.RGBAt(k, l)
					fr = math.Max(fr, float64(r))
					fg = math.Max(fg, float64(g))
					fb = math.Max(fb, float64(b))
				}
			}
		}
		dst.Set(x/incr, y/incr, color.RGBA{uint8(fr), uint8(fg), uint8(fb), 255})
	}
	ximg := img.Apply(incr, f, nil)
	ximg.isGray = img.isGray

	return ximg
}

// Normalize stretches the histogram to [minimum, maximum].
func (img *Ximg) Normalize(minimum, maximum uint8) *Ximg {
	w, h := img.Size()
	ximg := NewRGBA(w, h)

	rMin, rMax := minMaxChannel(w, h, func(x, y int) uint8 { r, _, _ := img.RGBAt(x, y); return r })
	gMin, gMax := minMaxChannel(w, h, func(x, y int) uint8 { _, g, _ := img.RGBAt(x, y); return g })
	bMin, bMax := minMaxChannel(w, h, func(x, y int) uint8 { _, _, b := img.RGBAt(x, y); return b })

	for x := range w {
		for y := range h {
			r, g, b, a := img.RGBAAt(x, y)
			ximg.Set(x, y, color.RGBA{
				rescaleU8(r, rMin, rMax, minimum, maximum),
				rescaleU8(g, gMin, gMax, minimum, maximum),
				rescaleU8(b, bMin, bMax, minimum, maximum),
				a,
			})
		}
	}
	ximg.isGray = img.isGray

	return ximg
}

// Merge blends two images with a factor in [0, 1] applied to `with`.
func (img *Ximg) Merge(with *Ximg, factor float64) *Ximg {
	if factor < 0 {
		factor = 0
	} else if factor > 1 {
		factor = 1
	}
	w, h := img.Size()
	ximg := NewRGBA(w, h)

	for x := range w {
		for y := range h {
			r1, g1, b1, a1 := img.RGBAAt(x, y)
			r2, g2, b2, a2 := with.RGBAAt(x, y)
			color := color.RGBA{
				uint8(float64(r1)*(1-factor) + float64(r2)*factor),
				uint8(float64(g1)*(1-factor) + float64(g2)*factor),
				uint8(float64(b1)*(1-factor) + float64(b2)*factor),
				uint8(float64(a1)*(1-factor) + float64(a2)*factor),
			}
			ximg.Set(x, y, color)
		}
	}
	ximg.isGray = img.isGray && with.isGray

	return ximg
}

func (img *Ximg) Fuzion(x, y int, with *Ximg) *Ximg {
	ximg := img.Clone()
	width, height := with.Size()
	r := image.Rect(x, y, x+width, y+height)
	draw.Draw(ximg, r, with, image.Point{X: 0, Y: 0}, draw.Over)
	ximg.isGray = img.IsGray() && with.isGray

	return ximg
}

// Contrast adjusts the image contrast using a sigmoid activation function.
func (img *Ximg) Contrast(a float64) *Ximg {
	activation := func(x float64) float64 {
		return float64(256.0) / (1.0 + math.Exp(-a*(x-128.0)))
	}
	f := func(src *Ximg, dst *Ximg, x int, y int) {
		r, g, b, a := src.RGBAAt(x, y)
		fr := activation(float64(r))
		fg := activation(float64(g))
		fb := activation(float64(b))
		fr = math.Min(255.0, math.Max(0.0, fr))
		fg = math.Min(255.0, math.Max(0.0, fg))
		fb = math.Min(255.0, math.Max(0.0, fb))
		dst.Set(x, y, color.RGBA{uint8(fr), uint8(fg), uint8(fb), a})
	}
	ximg := img.Apply(1, f, nil)
	ximg.isGray = img.isGray

	return ximg
}

// Brightness adds a delta to each color channel.
func (img *Ximg) Brightness(delta int) *Ximg {
	f := func(src *Ximg, dst *Ximg, x int, y int) {
		r, g, b, a := src.RGBAAt(x, y)
		dst.Set(x, y, color.RGBA{
			clampU8(int(r) + delta),
			clampU8(int(g) + delta),
			clampU8(int(b) + delta),
			a,
		})
	}
	ximg := img.Apply(1, f, nil)
	ximg.isGray = img.isGray

	return ximg
}

// Hue shifts the hue of each pixel by the given angle in degrees.
func (img *Ximg) Hue(shift float64) *Ximg {
	f := func(src *Ximg, dst *Ximg, x int, y int) {
		r, g, b, a := src.RGBAAt(x, y)
		h, s, l := rgb2hsl(r, g, b)
		h = math.Mod(h+shift/360, 1)
		if h < 0 {
			h += 1
		}
		nr, ng, nb := hsl2rgb(h, s, l)
		dst.Set(x, y, color.RGBA{nr, ng, nb, a})
	}
	ximg := img.Apply(1, f, nil)

	return ximg
}

// Rotate rotates the image by angle in degrees (counterclockwise) with transparent background.
func (img *Ximg) Rotate(angle float64) *Ximg {
	w, h := img.Size()
	cx := float64(w-1) / 2
	cy := float64(h-1) / 2

	rad := angle * math.Pi / 180
	cos := math.Cos(rad)
	sin := math.Sin(rad)

	nw := int(math.Ceil(math.Abs(cos)*float64(w) + math.Abs(sin)*float64(h)))
	nh := int(math.Ceil(math.Abs(sin)*float64(w) + math.Abs(cos)*float64(h)))
	ncx := float64(nw-1) / 2
	ncy := float64(nh-1) / 2

	dst := NewRGBA(nw, nh)

	for y := range nh {
		for x := range nw {
			dx := float64(x) - ncx
			dy := float64(y) - ncy
			sx := cos*dx + sin*dy + cx
			sy := -sin*dx + cos*dy + cy

			if sx >= 0 && sx < float64(w) && sy >= 0 && sy < float64(h) {
				r, g, b, a := bilinearSample(img, sx, sy)
				dst.Set(x, y, color.RGBA{r, g, b, a})
			}
		}
	}
	dst.isGray = img.isGray

	return dst
}

// ColorToAlpha sets pixels matching a target color to transparent within a tolerance in [0, 1].
// Tolerance 0 matches the exact color; 1 matches everything.
func (img *Ximg) ColorToAlpha(col color.RGBA, tolerance float32) *Ximg {
	if tolerance < 0 {
		tolerance = 0
	} else if tolerance > 1 {
		tolerance = 1
	}
	f := func(src *Ximg, dst *Ximg, x int, y int) {
		r, g, b, a := src.RGBAAt(x, y)
		// Chebyshev distance
		dr := max(r, col.R) - min(r, col.R)
		dg := max(g, col.G) - min(g, col.G)
		db := max(b, col.B) - min(b, col.B)
		maxDist := float32(max(dr, max(dg, db))) / 255
		if maxDist <= tolerance {
			a = 0
		}
		dst.Set(x, y, color.RGBA{r, g, b, a})
	}
	ximg := img.Apply(1, f, nil)

	return ximg
}

// AlphaToColor replaces transparency with a solid background color using alpha blending.
func (img *Ximg) AlphaToColor(bg color.RGBA) *Ximg {
	f := func(src *Ximg, dst *Ximg, x int, y int) {
		r, g, b, a := src.RGBAAt(x, y)
		af := float64(a) / 255
		nr := uint8(float64(bg.R)*(1-af) + float64(r)*af)
		ng := uint8(float64(bg.G)*(1-af) + float64(g)*af)
		nb := uint8(float64(bg.B)*(1-af) + float64(b)*af)
		dst.Set(x, y, color.RGBA{nr, ng, nb, 255})
	}
	ximg := img.Apply(1, f, nil)

	return ximg
}

// RemoveAlpha removes the alpha channel.
func (img *Ximg) RemoveAlpha() *Ximg {
	f := func(src *Ximg, dst *Ximg, x int, y int) {
		r, g, b := src.RGBAt(x, y)
		dst.Set(x, y, color.RGBA{r, g, b, 255})
	}
	ximg := img.Apply(1, f, nil)

	return ximg
}

// ConcatV concatenates images vertically, resizing them to the given width.
func ConcatV(width int, images ...*Ximg) *Ximg {
	if len(images) == 1 {
		return images[0]
	}
	if width == 0 {
		width, _ = images[0].Size()
	}
	height := 0
	for _, img := range images {
		w, h := img.Size()
		height += h * width / w
	}
	ximg := NewRGBA(width, height)

	yOff := 0
	for _, img := range images {
		iw, ih := img.Size()
		if iw != width {
			img = img.Resize(width, 0)
			_, ih = img.Size()
		}
		for x := range width {
			for y := range ih {
				r, g, b, a := img.RGBAAt(x, y)
				ximg.Set(x, y+yOff, color.RGBA{r, g, b, a})
			}
		}
		yOff += ih
	}

	return ximg
}

// ConcatH concatenates images horizontally, resizing them to the given height.
func ConcatH(height int, images ...*Ximg) *Ximg {
	if len(images) == 1 {
		return images[0]
	}
	if height == 0 {
		height, _ = images[0].Size()
	}
	width := 0
	for _, img := range images {
		w, h := img.Size()
		width += w * height / h
	}
	ximg := NewRGBA(width, height)

	xOff := 0
	for _, img := range images {
		iw, ih := img.Size()
		if ih != height {
			img = img.Resize(0, height)
			iw, _ = img.Size()
		}
		for x := range iw {
			for y := range height {
				r, g, b, a := img.RGBAAt(x, y)
				ximg.Set(x+xOff, y, color.RGBA{r, g, b, a})
			}
		}
		xOff += iw
	}

	return ximg
}
