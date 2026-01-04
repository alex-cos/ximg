package ximg

import "math"

// NormalizeFloat rescales data to [0, 1] using min-max normalization.
func NormalizeFloat[F Float](data []F) []F {
	if len(data) == 0 {
		return nil
	}

	vmin, vmax := data[0], data[0]
	for _, v := range data {
		if v < vmin {
			vmin = v
		}
		if v > vmax {
			vmax = v
		}
	}
	if vmax == vmin {
		out := make([]F, len(data))
		copy(out, data)
		return out
	}

	scale := 1 / (vmax - vmin)
	out := make([]F, len(data))
	for i, v := range data {
		out[i] = (v - vmin) * scale
	}

	return out
}

// NormalizeComplex scales a 1D complex64 slice by the maximum magnitude.
func NormalizeComplex(data []complex64) []complex64 {
	if len(data) == 0 {
		return nil
	}

	var maxMag float64
	for _, z := range data {
		mag := math.Sqrt(float64(real(z)*real(z) + imag(z)*imag(z)))
		if mag > maxMag {
			maxMag = mag
		}
	}
	if maxMag == 0 {
		out := make([]complex64, len(data))
		copy(out, data)
		return out
	}

	scale := float32(1 / maxMag)
	out := make([]complex64, len(data))
	for i, z := range data {
		out[i] = complex(real(z)*scale, imag(z)*scale)
	}

	return out
}

// NormalizeComplex2D scales a 2D complex64 slice by the maximum magnitude.
func NormalizeComplex2D(data [][]complex64) [][]complex64 {
	var maxMag float64

	if len(data) == 0 {
		return nil
	}
	for _, row := range data {
		for _, z := range row {
			mag := math.Sqrt(float64(real(z)*real(z) + imag(z)*imag(z)))
			if mag > maxMag {
				maxMag = mag
			}
		}
	}
	if maxMag == 0 {
		out := make([][]complex64, len(data))
		for i, row := range data {
			out[i] = make([]complex64, len(row))
			copy(out[i], row)
		}
		return out
	}

	scale := float32(1 / maxMag)
	out := make([][]complex64, len(data))
	for i, row := range data {
		out[i] = make([]complex64, len(row))
		for j, z := range row {
			out[i][j] = complex(real(z)*scale, imag(z)*scale)
		}
	}

	return out
}
