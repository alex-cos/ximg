package ximg_test

import (
	"fmt"
	"image/color"
	"path/filepath"
	"strings"
	"testing"

	"github.com/alex-cos/ximg"
	"github.com/stretchr/testify/require"
)

func TestResize(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
		width    int
		height   int
	}{
		{"img_01.jpg", 256, 256},
		{"img_02.jpg", 128, 128},
		{"img_03.jpg", 64, 64},
	}

	for _, test := range tests {
		ext := filepath.Ext(test.filename)
		infile := filepath.Join("testdata", test.filename)
		outfile := fmt.Sprintf("%s_Resize%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		res := img.Resize(test.width, test.height)
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestSplit(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
		width    int
		shift    int
	}{
		{"img_01.jpg", 6000, 4000},
		{"img_02.jpg", 2048, 2048},
		{"img_03.jpg", 1024, 512},
	}
	for _, test := range tests {
		ext := filepath.Ext(test.filename)
		infile := filepath.Join("testdata", test.filename)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		array := img.Split(test.width, test.shift)
		for i := range array {
			outfile := fmt.Sprintf("%s_Split_%d%s", strings.ReplaceAll(test.filename, ext, ""), i, ext)
			out := filepath.Join("output", outfile)
			err = array[i].Save(out, 90)
			require.NoError(t, err)
		}
	}
}

func TestRed(t *testing.T) {
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
		outfile := fmt.Sprintf("%s_Red%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		res := img.Red()
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestGreen(t *testing.T) {
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
		outfile := fmt.Sprintf("%s_Green%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		res := img.Green()
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestBlue(t *testing.T) {
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
		outfile := fmt.Sprintf("%s_Blue%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		res := img.Blue()
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestGrey(t *testing.T) {
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
		outfile := fmt.Sprintf("%s_Grey%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		res := img.Grey()
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestGray(t *testing.T) {
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
		outfile := fmt.Sprintf("%s_Gray%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		res := img.Gray()
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestInvert(t *testing.T) {
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
		outfile := fmt.Sprintf("%s_Invert%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		res := img.Invert()
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestFlipH(t *testing.T) {
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
		outfile := fmt.Sprintf("%s_FlipH%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		res := img.FlipH()
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestFlipV(t *testing.T) {
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
		outfile := fmt.Sprintf("%s_FlipV%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		res := img.FlipV()
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestNormalize(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
		min      uint8
		max      uint8
	}{
		{"img_01.jpg", 0, 255},
		{"img_02.jpg", 2, 248},
		{"img_03.jpg", 32, 212},
	}
	for _, test := range tests {
		ext := filepath.Ext(test.filename)
		infile := filepath.Join("testdata", test.filename)
		outfile := fmt.Sprintf("%s_Normalize%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		res := img.Normalize(test.min, test.max)
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestMerge(t *testing.T) {
	t.Parallel()
	setupTest(t)

	img1, err := ximg.Load("testdata/img_02.jpg")
	require.NoError(t, err)

	img2, err := ximg.Load("testdata/img_03.jpg")
	require.NoError(t, err)

	res := img1.Merge(img2)
	err = res.Save("output/merge.jpg", 90)
	require.NoError(t, err)
}

func TestFuzion(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
	}{
		{"img_01.jpg"}, {"img_02.jpg"}, {"img_03.jpg"},
	}
	logo, err := ximg.Load("testdata/logo.png")
	require.NoError(t, err)
	for _, test := range tests {
		ext := filepath.Ext(test.filename)
		infile := filepath.Join("testdata", test.filename)
		outfile := fmt.Sprintf("%s_Fuzion%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		res := img.Fuzion(50, 50, logo)
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestMaxPool(t *testing.T) {
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
		outfile := fmt.Sprintf("%s_MaxPool%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		res := img.MaxPool(2)
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestContrast(t *testing.T) {
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
		outfile := fmt.Sprintf("%s_Contrast%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		res := img.Contrast(0.05)
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestBrightness(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
		delta    int
	}{
		{"img_01.jpg", -50},
		{"img_02.jpg", 50},
		{"img_03.jpg", 20},
	}
	for _, test := range tests {
		ext := filepath.Ext(test.filename)
		infile := filepath.Join("testdata", test.filename)
		outfile := fmt.Sprintf("%s_Brightness%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		res := img.Brightness(test.delta)
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestHue(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
		shift    float64
	}{
		{"img_01.jpg", 90},
		{"img_02.jpg", -45},
		{"img_03.jpg", 20},
	}
	for _, test := range tests {
		ext := filepath.Ext(test.filename)
		infile := filepath.Join("testdata", test.filename)
		outfile := fmt.Sprintf("%s_Hue%s", strings.ReplaceAll(test.filename, ext, ""), ext)
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		res := img.Hue(test.shift)
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestColorToAlpha(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename  string
		tolerance float32
	}{
		{"img_01.jpg", 0},
		{"img_02.jpg", 0.12},
		{"img_03.jpg", 0.1},
	}
	for _, test := range tests {
		ext := filepath.Ext(test.filename)
		infile := filepath.Join("testdata", test.filename)
		outfile := fmt.Sprintf("%s_ColorToAlpha%s", strings.ReplaceAll(test.filename, ext, ""), ".png")
		out := filepath.Join("output", outfile)
		img, err := ximg.Load(infile)
		require.NoError(t, err)

		res := img.ColorToAlpha(color.RGBA{255, 255, 255, 255}, test.tolerance)
		err = res.Save(out, 90)
		require.NoError(t, err)
	}
}

func TestAlphaToColor(t *testing.T) {
	t.Parallel()
	setupTest(t)

	logo, err := ximg.Load("testdata/logo.png")
	require.NoError(t, err)

	res := logo.AlphaToColor(color.RGBA{255, 0, 0, 255})
	err = res.Save("output/logo_AlphaToColor.png", 90)
	require.NoError(t, err)
}

func TestRemoveAlpha(t *testing.T) {
	t.Parallel()
	setupTest(t)

	logo, err := ximg.Load("testdata/logo.png")
	require.NoError(t, err)

	res := logo.RemoveAlpha()
	err = res.Save("output/logo_AlphaToColor.png", 90)
	require.NoError(t, err)
}

func TestConcatV(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
	}{
		{"img_01.jpg"}, {"img_02.jpg"}, {"img_03.jpg"},
	}
	ext := filepath.Ext(tests[0].filename)
	images := make([]*ximg.Ximg, 0, len(tests))
	for _, test := range tests {
		infile := filepath.Join("testdata", test.filename)
		img, err := ximg.Load(infile)
		require.NoError(t, err)
		images = append(images, img.Resize(0, 256))
	}
	outfile := fmt.Sprintf("ConcatV%s", ext)
	out := filepath.Join("output", outfile)

	res := ximg.ConcatV(256, images...)
	err := res.Save(out, 90)
	require.NoError(t, err)
}

func TestConcatH(t *testing.T) {
	t.Parallel()
	setupTest(t)

	var tests = []struct {
		filename string
	}{
		{"img_01.jpg"}, {"img_02.jpg"}, {"img_03.jpg"},
	}
	ext := filepath.Ext(tests[0].filename)
	images := make([]*ximg.Ximg, 0, len(tests))
	for _, test := range tests {
		infile := filepath.Join("testdata", test.filename)
		img, err := ximg.Load(infile)
		require.NoError(t, err)
		images = append(images, img.Resize(0, 256))
	}
	outfile := fmt.Sprintf("ConcatH%s", ext)
	out := filepath.Join("output", outfile)

	res := ximg.ConcatH(256, images...)
	err := res.Save(out, 90)
	require.NoError(t, err)
}
