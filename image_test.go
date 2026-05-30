package ximg_test

import (
	"image"
	"image/color"
	"testing"

	"github.com/alex-cos/ximg"
	"github.com/stretchr/testify/assert"
)

func pixelGray(r, g, b uint8) uint8 {
	return uint8(0.2989*float64(r) + 0.5870*float64(g) + 0.1140*float64(b))
}

func testImage() *ximg.Ximg {
	img := ximg.NewRGBA(2, 2)
	img.Set(0, 0, color.RGBA{255, 0, 0, 255})
	img.Set(1, 0, color.RGBA{0, 255, 0, 255})
	img.Set(0, 1, color.RGBA{0, 0, 255, 255})
	img.Set(1, 1, color.RGBA{128, 128, 128, 255})
	return img
}

func TestToFloat32(t *testing.T) {
	t.Parallel()
	img := testImage()
	res := img.ToFloat32()

	assert.Len(t, res, 4)
	exp := [4]float64{
		float64(pixelGray(255, 0, 0)) / 255,
		float64(pixelGray(0, 255, 0)) / 255,
		float64(pixelGray(0, 0, 255)) / 255,
		float64(pixelGray(128, 128, 128)) / 255,
	}
	for i, v := range res {
		assert.InDelta(t, exp[i], float64(v), 1e-6, "index %d", i)
	}
}

func TestToFloat64(t *testing.T) {
	t.Parallel()
	img := testImage()
	res := img.ToFloat64()

	assert.Len(t, res, 4)
	exp := [4]float64{
		float64(pixelGray(255, 0, 0)) / 255,
		float64(pixelGray(0, 255, 0)) / 255,
		float64(pixelGray(0, 0, 255)) / 255,
		float64(pixelGray(128, 128, 128)) / 255,
	}
	for i, v := range res {
		assert.InDelta(t, exp[i], v, 1e-12, "index %d", i)
	}
}

func TestToFloat32_2D(t *testing.T) {
	t.Parallel()
	img := testImage()
	res := img.ToFloat32_2D()

	assert.Len(t, res, 2)
	assert.Len(t, res[0], 2)

	exp := [2][2]float64{
		{float64(pixelGray(255, 0, 0)) / 255, float64(pixelGray(0, 0, 255)) / 255},
		{float64(pixelGray(0, 255, 0)) / 255, float64(pixelGray(128, 128, 128)) / 255},
	}
	for x := range 2 {
		for y := range 2 {
			assert.InDelta(t, exp[x][y], float64(res[x][y]), 1e-6, "res[%d][%d]", x, y)
		}
	}
}

func TestToFloat64_2D(t *testing.T) {
	t.Parallel()
	img := testImage()
	res := img.ToFloat64_2D()

	assert.Len(t, res, 2)
	assert.Len(t, res[0], 2)

	exp := [2][2]float64{
		{float64(pixelGray(255, 0, 0)) / 255, float64(pixelGray(0, 0, 255)) / 255},
		{float64(pixelGray(0, 255, 0)) / 255, float64(pixelGray(128, 128, 128)) / 255},
	}
	for x := range 2 {
		for y := range 2 {
			assert.InDelta(t, exp[x][y], res[x][y], 1e-12, "res[%d][%d]", x, y)
		}
	}
}

func TestToComplex(t *testing.T) {
	t.Parallel()
	img := testImage()
	res := img.ToComplex()

	assert.Len(t, res, 4)

	exp := [4]float64{
		float64(pixelGray(255, 0, 0)) / 255,
		float64(pixelGray(0, 255, 0)) / 255,
		float64(pixelGray(0, 0, 255)) / 255,
		float64(pixelGray(128, 128, 128)) / 255,
	}
	for i, v := range res {
		assert.Equal(t, float32(0), imag(v), "index %d imag", i)
		assert.InDelta(t, exp[i], float64(real(v)), 1e-6, "index %d real", i)
	}
}

func TestToComplex_2D(t *testing.T) {
	t.Parallel()
	img := testImage()
	res := img.ToComplex_2D()

	assert.Len(t, res, 2)
	assert.Len(t, res[0], 2)

	exp := [2][2]float64{
		{float64(pixelGray(255, 0, 0)) / 255, float64(pixelGray(0, 0, 255)) / 255},
		{float64(pixelGray(0, 255, 0)) / 255, float64(pixelGray(128, 128, 128)) / 255},
	}
	for x := range 2 {
		for y := range 2 {
			assert.Equal(t, float32(0), imag(res[x][y]), "res[%d][%d] imag", x, y)
			assert.InDelta(t, exp[x][y], float64(real(res[x][y])), 1e-6, "res[%d][%d] real", x, y)
		}
	}
}

func TestToComplex128(t *testing.T) {
	t.Parallel()
	img := testImage()
	res := img.ToComplex128()

	assert.Len(t, res, 4)

	exp := [4]float64{
		float64(pixelGray(255, 0, 0)) / 255,
		float64(pixelGray(0, 255, 0)) / 255,
		float64(pixelGray(0, 0, 255)) / 255,
		float64(pixelGray(128, 128, 128)) / 255,
	}
	for i, v := range res {
		assert.Equal(t, 0.0, imag(v), "index %d imag", i)
		assert.InDelta(t, exp[i], real(v), 1e-12, "index %d real", i)
	}
}

func TestToComplex128_2D(t *testing.T) {
	t.Parallel()
	img := testImage()
	res := img.ToComplex128_2D()

	assert.Len(t, res, 2)
	assert.Len(t, res[0], 2)

	exp := [2][2]float64{
		{float64(pixelGray(255, 0, 0)) / 255, float64(pixelGray(0, 0, 255)) / 255},
		{float64(pixelGray(0, 255, 0)) / 255, float64(pixelGray(128, 128, 128)) / 255},
	}
	for x := range 2 {
		for y := range 2 {
			assert.Equal(t, 0.0, imag(res[x][y]), "res[%d][%d] imag", x, y)
			assert.InDelta(t, exp[x][y], real(res[x][y]), 1e-12, "res[%d][%d] real", x, y)
		}
	}
}

func TestToFloatGrayImage(t *testing.T) {
	t.Parallel()

	img := ximg.NewRGBA(2, 1)
	img.Set(0, 0, color.RGBA{128, 128, 128, 255})
	img.Set(1, 0, color.RGBA{200, 200, 200, 255})

	res := img.ToFloat32()
	assert.Len(t, res, 2)
	assert.Equal(t, float32(pixelGray(128, 128, 128))/255, res[0])
	assert.Equal(t, float32(pixelGray(200, 200, 200))/255, res[1])
}

func TestToFloatEmptyImage(t *testing.T) {
	t.Parallel()

	img := ximg.NewRGBA(0, 0)

	assert.Empty(t, img.ToFloat64())
	assert.Empty(t, img.ToFloat64_2D())
}

func TestToComplex128EmptyImage(t *testing.T) {
	t.Parallel()

	img := ximg.NewRGBA(0, 0)

	assert.Empty(t, img.ToComplex128())
	assert.Empty(t, img.ToComplex128_2D())
}

func TestSize(t *testing.T) {
	t.Parallel()

	img := ximg.NewRGBA(10, 20)
	w, h := img.Size()
	assert.Equal(t, 10, w)
	assert.Equal(t, 20, h)
}

func TestRGBAt(t *testing.T) {
	t.Parallel()

	img := ximg.NewRGBA(1, 1)
	img.Set(0, 0, color.RGBA{100, 150, 200, 255})

	r, g, b := img.RGBAt(0, 0)
	assert.Equal(t, uint8(100), r)
	assert.Equal(t, uint8(150), g)
	assert.Equal(t, uint8(200), b)
}

func TestRGBAAt(t *testing.T) {
	t.Parallel()

	img := ximg.NewRGBA(1, 1)
	img.Set(0, 0, color.RGBA{100, 150, 200, 128})

	r, g, b, a := img.RGBAAt(0, 0)
	assert.Equal(t, uint8(100), r)
	assert.Equal(t, uint8(150), g)
	assert.Equal(t, uint8(200), b)
	assert.Equal(t, uint8(128), a)
}

func TestIsGray(t *testing.T) {
	t.Parallel()

	img := ximg.NewRGBA(1, 1)
	assert.False(t, img.IsGray())

	gray := img.Gray()
	assert.True(t, gray.IsGray())
}

func TestNewFromImage(t *testing.T) {
	t.Parallel()

	src := image.NewRGBA(image.Rect(0, 0, 3, 3))
	img, err := ximg.New(src)
	assert.NoError(t, err)
	w, h := img.Size()
	assert.Equal(t, 3, w)
	assert.Equal(t, 3, h)
}

func TestFromFloatRoundTrip(t *testing.T) {
	t.Parallel()

	img := testImage()
	gray := img.Gray()
	data := ximg.ToFloat[float32](gray)
	recovered := ximg.FromFloat(data, 2, 2)

	for x := range 2 {
		for y := range 2 {
			r1, _, _ := gray.RGBAt(x, y)
			r2, _, _ := recovered.RGBAt(x, y)
			assert.InDelta(t, float64(r1), float64(r2), 1, "pixel (%d,%d)", x, y)
		}
	}
}

func TestFromFloatEmpty(t *testing.T) {
	t.Parallel()

	img := ximg.FromFloat([]float32{}, 0, 0)
	w, h := img.Size()
	assert.Equal(t, 0, w)
	assert.Equal(t, 0, h)
}

func TestFromFloatRoundTripFloat64(t *testing.T) {
	t.Parallel()

	img := testImage()
	data := img.ToFloat64()
	recovered := ximg.FromFloat(data, 2, 2)

	r1, _, _ := img.RGBAt(1, 1)
	r2, _, _ := recovered.RGBAt(1, 1)
	assert.InDelta(t, float64(r1), float64(r2), 1, "pixel (1,1)")
}

func TestFromFloatIsGray(t *testing.T) {
	t.Parallel()

	img := ximg.FromFloat([]float32{0.5, 0.5, 0.5, 0.5}, 2, 2)
	assert.True(t, img.IsGray())
}

func TestFromFloat2DRoundTrip(t *testing.T) {
	t.Parallel()

	img := testImage()
	gray := img.Gray()
	data := ximg.ToFloat2D[float32](gray)
	recovered := ximg.FromFloat2D(data)

	for x := range 2 {
		for y := range 2 {
			r1, _, _ := gray.RGBAt(x, y)
			r2, _, _ := recovered.RGBAt(x, y)
			assert.InDelta(t, float64(r1), float64(r2), 1, "pixel (%d,%d)", x, y)
		}
	}
}

func TestFromFloat2DEmpty(t *testing.T) {
	t.Parallel()

	img := ximg.FromFloat2D([][]float32{})
	w, h := img.Size()
	assert.Equal(t, 0, w)
	assert.Equal(t, 0, h)
}

func TestFromFloat2DIsGray(t *testing.T) {
	t.Parallel()

	data := [][]float32{{0.25, 0.5}, {0.75, 1.0}}
	img := ximg.FromFloat2D(data)
	assert.True(t, img.IsGray())
	w, h := img.Size()
	assert.Equal(t, 2, w)
	assert.Equal(t, 2, h)
}

func TestFromComplexMagnitude(t *testing.T) {
	t.Parallel()

	data := []complex64{1 + 0i, 0 + 1i, 3 + 4i, 0 + 0i}
	img := ximg.FromComplex(data, 2, 2)

	r, _, _ := img.RGBAt(0, 0)
	assert.Equal(t, uint8(255), r, "|1+0i| = 1")

	r, _, _ = img.RGBAt(1, 0)
	assert.Equal(t, uint8(255), r, "|0+1i| = 1")

	r, _, _ = img.RGBAt(0, 1)
	assert.Equal(t, uint8(255), r, "|3+4i| = 5, clamped to 1")

	r, _, _ = img.RGBAt(1, 1)
	assert.Equal(t, uint8(0), r, "|0+0i| = 0")
}

func TestFromComplexRoundTrip(t *testing.T) {
	t.Parallel()

	img := testImage()
	gray := img.Gray()
	data := ximg.ToComplex(gray)
	recovered := ximg.FromComplex(data, 2, 2)

	for x := range 2 {
		for y := range 2 {
			r1, _, _ := gray.RGBAt(x, y)
			r2, _, _ := recovered.RGBAt(x, y)
			assert.InDelta(t, float64(r1), float64(r2), 1, "pixel (%d,%d)", x, y)
		}
	}
}

func TestFromComplexZero(t *testing.T) {
	t.Parallel()

	img := ximg.FromComplex([]complex64{0 + 0i, 0 + 0i}, 2, 1)
	r, _, _ := img.RGBAt(0, 0)
	assert.Equal(t, uint8(0), r)
	assert.True(t, img.IsGray())
}

func TestFromComplex2D(t *testing.T) {
	t.Parallel()

	data := [][]complex64{
		{1 + 0i, 0.5 + 0i},
		{0 + 0i, 3 + 4i},
	}
	img := ximg.FromComplex2D(data)

	r, _, _ := img.RGBAt(0, 0)
	assert.Equal(t, uint8(255), r)

	r, _, _ = img.RGBAt(0, 1)
	assert.InDelta(t, uint8(128), r, 1)

	r, _, _ = img.RGBAt(1, 0)
	assert.Equal(t, uint8(0), r)

	assert.True(t, img.IsGray())
}

func TestFromComplexEmpty(t *testing.T) {
	t.Parallel()

	img := ximg.FromComplex([]complex64{}, 0, 0)
	w, h := img.Size()
	assert.Equal(t, 0, w)
	assert.Equal(t, 0, h)
}

func TestFromComplex2DEmpty(t *testing.T) {
	t.Parallel()

	img := ximg.FromComplex2D([][]complex64{})
	w, h := img.Size()
	assert.Equal(t, 0, w)
	assert.Equal(t, 0, h)
}

func TestFromComplex128Magnitude(t *testing.T) {
	t.Parallel()

	data := []complex128{1 + 0i, 0 + 1i, 3 + 4i, 0 + 0i}
	img := ximg.FromComplex128(data, 2, 2)

	r, _, _ := img.RGBAt(0, 0)
	assert.Equal(t, uint8(255), r)

	r, _, _ = img.RGBAt(1, 0)
	assert.Equal(t, uint8(255), r)

	r, _, _ = img.RGBAt(0, 1)
	assert.Equal(t, uint8(255), r, "|3+4i| clamped")

	r, _, _ = img.RGBAt(1, 1)
	assert.Equal(t, uint8(0), r)
}

func TestFromComplex128_2D(t *testing.T) {
	t.Parallel()

	data := [][]complex128{
		{1 + 0i, 0 + 0i},
		{0 + 0i, 0 + 0i},
	}
	img := ximg.FromComplex128_2D(data)

	r, _, _ := img.RGBAt(0, 0)
	assert.Equal(t, uint8(255), r)

	r, _, _ = img.RGBAt(1, 0)
	assert.Equal(t, uint8(0), r)

	assert.True(t, img.IsGray())
}
