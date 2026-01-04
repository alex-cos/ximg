package ximg

import (
	"image/color"
	"math"
)

// FromFloat creates a grayscale image from a flat []F in [0, 1].
// Values are clamped to [0, 1] before conversion.
func FromFloat[F Float](data []F, width, height int) *Ximg {
	ximg := NewRGBA(width, height)
	for x := range width {
		for y := range height {
			v := uint8(F(255)*data[x+width*y] + F(0.5))
			ximg.Set(x, y, color.RGBA{v, v, v, 255})
		}
	}
	ximg.isGray = true
	return ximg
}

// FromFloat2D creates a grayscale image from a 2D [][]F in [0, 1] (width × height).
func FromFloat2D[F Float](data [][]F) *Ximg {
	if len(data) == 0 {
		return NewRGBA(0, 0)
	}
	width := len(data)
	height := len(data[0])
	ximg := NewRGBA(width, height)
	for x := range width {
		for y := range height {
			v := uint8(F(255)*data[x][y] + F(0.5))
			ximg.Set(x, y, color.RGBA{v, v, v, 255})
		}
	}
	ximg.isGray = true
	return ximg
}

// FromComplex creates a grayscale image from a flat []complex64 using magnitude.
// Input should be normalized to [0, 1] first.
func FromComplex(data []complex64, width, height int) *Ximg {
	ximg := NewRGBA(width, height)
	for x := range width {
		for y := range height {
			z := data[x+width*y]
			mag := math.Sqrt(float64(real(z)*real(z) + imag(z)*imag(z)))
			mag = max(0, min(mag, 1))
			v := uint8(mag*255 + 0.5)
			ximg.Set(x, y, color.RGBA{v, v, v, 255})
		}
	}
	ximg.isGray = true
	return ximg
}

// FromComplex2D creates a grayscale image from a 2D [][]complex64 using magnitude.
func FromComplex2D(data [][]complex64) *Ximg {
	if len(data) == 0 {
		return NewRGBA(0, 0)
	}
	width := len(data)
	height := len(data[0])
	ximg := NewRGBA(width, height)
	for x := range width {
		for y := range height {
			z := data[x][y]
			mag := math.Sqrt(float64(real(z)*real(z) + imag(z)*imag(z)))
			mag = max(0, min(mag, 1))
			v := uint8(mag*255 + 0.5)
			ximg.Set(x, y, color.RGBA{v, v, v, 255})
		}
	}
	ximg.isGray = true
	return ximg
}

// FromComplex128 creates a grayscale image from a flat []complex128 using magnitude.
func FromComplex128(data []complex128, width, height int) *Ximg {
	ximg := NewRGBA(width, height)
	for x := range width {
		for y := range height {
			z := data[x+width*y]
			mag := math.Sqrt(real(z)*real(z) + imag(z)*imag(z))
			mag = max(0, min(mag, 1))
			v := uint8(mag*255 + 0.5)
			ximg.Set(x, y, color.RGBA{v, v, v, 255})
		}
	}
	ximg.isGray = true
	return ximg
}

// FromComplex128_2D creates a grayscale image from a 2D [][]complex128 using magnitude.
func FromComplex128_2D(data [][]complex128) *Ximg {
	if len(data) == 0 {
		return NewRGBA(0, 0)
	}
	width := len(data)
	height := len(data[0])
	ximg := NewRGBA(width, height)
	for x := range width {
		for y := range height {
			z := data[x][y]
			mag := math.Sqrt(real(z)*real(z) + imag(z)*imag(z))
			mag = max(0, min(mag, 1))
			v := uint8(mag*255 + 0.5)
			ximg.Set(x, y, color.RGBA{v, v, v, 255})
		}
	}
	ximg.isGray = true
	return ximg
}

// ToFloat converts a grayscale image to a flat []F in [0, 1].
func ToFloat[F Float](img *Ximg) []F {
	g := img.Gray()
	w, h := g.Size()
	values := make([]F, w*h)
	for x := range w {
		for y := range h {
			r, _, _ := g.RGBAt(x, y)
			values[x+w*y] = F(r) / F(255)
		}
	}
	return values
}

// ToFloat2D converts a grayscale image to [][]F in [0, 1] (width × height).
func ToFloat2D[F Float](img *Ximg) [][]F {
	g := img.Gray()
	w, h := g.Size()
	values := make([][]F, w)
	for x := range w {
		values[x] = make([]F, h)
		for y := range h {
			r, _, _ := g.RGBAt(x, y)
			values[x][y] = F(r) / F(255)
		}
	}
	return values
}

// ToComplex converts a grayscale image to a flat []complex64.
func ToComplex(img *Ximg) []complex64 {
	g := img.Gray()
	w, h := g.Size()
	values := make([]complex64, w*h)
	for x := range w {
		for y := range h {
			r, _, _ := g.RGBAt(x, y)
			values[x+w*y] = complex(float32(r)/float32(255), float32(0))
		}
	}
	return values
}

// ToComplex_2D converts a grayscale image to [][]complex64 (width × height).
func ToComplex_2D(img *Ximg) [][]complex64 {
	g := img.Gray()
	w, h := g.Size()
	values := make([][]complex64, w)
	for x := range w {
		values[x] = make([]complex64, h)
		for y := range h {
			r, _, _ := g.RGBAt(x, y)
			values[x][y] = complex(float32(r)/float32(255), float32(0))
		}
	}
	return values
}

// ToComplex128 converts a grayscale image to a flat []complex128.
func ToComplex128(img *Ximg) []complex128 {
	g := img.Gray()
	w, h := g.Size()
	values := make([]complex128, w*h)
	for x := range w {
		for y := range h {
			r, _, _ := g.RGBAt(x, y)
			values[x+w*y] = complex(float64(r)/float64(255), float64(0))
		}
	}
	return values
}

// ToComplex128_2D converts a grayscale image to [][]complex128 (width × height).
func ToComplex128_2D(img *Ximg) [][]complex128 {
	g := img.Gray()
	w, h := g.Size()
	values := make([][]complex128, w)
	for x := range w {
		values[x] = make([]complex128, h)
		for y := range h {
			r, _, _ := g.RGBAt(x, y)
			values[x][y] = complex(float64(r)/float64(255), float64(0))
		}
	}
	return values
}
