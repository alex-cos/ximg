package ximg

import "errors"

var (
	ErrInputImageIsNil = errors.New("given input image is nil")

	ErrInputImagesNotSameSize = errors.New("input images should have the same size")
)
