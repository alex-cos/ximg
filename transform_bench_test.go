package ximg_test

import (
	"testing"

	"github.com/alex-cos/ximg"
)

func BenchmarkGrey(b *testing.B) {
	img, err := ximg.Load("testdata/img_01.jpg")
	if err != nil {
		b.Fatal(err)
	}
	for b.Loop() {
		img.Grey()
	}
}

func BenchmarkGray(b *testing.B) {
	img, err := ximg.Load("testdata/img_01.jpg")
	if err != nil {
		b.Fatal(err)
	}
	for b.Loop() {
		img.Gray()
	}
}

func BenchmarkInvert(b *testing.B) {
	img, err := ximg.Load("testdata/img_01.jpg")
	if err != nil {
		b.Fatal(err)
	}
	for b.Loop() {
		img.Invert()
	}
}

func BenchmarkFlipH(b *testing.B) {
	img, err := ximg.Load("testdata/img_01.jpg")
	if err != nil {
		b.Fatal(err)
	}
	for b.Loop() {
		img.FlipH()
	}
}

func BenchmarkFlipV(b *testing.B) {
	img, err := ximg.Load("testdata/img_01.jpg")
	if err != nil {
		b.Fatal(err)
	}
	for b.Loop() {
		img.FlipV()
	}
}

func BenchmarkMaxPool(b *testing.B) {
	img, err := ximg.Load("testdata/img_01.jpg")
	if err != nil {
		b.Fatal(err)
	}
	for b.Loop() {
		img.MaxPool(2)
	}
}
