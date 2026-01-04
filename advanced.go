package ximg

import (
	"sync"
)

func Convolution1(img *Ximg) []*Ximg {
	out := img.
		Resize(1024, 0).
		Contrast(0.025).
		Blur().
		Gray().
		Normalize(5, 255).
		Gaussian7().
		Resize(128, 0)

	filter1 := out.Sobel().MaxPool(2)
	filter1 = filter1.Contrast(0.025)
	filter3 := out.MaxPool(2)
	filter3 = filter3.Contrast(0.5)

	return []*Ximg{filter1, filter3}
}

func Convolution2(img *Ximg) []*Ximg {
	out := img.
		Resize(1024, 0).
		Contrast(0.025).
		Blur().
		Gray().
		Normalize(5, 255).
		Gaussian7().
		Resize(128, 0)

	filter1 := out.Sobel().MaxPool(2)
	filter1 = filter1.Contrast(0.025)
	filter2 := out.EdgeDetect().MaxPool(2)
	filter2 = filter2.Contrast(0.025)
	filter3 := out.MaxPool(2)
	filter3 = filter3.Contrast(0.5)

	return []*Ximg{filter1, filter2, filter3}
}

func Convolution3(img *Ximg) []*Ximg {
	var (
		f1, f2 *Ximg
	)

	out := img.
		Resize(1024, 0).
		Gaussian5().
		Resize(128, 0).
		Normalize(5, 255).
		Gray()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		f1 = out.Sobel().MaxPool(2)
	}()
	go func() {
		defer wg.Done()
		f2 = out.Contrast(0.5).MaxPool(2)
	}()
	wg.Wait()

	return []*Ximg{f1, f2}
}
