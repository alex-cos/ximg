package ximg_test

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/alex-cos/ximg"
	"github.com/stretchr/testify/require"
)

func TestBlur(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
	}{
		{"img_01.jpg"}, {"img_02.jpg"}, {"img_03.jpg"},
	}
	for _, test := range tests {
		ext := filepath.Ext(test.filename)
		infile := filepath.Join("testdata", test.filename)
		outfile := fmt.Sprintf("%s_Blur%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		img = img.Resize(256, 0)
		res := img.Blur()
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestSharpen(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
	}{
		{"img_01.jpg"}, {"img_02.jpg"}, {"img_03.jpg"},
	}
	for _, test := range tests {
		ext := filepath.Ext(test.filename)
		infile := filepath.Join("testdata", test.filename)
		outfile := fmt.Sprintf("%s_Sharpen%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		img = img.Resize(256, 0)
		res := img.Sharpen()
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestGaussian5(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
	}{
		{"img_01.jpg"}, {"img_02.jpg"}, {"img_03.jpg"},
	}
	for _, test := range tests {
		ext := filepath.Ext(test.filename)
		infile := filepath.Join("testdata", test.filename)
		outfile := fmt.Sprintf("%s_Gaussian5%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		img = img.Resize(256, 0)
		res := img.Gaussian5()
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestGaussian7(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
	}{
		{"img_01.jpg"}, {"img_02.jpg"}, {"img_03.jpg"},
	}
	for _, test := range tests {
		ext := filepath.Ext(test.filename)
		infile := filepath.Join("testdata", test.filename)
		outfile := fmt.Sprintf("%s_Gaussian7%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		img = img.Resize(256, 0)
		res := img.Gaussian7()
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestEdgeDetect(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
	}{
		{"img_01.jpg"}, {"img_02.jpg"}, {"img_03.jpg"},
	}
	for _, test := range tests {
		ext := filepath.Ext(test.filename)
		infile := filepath.Join("testdata", test.filename)
		outfile := fmt.Sprintf("%s_EdgeDetect%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		img = img.Resize(256, 0)
		res := img.EdgeDetect()
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestSobel(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
	}{
		{"img_01.jpg"}, {"img_02.jpg"}, {"img_03.jpg"},
	}
	for _, test := range tests {
		ext := filepath.Ext(test.filename)
		infile := filepath.Join("testdata", test.filename)
		outfile := fmt.Sprintf("%s_Sobel%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		img = img.Resize(256, 0)
		res := img.Sobel()
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestSobel5(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
	}{
		{"img_01.jpg"}, {"img_02.jpg"}, {"img_03.jpg"},
	}
	for _, test := range tests {
		ext := filepath.Ext(test.filename)
		infile := filepath.Join("testdata", test.filename)
		outfile := fmt.Sprintf("%s_Sobel5%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		img = img.Resize(256, 0)
		img = img.Gray()
		res := img.Sobel5()
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestSobel2(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
	}{
		{"img_01.jpg"}, {"img_02.jpg"}, {"img_03.jpg"},
	}
	for _, test := range tests {
		ext := filepath.Ext(test.filename)
		infile := filepath.Join("testdata", test.filename)
		outfile := fmt.Sprintf("%s_Sobel2%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		img = img.Resize(256, 0)
		res := img.Sobel2()
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}
