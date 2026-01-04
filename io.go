package ximg

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// Load decodes an image from a file (JPEG, PNG, GIF).
func Load(filename string) (*Ximg, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	return New(img), nil
}

// Save encodes the image to a file. Supported formats: .jpg/.jpeg, .png, .gif.
func (img *Ximg) Save(filename string, quality int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	switch strings.ToLower(filepath.Ext(filename)) {
	case ".jpg", ".jpeg":
		return jpeg.Encode(file, img, &jpeg.Options{Quality: quality})
	case ".png":
		return png.Encode(file, img)
	case ".gif":
		return gif.Encode(file, img, nil)
	default:
		return fmt.Errorf("unsupported format: %s", filepath.Ext(filename))
	}
}
