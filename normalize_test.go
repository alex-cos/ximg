package ximg_test

import (
	"math"
	"testing"

	"github.com/alex-cos/ximg"
	"github.com/stretchr/testify/assert"
)

func TestNormalizeComplex(t *testing.T) {
	t.Parallel()

	data := []complex64{0 + 0i, 3 + 4i, 1 + 0i}
	res := ximg.NormalizeComplex(data)

	assert.Len(t, res, 3)
	assert.InDelta(t, 0, real(res[0]), 1e-6)
	assert.InDelta(t, 0, imag(res[0]), 1e-6)

	// |3+4i| = 5, scaled by 1/5
	assert.InDelta(t, 0.6, real(res[1]), 1e-6)
	assert.InDelta(t, 0.8, imag(res[1]), 1e-6)

	// |1+0i| = 1, scaled by 1/5
	assert.InDelta(t, 0.2, real(res[2]), 1e-6)
	assert.InDelta(t, 0, imag(res[2]), 1e-6)
}

func TestNormalizeComplexPreservesPhase(t *testing.T) {
	t.Parallel()

	data := []complex64{2 + 2i, 4 + 0i}
	res := ximg.NormalizeComplex(data)

	// phase of 2+2i = pi/4, magnitude = 2√2 ≈ 2.828
	// max = 4, scale = 1/4
	// 2+2i → 0.5+0.5i, phase should still be pi/4
	expected := float32(math.Pi / 4)
	got := float32(math.Atan2(float64(imag(res[0])), float64(real(res[0]))))
	assert.InDelta(t, expected, got, 1e-6)
}

func TestNormalizeComplexAllZeros(t *testing.T) {
	t.Parallel()

	data := []complex64{0 + 0i, 0 + 0i}
	res := ximg.NormalizeComplex(data)

	assert.Len(t, res, 2)
	assert.Equal(t, float32(0), real(res[0]))
	assert.Equal(t, float32(0), imag(res[0]))
}

func TestNormalizeComplexEmpty(t *testing.T) {
	t.Parallel()

	res := ximg.NormalizeComplex(nil)
	assert.Nil(t, res)
}

func TestNormalizeComplex2D(t *testing.T) {
	t.Parallel()

	data := [][]complex64{
		{0 + 0i, 3 + 4i},
		{1 + 0i, 0 + 1i},
	}
	res := ximg.NormalizeComplex2D(data)

	assert.Len(t, res, 2)
	assert.Len(t, res[0], 2)

	// max = 5, scale = 1/5
	assert.InDelta(t, 0.6, real(res[0][1]), 1e-6)
	assert.InDelta(t, 0.8, imag(res[0][1]), 1e-6)
	assert.InDelta(t, 0.2, real(res[1][0]), 1e-6)
	assert.InDelta(t, 0.2, imag(res[1][1]), 1e-6)
}

func TestNormalizeComplex2DEmpty(t *testing.T) {
	t.Parallel()

	res := ximg.NormalizeComplex2D([][]complex64{})
	assert.Nil(t, res)
}

func TestNormalizeComplex2DAllZeros(t *testing.T) {
	t.Parallel()

	data := [][]complex64{{0 + 0i, 0 + 0i}, {0 + 0i, 0 + 0i}}
	res := ximg.NormalizeComplex2D(data)

	assert.Len(t, res, 2)
	assert.Len(t, res[0], 2)
	assert.Equal(t, float32(0), real(res[0][0]))
	assert.Equal(t, float32(0), imag(res[0][0]))
}

func TestNormalizeComplexFromFFT(t *testing.T) {
	t.Parallel()

	n := 32
	input := make([]float64, n)
	for i := range input {
		input[i] = math.Sin(2 * math.Pi * float64(i) / float64(n))
	}
	spectrum := ximg.FFT(input)
	normalized := ximg.NormalizeComplex(spectrum)

	// all magnitudes should be in [0,1]
	for _, z := range normalized {
		mag := math.Sqrt(float64(real(z)*real(z) + imag(z)*imag(z)))
		assert.LessOrEqual(t, mag, 1.0+1e-6)
	}

	img := ximg.FromComplex(normalized, n, 1)
	assert.NotNil(t, img)
}
