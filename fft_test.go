package ximg_test

import (
	"fmt"
	"math"
	"path/filepath"
	"strings"
	"testing"

	"github.com/alex-cos/ximg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFFTImage(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
	}{
		{"img_04.jpg"}, {"img_05.jpg"},
	}
	for _, test := range tests {
		ext := filepath.Ext(test.filename)
		infile := filepath.Join("testdata", test.filename)
		outfile := fmt.Sprintf("%s_FFT%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)
		input := img.ToFloat64()
		spectrum := ximg.FFT(input)
		mag := ximg.LogMagnitude(spectrum)

		w, h := img.Size()
		res := ximg.FromFloat(ximg.NormalizeFloat(mag), w, h)
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestFFTConstant(t *testing.T) {
	t.Parallel()

	input := []float64{1, 1, 1, 1}
	res := ximg.FFT(input)

	assert.Len(t, res, 4)
	assert.InDelta(t, 4, real(res[0]), 1e-6, "DC component")
	assert.InDelta(t, 0, real(res[1]), 1e-6, "bin 1 real")
	assert.InDelta(t, 0, imag(res[1]), 1e-6, "bin 1 imag")
	assert.InDelta(t, 0, real(res[2]), 1e-6, "bin 2 real")
	assert.InDelta(t, 0, imag(res[3]), 1e-6, "bin 3 imag")
}

func TestFFTImpulse(t *testing.T) {
	t.Parallel()

	input := []float64{1, 0, 0, 0}
	res := ximg.FFT(input)

	assert.Len(t, res, 4)
	for i := range res {
		assert.InDelta(t, 1, real(res[i]), 1e-6, "bin %d real", i)
		assert.InDelta(t, 0, imag(res[i]), 1e-6, "bin %d imag", i)
	}
}

func TestFFTSine(t *testing.T) {
	t.Parallel()

	n := 64
	input := make([]float64, n)
	for i := range input {
		input[i] = math.Sin(2 * math.Pi * float64(i) / float64(n))
	}

	res := ximg.FFT(input)

	assert.Len(t, res, 64)
	mag := math.Sqrt(float64(real(res[1])*real(res[1]) + imag(res[1])*imag(res[1])))
	assert.InDelta(t, 32, mag, 1, "bin 1 magnitude")
}

func TestFFTPadding(t *testing.T) {
	t.Parallel()

	input := []float32{1, 2, 3}
	res := ximg.FFT(input)

	assert.Len(t, res, 4, "should pad to next power of 2")
}

func TestFFTSymmetry(t *testing.T) {
	t.Parallel()

	n := 32
	input := make([]float64, n)
	for i := range input {
		input[i] = float64(i)
	}

	res := ximg.FFT(input)

	assert.Len(t, res, 32)
	for i := 1; i < n/2; i++ {
		re1, im1 := real(res[i]), imag(res[i])
		re2, im2 := real(res[n-i]), imag(res[n-i])
		assert.InDelta(t, re1, re2, 1e-5, "real symmetry broken at bin %d", i)
		assert.InDelta(t, im1, -im2, 1e-5, "imag antisymmetry broken at bin %d", i)
	}
}

func TestFFTEmpty(t *testing.T) {
	t.Parallel()

	input := []float64{}
	res := ximg.FFT(input)

	assert.Len(t, res, 1, "empty input should pad to 1")
}

// FFT2D

func TestFFT2DImage(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
	}{
		{"img_04.jpg"}, {"img_05.jpg"},
	}

	for _, test := range tests {
		ext := filepath.Ext(test.filename)
		infile := filepath.Join("testdata", test.filename)
		outfile := fmt.Sprintf("%s_FFT2D%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		data := img.Gray().ToFloat64_2D()
		spectrum := ximg.FFT2D(data)
		centered := ximg.FFTCenter(spectrum)

		pw := len(centered)
		ph := len(centered[0])
		mag := make([]float64, 0, pw*ph)
		for x := range pw {
			mag = append(mag, ximg.LogMagnitude(centered[x])...)
		}

		res := ximg.FromFloat(ximg.NormalizeFloat(mag), pw, ph)
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestFFT2DConstant(t *testing.T) {
	t.Parallel()

	data := [][]float64{
		{1, 1},
		{1, 1},
	}
	res := ximg.FFT2D(data)

	assert.Len(t, res, 2)
	assert.Len(t, res[0], 2)

	assert.InDelta(t, 4, real(res[0][0]), 1e-5, "DC")
	assert.InDelta(t, 0, real(res[1][0]), 1e-5, "AC x")
	assert.InDelta(t, 0, real(res[0][1]), 1e-5, "AC y")
	assert.InDelta(t, 0, real(res[1][1]), 1e-5, "AC xy")
}

func TestFFT2DImpulse(t *testing.T) {
	t.Parallel()

	data := [][]float64{
		{1, 0},
		{0, 0},
	}
	res := ximg.FFT2D(data)

	assert.Len(t, res, 2)
	assert.Len(t, res[0], 2)

	for x := range 2 {
		for y := range 2 {
			assert.InDelta(t, 1, real(res[x][y]), 1e-5, "(%d,%d)", x, y)
			assert.InDelta(t, 0, imag(res[x][y]), 1e-5, "(%d,%d)", x, y)
		}
	}
}

func TestFFT2DPadding(t *testing.T) {
	t.Parallel()

	data := [][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	res := ximg.FFT2D(data)

	assert.Len(t, res, 4, "width padded to next pow2")
	assert.Len(t, res[0], 2, "height already pow2")
}

func TestFFT2DEmpty(t *testing.T) {
	t.Parallel()

	assert.Nil(t, ximg.FFT2D([][]float64{}))
	assert.Nil(t, ximg.FFT2D([][]float64{{}}))
}

func TestFFTCenterEven(t *testing.T) {
	t.Parallel()

	data := [][]complex64{
		{1, 2},
		{3, 4},
	}
	res := ximg.FFTCenter(data)

	assert.Equal(t, complex64(4), res[0][0], "DC to center")
	assert.Equal(t, complex64(1), res[1][1], "TL to BR")
	assert.Equal(t, complex64(2), res[1][0], "TR to BL")
	assert.Equal(t, complex64(3), res[0][1], "BL to TR")
}

func TestFFTCenterOdd(t *testing.T) {
	t.Parallel()

	data := [][]complex64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	res := ximg.FFTCenter(data)

	assert.Equal(t, complex64(9), res[1][1], "center element")
}

func TestFFTCenterTwice(t *testing.T) {
	t.Parallel()

	data := [][]complex64{
		{1, 2},
		{3, 4},
	}
	res := ximg.FFTCenter(ximg.FFTCenter(data))

	for x := range 2 {
		for y := range 2 {
			assert.Equal(t, data[x][y], res[x][y], "(%d,%d)", x, y)
		}
	}
}

func TestFFTCenterEmpty(t *testing.T) {
	t.Parallel()

	assert.Nil(t, ximg.FFTCenter([][]complex64{}))
	assert.Nil(t, ximg.FFTCenter([][]complex64{{}}))
}
