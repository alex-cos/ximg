package ximg

import (
	"image/color"
	"math"
	"sync"
)

var (
	// BlurFilter is a 3×3 averaging kernel.
	BlurFilter = [][]float64{
		{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
		{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
		{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
	}
	// SharpenFilter is a 3×3 sharpening kernel.
	SharpenFilter = [][]float64{
		{0.0, -0.5, 0.0},
		{-0.5, 3.0, -0.5},
		{0.0, -0.5, 0.0},
	}
	// Sharpen2Filter is a stronger 3×3 sharpening kernel.
	Sharpen2Filter = [][]float64{
		{0.0, -1.0, 0.0},
		{-1.0, 5.0, -1.0},
		{0.0, -1.0, 0.0},
	}
	// LaplacianFilter is a 3×3 Laplacian edge detection kernel.
	LaplacianFilter = [][]float64{
		{-1.0, -1.0, -1.0},
		{-1.0, 8.0, -1.0},
		{-1.0, -1.0, -1.0},
	}
	// Gaussian5Filter is a 5×5 Gaussian blur kernel.
	Gaussian5Filter = [][]float64{
		{2.0 / 159, 4.0 / 159, 5.0 / 159, 4.0 / 159, 2.0 / 159},
		{4.0 / 159, 9.0 / 159, 12.0 / 159, 9.0 / 159, 4.0 / 159},
		{5.0 / 159, 12.0 / 159, 15.0 / 159, 12.0 / 159, 5.0 / 159},
		{4.0 / 159, 9.0 / 159, 12.0 / 159, 9.0 / 159, 4.0 / 159},
		{2.0 / 159, 4.0 / 159, 5.0 / 159, 4.0 / 159, 2.0 / 159},
	}
	// Gaussian7Filter is a 7×7 Gaussian blur kernel.
	Gaussian7Filter = [][]float64{
		{0.00000067, 0.00002292, 0.00019117, 0.00038771, 0.00019117, 0.00002292, 0.00000067},
		{0.00002292, 0.00078633, 0.00655965, 0.01330373, 0.00655965, 0.00078633, 0.00002292},
		{0.00019117, 0.00655965, 0.05472157, 0.11098164, 0.05472157, 0.00655965, 0.00019117},
		{0.00038771, 0.01330373, 0.11098164, 0.22508352, 0.11098164, 0.01330373, 0.00038771},
		{0.00019117, 0.00655965, 0.05472157, 0.11098164, 0.05472157, 0.00655965, 0.00019117},
		{0.00002292, 0.00078633, 0.00655965, 0.01330373, 0.00655965, 0.00078633, 0.00002292},
		{0.00000067, 0.00002292, 0.00019117, 0.00038771, 0.00019117, 0.00002292, 0.00000067},
	}
	// EdgeDetectFilter is a 3×3 Laplacian-like edge detection kernel.
	EdgeDetectFilter = [][]float64{
		{-1.5, -1.5, -1.5},
		{-1.5, 12, -1.5},
		{-1.5, -1.5, -1.5},
	}
	// EdgeDetectHFilter is a 3×3 horizontal edge detection kernel.
	EdgeDetectHFilter = [][]float64{
		{-1, -1, -1},
		{2, 2, 2},
		{-1, -1, -1},
	}
	// EdgeDetectVFilter is a 3×3 vertical edge detection kernel.
	EdgeDetectVFilter = [][]float64{
		{-1, 2, -1},
		{-1, 2, -1},
		{-1, 2, -1},
	}
	// EdgeDetectD1Filter is a 3×3 diagonal (top-left to bottom-right) edge detection kernel.
	EdgeDetectD1Filter = [][]float64{
		{2, -1, -1},
		{-1, 2, -1},
		{-1, -1, 2},
	}
	// EdgeDetectD2Filter is a 3×3 diagonal (top-right to bottom-left) edge detection kernel.
	EdgeDetectD2Filter = [][]float64{
		{-1, -1, 2},
		{-1, 2, -1},
		{2, -1, -1},
	}
	// SobelHFilter is a 3×3 horizontal Sobel kernel.
	SobelHFilter = [][]float64{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}
	// SobelVFilter is a 3×3 vertical Sobel kernel.
	SobelVFilter = [][]float64{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}
	// Sobel5HFilter is a 5×5 horizontal Sobel kernel.
	Sobel5HFilter = [][]float64{
		{1, 2, 0, -2, -1},
		{4, 8, 0, -8, -4},
		{6, 12, 0, -12, -6},
		{4, 8, 0, -8, -4},
		{1, 2, 0, -2, -1},
	}
	// Sobel5VFilter is a 5×5 vertical Sobel kernel.
	Sobel5VFilter = [][]float64{
		{1, 4, 6, 4, 1},
		{2, 8, 12, 8, 2},
		{0, 0, 0, 0, 0},
		{-2, -8, -12, -8, -2},
		{-1, -4, -6, -4, -1},
	}
)

// ApplyFilter applies a convolution filter to the image.
func (img *Ximg) ApplyFilter(filter [][]float64) *Ximg {
	w, h := img.Size()
	filterSize := len(filter)

	f := func(src *Ximg, dst *Ximg, x int, y int) {
		fr := float64(0.0)
		fg := float64(0.0)
		fb := float64(0.0)
		a := src.AlphaAt(x, y)
		for i := range filterSize {
			for j := range filterSize {
				k := x - (filterSize / 2) + i
				l := y - (filterSize / 2) + j
				if k >= 0 && k < w && l >= 0 && l < h {
					r, g, b := src.RGBAt(k, l)
					fr += filter[i][j] * float64(r)
					if !img.IsGray() {
						fg += filter[i][j] * float64(g)
						fb += filter[i][j] * float64(b)
					}
				}
			}
		}
		fr = math.Min(255.0, math.Max(0.0, fr))
		if img.IsGray() {
			fg = fr
			fb = fr
		} else {
			fg = math.Min(255.0, math.Max(0.0, fg))
			fb = math.Min(255.0, math.Max(0.0, fb))
		}
		dst.Set(x, y, color.RGBA{uint8(fr), uint8(fg), uint8(fb), a})
	}
	ximg := img.Apply(1, f, nil)
	ximg.isGray = img.isGray

	return ximg
}

// Blur applies a 3×3 box blur.
func (img *Ximg) Blur() *Ximg {
	return img.ApplyFilter(BlurFilter)
}

// Sharpen applies a 3×3 sharpen filter.
func (img *Ximg) Sharpen() *Ximg {
	return img.ApplyFilter(SharpenFilter)
}

// Sharpen2 applies a stronger 3×3 sharpen filter.
func (img *Ximg) Sharpen2() *Ximg {
	return img.ApplyFilter(Sharpen2Filter)
}

// Gaussian5 applies a 5×5 Gaussian blur.
func (img *Ximg) Gaussian5() *Ximg {
	return img.ApplyFilter(Gaussian5Filter)
}

// Gaussian7 applies a 7×7 Gaussian blur.
func (img *Ximg) Gaussian7() *Ximg {
	return img.ApplyFilter(Gaussian7Filter)
}

// EdgeDetect applies a 3×3 Laplacian edge detection filter.
func (img *Ximg) EdgeDetect() *Ximg {
	return img.ApplyFilter(EdgeDetectFilter)
}

// EdgeDetectH detects horizontal edges with a 3×3 filter.
func (img *Ximg) EdgeDetectH() *Ximg {
	return img.ApplyFilter(EdgeDetectHFilter)
}

// EdgeDetectV detects vertical edges with a 3×3 filter.
func (img *Ximg) EdgeDetectV() *Ximg {
	return img.ApplyFilter(EdgeDetectVFilter)
}

// EdgeDetectD1 detects diagonal edges (top-left to bottom-right).
func (img *Ximg) EdgeDetectD1() *Ximg {
	return img.ApplyFilter(EdgeDetectD1Filter)
}

// EdgeDetectD2 detects diagonal edges (top-right to bottom-left).
func (img *Ximg) EdgeDetectD2() *Ximg {
	return img.ApplyFilter(EdgeDetectD2Filter)
}

// SobelH applies a 3×3 horizontal Sobel filter.
func (img *Ximg) SobelH() *Ximg {
	return img.ApplyFilter(SobelHFilter)
}

// SobelV applies a 3×3 vertical Sobel filter.
func (img *Ximg) SobelV() *Ximg {
	return img.ApplyFilter(SobelVFilter)
}

// Sobel5H applies a 5×5 horizontal Sobel filter.
func (img *Ximg) Sobel5H() *Ximg {
	return img.ApplyFilter(Sobel5HFilter)
}

// Sobel5V applies a 5×5 vertical Sobel filter.
func (img *Ximg) Sobel5V() *Ximg {
	return img.ApplyFilter(Sobel5VFilter)
}

func (img *Ximg) applySobel(hFunc, vFunc func() *Ximg) *Ximg {
	var sHori, sVert *Ximg
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		sHori = hFunc()
	}()
	go func() {
		defer wg.Done()
		sVert = vFunc()
	}()
	wg.Wait()
	f := func(src *Ximg, dst *Ximg, x int, y int) {
		r1, g1, b1 := sHori.RGBAt(x, y)
		r2, g2, b2 := sVert.RGBAt(x, y)
		color := color.RGBA{
			uint8(math.Sqrt(math.Pow(float64(r1), 2) + math.Pow(float64(r2), 2))),
			uint8(math.Sqrt(math.Pow(float64(g1), 2) + math.Pow(float64(g2), 2))),
			uint8(math.Sqrt(math.Pow(float64(b1), 2) + math.Pow(float64(b2), 2))),
			255,
		}
		dst.Set(x, y, color)
	}
	ximg := img.Apply(1, f, nil)
	ximg = ximg.Normalize(0, 255)
	ximg.isGray = img.isGray

	return ximg
}

// Sobel applies 3×3 horizontal and vertical Sobel filters and combines them by magnitude.
func (img *Ximg) Sobel() *Ximg {
	return img.applySobel(img.SobelH, img.SobelV)
}

// Sobel5 applies 5×5 horizontal and vertical Sobel filters and combines them by magnitude.
func (img *Ximg) Sobel5() *Ximg {
	return img.applySobel(img.Sobel5H, img.Sobel5V)
}

// Sobel2 applies 3×3 Sobel with non-maximum suppression (Canny-like thinning).
func (img *Ximg) Sobel2() *Ximg {
	var (
		sHori, sVert *Ximg
	)
	w, h := img.Size()
	ximg := img.Gray()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		sHori = img.SobelH()
	}()
	go func() {
		defer wg.Done()
		sVert = img.SobelV()
	}()
	wg.Wait()
	f := func(src *Ximg, dst *Ximg, x int, y int) {
		alpha := src.AlphaAt(x, y)
		if x < 1 || x > (w-1) || y < 1 || y > (h-1) {
			dst.Set(x, y, color.RGBA{0, 0, 0, alpha})
			return
		}
		q := uint8(255)
		r := uint8(255)
		hg, _, _ := sHori.RGBAt(x, y)
		vg, _, _ := sVert.RGBAt(x, y)

		a := math.Atan2(float64(vg), float64(hg)) * 180.0 / math.Pi
		if a < 0 {
			a += 180.0
		}

		switch {
		case (a >= 0 && a < 22.5) || (a >= 157.5 && a <= 180):
			q, _, _ = src.RGBAt(x, y+1)
			r, _, _ = src.RGBAt(x, y-1)
		case a >= 22.5 && a < 67.5:
			q, _, _ = src.RGBAt(x+1, y-1)
			r, _, _ = src.RGBAt(x-1, y+1)
		case a >= 67.5 && a < 112.5:
			q, _, _ = src.RGBAt(x+1, y)
			r, _, _ = src.RGBAt(x-1, y)
		case a >= 112.5 && a < 157.5:
			q, _, _ = src.RGBAt(x-1, y-1)
			r, _, _ = src.RGBAt(x+1, y+1)
		}
		v, _, _ := src.RGBAt(x, y)
		if (v >= q) && (v >= r) {
			dst.Set(x, y, color.RGBA{v, v, v, alpha})
		} else {
			dst.Set(x, y, color.RGBA{0, 0, 0, alpha})
		}
	}
	ximg = ximg.Apply(1, f, nil)
	ximg = ximg.Normalize(0, 255)
	ximg.isGray = true

	return ximg
}
