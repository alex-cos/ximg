package ximg

import (
	"runtime"
	"sync"
)

// Apply applies a function to every pixel (or every stride-th pixel) in parallel.
func (img *Ximg) Apply(stride int, applyFunc ApplyFunc, opts *Options) *Ximg {
	w, h := img.Size()
	s := stride
	if s <= 0 {
		s = 1
	}
	ximg := NewRGBA(w/s, h/s)

	nbWorkers := runtime.GOMAXPROCS(0)
	if opts != nil && opts.Workers > 0 {
		nbWorkers = opts.Workers
	}

	n := h / nbWorkers
	if n <= 0 {
		n = 1
	}
	wg := sync.WaitGroup{}
	wg.Add(nbWorkers)
	for i := range nbWorkers {
		go func(i int) {
			defer wg.Done()
			end := (i + 1) * n
			if i+1 == nbWorkers {
				end = h
			}
			for y := i * n; y < end; y += s {
				for x := 0; x < w; x += s {
					applyFunc(img, ximg, x, y)
				}
			}
		}(i)
	}
	wg.Wait()

	return ximg
}
