package ximg

import "image"

// Ximg wraps an image.RGBA with additional metadata.
type Ximg struct {
	image.RGBA
	isGray bool
}

// ApplyFunc is the function signature used by Apply.
type ApplyFunc func(src *Ximg, dst *Ximg, x int, y int)

// Options provides configuration for Apply.
type Options struct {
	Workers int // 0 = auto (GOMAXPROCS)
}

// Float constraint for float32 and float64.
type Float interface {
	float32 | float64
}

// Integer constraint for signed integer types.
type Integer interface {
	int | int8 | int16 | int32 | int64
}

// Complex constraint for complex64 and complex128.
type Complex interface {
	complex64 | complex128
}

// nolint: iface
// Number constraint covering all numeric types used in the library.
type Number interface {
	Integer | Float | Complex
}
