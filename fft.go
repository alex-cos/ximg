package ximg

import "math"

// FFT computes the 1D FFT (Cooley-Tukey radix-2) with zero-padding to next power of 2.
func FFT[F Float](input []F) []complex64 {
	n := nextPow2(len(input))
	a := make([]complex64, n)
	for i, v := range input {
		a[i] = complex(float32(v), 0)
	}

	fftInplace(a)
	return a
}

// FFT2D computes the 2D FFT by applying 1D FFT on rows then columns.
// Data layout: data[x][y] (width × height).
// Output is padded to the next power of 2 in each dimension.
func FFT2D[F Float](data [][]F) [][]complex64 {
	if len(data) == 0 {
		return nil
	}
	w := len(data)
	h := len(data[0])
	if h == 0 {
		return nil
	}

	pw := nextPow2(w)
	ph := nextPow2(h)

	out := make([][]complex64, pw)
	for x := range pw {
		out[x] = make([]complex64, ph)
		if x < w {
			for y := range ph {
				if y < h {
					out[x][y] = complex(float32(data[x][y]), 0)
				}
			}
		}
	}

	row := make([]complex64, pw)
	for y := range ph {
		for x := range pw {
			row[x] = out[x][y]
		}
		fftInplace(row)
		for x := range pw {
			out[x][y] = row[x]
		}
	}

	for x := range pw {
		fftInplace(out[x])
	}

	return out
}

// FFTCenter shifts the DC component to the center of the spectrum (fftshift).
func FFTCenter[C Complex](data [][]C) [][]C {
	if len(data) == 0 {
		return nil
	}
	w := len(data)
	h := len(data[0])
	if h == 0 {
		return nil
	}

	out := make([][]C, w)
	for x := range w {
		out[x] = make([]C, h)
	}

	cx := (w + 1) / 2
	cy := (h + 1) / 2

	for x := range w {
		for y := range h {
			nx := (x + cx) % w
			ny := (y + cy) % h
			out[nx][ny] = data[x][y]
		}
	}

	return out
}

// LogMagnitude computes log(1 + |z|) for each complex64 value.
func LogMagnitude(data []complex64) []float64 {
	out := make([]float64, len(data))
	for i, z := range data {
		mag := math.Sqrt(float64(real(z)*real(z) + imag(z)*imag(z)))
		out[i] = math.Log(1 + mag)
	}
	return out
}

func fftInplace(a []complex64) {
	n := len(a)
	if n <= 1 {
		return
	}
	bitReverse(a)
	for length := 2; length <= n; length <<= 1 {
		angle := -2 * math.Pi / float64(length)
		for i := 0; i < n; i += length {
			half := length >> 1
			for j := range half {
				w := complexFromPolar(angle * float64(j))
				u := a[i+j]
				v := a[i+j+half] * w
				a[i+j] = u + v
				a[i+j+half] = u - v
			}
		}
	}
}
