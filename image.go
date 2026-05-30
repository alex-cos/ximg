package ximg

import (
	"image"
	"image/color"
	"image/draw"
)

// New wraps an image.Image into an Ximg, converting to RGBA if needed.
func New(image image.Image) (*Ximg, error) {
	if image == nil {
		return nil, ErrInputImageIsNil
	}
	return &Ximg{
		RGBA:   *imageToRGBA(image),
		isGray: false,
	}, nil
}

// NewRGBA creates a new blank RGBA image with the given dimensions.
func NewRGBA(width, height int) *Ximg {
	ximg, _ := New(image.NewRGBA(image.Rect(0, 0, width, height)))
	return ximg
}

// NewFromRGB combines three single-channel images into a color image.
func NewFromRGB(red, green, blue *Ximg) (*Ximg, error) {
	width, height := red.Size()
	ximg := NewRGBA(width, height)

	wg, hg := green.Size()
	wb, hb := blue.Size()
	if (wg != width) || (hg != height) || (wb != width) || (hb != height) {
		return nil, ErrInputImagesNotSameSize
	}

	for y := range height {
		for x := range width {
			r, _, _ := red.RGBAt(x, y)
			g, _, _ := green.RGBAt(x, y)
			b, _, _ := blue.RGBAt(x, y)
			ximg.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
	return ximg, nil
}

// ----------------------------------------------------------------------------
// Exported functions
// ----------------------------------------------------------------------------

func (img *Ximg) Clone() *Ximg {
	width, height := img.Size()
	dst := NewRGBA(width, height)
	copy(dst.Pix, img.Pix)
	dst.isGray = img.isGray

	return dst
}

// Size returns the width and height of the image.
func (img *Ximg) Size() (int, int) {
	size := img.Bounds().Size()
	return size.X, size.Y
}

// RGBAt returns the red, green, blue values at (x, y).
func (img *Ximg) RGBAt(x, y int) (uint8, uint8, uint8) {
	r, g, b, _ := img.At(x, y).RGBA()
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8) // nolint: gosec
}

// RGBAAt returns the red, green, blue and alpha values at (x, y).
func (img *Ximg) RGBAAt(x, y int) (uint8, uint8, uint8, uint8) {
	r, g, b, a := img.At(x, y).RGBA()
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8) // nolint: gosec
}

// AlphaAt returns the alpha value at (x, y).
func (img *Ximg) AlphaAt(x, y int) uint8 {
	_, _, _, a := img.At(x, y).RGBA() // nolint: dogsled
	return uint8(a >> 8)              // nolint: gosec
}

// IsGray returns whether the image is known to be grayscale.
func (img *Ximg) IsGray() bool {
	return img.isGray
}

// ToFloat32 converts the grayscale image to a flat []float32 in [0, 1].
func (img *Ximg) ToFloat32() []float32 {
	return ToFloat[float32](img)
}

// ToFloat64 converts the grayscale image to a flat []float64 in [0, 1].
func (img *Ximg) ToFloat64() []float64 {
	return ToFloat[float64](img)
}

// ToFloat32_2D converts the grayscale image to [][]float32 in [0, 1] (width × height).
func (img *Ximg) ToFloat32_2D() [][]float32 {
	return ToFloat2D[float32](img)
}

// ToFloat64_2D converts the grayscale image to [][]float64 in [0, 1] (width × height).
func (img *Ximg) ToFloat64_2D() [][]float64 {
	return ToFloat2D[float64](img)
}

// ToComplex converts the grayscale image to a flat []complex64.
func (img *Ximg) ToComplex() []complex64 {
	return ToComplex(img)
}

// ToComplex_2D converts the grayscale image to [][]complex64 (width × height).
func (img *Ximg) ToComplex_2D() [][]complex64 {
	return ToComplex_2D(img)
}

// ToComplex128 converts the grayscale image to a flat []complex128.
func (img *Ximg) ToComplex128() []complex128 {
	return ToComplex128(img)
}

// ToComplex128_2D converts the grayscale image to [][]complex128 (width × height).
func (img *Ximg) ToComplex128_2D() [][]complex128 {
	return ToComplex128_2D(img)
}

// ----------------------------------------------------------------------------
// Unexported functions
// ----------------------------------------------------------------------------

func imageToRGBA(src image.Image) *image.RGBA {
	if dst, ok := src.(*image.RGBA); ok {
		return dst
	}
	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Src)
	return dst
}
