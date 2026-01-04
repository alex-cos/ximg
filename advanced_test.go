package ximg_test

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/alex-cos/ximg"
	"github.com/stretchr/testify/require"
)

func TestConvolution1(t *testing.T) {
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
		outfile := fmt.Sprintf("%s_Convolution1%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		images := ximg.Convolution1(img)
		res := ximg.ConcatV(0, images...)
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestConvolution2(t *testing.T) {
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
		outfile := fmt.Sprintf("%s_Convolution2%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		images := ximg.Convolution2(img)
		res := ximg.ConcatV(0, images...)
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestConvolution3(t *testing.T) {
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
		outfile := fmt.Sprintf("%s_Convolution3%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		img = img.Resize(256, 0)
		images := ximg.Convolution3(img)
		res := ximg.ConcatV(0, images...)
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}
