# ximg

[![Go Version](https://img.shields.io/badge/Go-1.24%2B-blue)](https://go.dev/)
[![Test Status](https://github.com/alex-cos/ximg/actions/workflows/test.yml/badge.svg)](https://github.com/alex-cos/ximg/actions/workflows/test.yml)
[![Codecov](https://codecov.io/gh/alex-cos/ximg/branch/main/graph/badge.svg)](https://codecov.io/gh/alex-cos/ximg)
[![Lint Status](https://github.com/alex-cos/ximg/actions/workflows/lint.yml/badge.svg)](https://github.com/alex-cos/ximg/actions/workflows/lint.yml)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/alex-cos/ximg)](https://goreportcard.com/report/github.com/alex-cos/ximg)

Image processing library for Go, built on top of `image.RGBA`.

![ximg](/ximg.png)

## Features

- **Filters**: blur, sharpen, Gaussian, edge detection, Sobel operators (3×3 and 5×5)
- **Transforms**: resize, flip, invert, grayscale, channel extraction, brightness, contrast, hue, color-to-alpha
- **Pooling**: max pooling (useful for CNN-like pipelines)
- **FFT**: 1D/2D FFT (Cooley-Tukey radix-2) with centering and log magnitude
- **Conversion**: to/from `float32`, `float64`, `complex64`, `complex128` (1D, 2D)
- **Normalization**: min-max float scaling, max-magnitude complex scaling
- **Parallel processing**: `Apply()` distributes pixel work across CPU cores
- **IO**: load JPEG/PNG/GIF, save to JPEG/PNG/GIF

## Installation

```bash
go get github.com/alex-cos/ximg
```

## Quick start

```go
package main

import (
    "fmt"
    "github.com/alex-cos/ximg"
)

func main() {
    img, err := ximg.Load("input.jpg")
    if err != nil {
        panic(err)
    }

    // Chain transformations
    result := img.
        Resize(512, 0).       // resize keeping aspect ratio
        Gray().               // convert to grayscale
        Contrast(0.05).       // adjust contrast (sigmoid)
        Sharpen().            // sharpen
        Normalize(0, 255)     // stretch values to full range

    result.Save("output.jpg", 90)
    fmt.Println("done")
}
```

## API overview

### Constructors

| Function | Description |
| --- | --- |
| `New(image.Image) *Ximg` | Wrap an existing image |
| `NewRGBA(width, height int) *Ximg` | Create a blank RGBA image |
| `NewFromRGB(red, green, blue *Ximg) (*Ximg, error)` | Combine three single-channel images into an RGB image |

### IO

| Function | Description |
| --- | --- |
| `Load(filename string) (*Ximg, error)` | Load JPEG, PNG, or GIF |
| `Save(filename string, quality int) error` | Save (format from extension: `.jpg`/`.jpeg`, `.png`, `.gif`) |

### Pixel access

| Method | Description |
| --- | --- |
| `Clone() *Ximg` | Return a deep copy of the image |
| `Size() (int, int)` | Get width and height |
| `RGBAt(x, y int) (uint8, uint8, uint8)` | Get RGB at pixel |
| `RGBAAt(x, y int) (uint8, uint8, uint8, uint8)` | Get RGBA at pixel |
| `AlphaAt(x, y int) uint8` | Get alpha at pixel |
| `IsGray() bool` | Is the image currently grayscale? |
| `ToFloat32() []float32` | Flatten to grayscale float32 [0–1] |
| `ToFloat64() []float64` | Flatten to grayscale float64 [0–1] |
| `ToFloat32_2D() [][]float32` | Grayscale float32 as 2D slice |
| `ToFloat64_2D() [][]float64` | Grayscale float64 as 2D slice |
| `ToComplex() []complex64` | Flatten to complex64 |
| `ToComplex_2D() [][]complex64` | Complex64 as 2D slice |
| `ToComplex128() []complex128` | Flatten to complex128 |
| `ToComplex128_2D() [][]complex128` | Complex128 as 2D slice |

### Filters

| Method | Description | Kernel |
| --- | --- | --- |
| `Blur()` | Box blur | 3×3 |
| `Sharpen()` | Sharpening | 3×3 |
| `Sharpen2()` | Stronger sharpening | 3×3 |
| `Gaussian5()` | Gaussian blur | 5×5 |
| `Gaussian7()` | Gaussian blur | 7×7 |
| `EdgeDetect()` | Edge detection | 3×3 |
| `EdgeDetectH()` | Horizontal edges | 3×3 |
| `EdgeDetectV()` | Vertical edges | 3×3 |
| `EdgeDetectD1()` | Diagonal edges (/) | 3×3 |
| `EdgeDetectD2()` | Diagonal edges (\) | 3×3 |
| `SobelH()` | Sobel horizontal | 3×3 |
| `SobelV()` | Sobel vertical | 3×3 |
| `Sobel()` | Sobel magnitude | 3×3 |
| `Sobel5H()` | Sobel horizontal | 5×5 |
| `Sobel5V()` | Sobel vertical | 5×5 |
| `Sobel5()` | Sobel magnitude | 5×5 |
| `Sobel2()` | Sobel with non-max suppression | 3×3 |
| `ApplyFilter(kernel [][]float64)` | Apply arbitrary kernel | any |

### Transforms

| Method | Description |
| --- | --- |
| `Resize(width, height int)` | Resize (set one to 0 for auto) |
| `Gray()` | Grayscale (ITU-R BT.601 luminance) |
| `Grey()` | Grayscale (equal weights) |
| `Red()` / `Green()` / `Blue()` | Extract single channel |
| `Invert()` | Invert colors |
| `FlipH()` / `FlipV()` | Horizontal / vertical flip |
| `Normalize(min, max uint8)` | Contrast stretch |
| `Contrast(a float64)` | Sigmoid contrast adjustment |
| `Brightness(delta int)` | Add delta to each channel |
| `Hue(shift float64)` | Shift hue by angle in degrees |
| `MaxPool(size int)` | Max pooling downsampling |
| `Split(width, shift int)` | Split into overlapping tiles |
| `Merge(other *Ximg)` | Blend two images |
| `Fuzion(x, y int, with *Ximg)` | Overlay `with` at position (x, y) with alpha compositing |
| `ColorToAlpha(col color.RGBA, tolerance float32)` | Set matching color to transparent |
| `RemoveAlpha()` | Set alpha to 255 |
| `AlphaToColor(bg color.RGBA)` | Replace transparency by compositing over a solid background |
| `AlphaToColor(bg color.RGBA)` | replaces transparency with a solid background |

### Conversion (from image to data)

| Function | Description |
| --- | --- |
| `ToFloat[F](img *Ximg) []F` | Generic grayscale to flat float |
| `ToFloat2D[F](img *Ximg) [][]F` | Generic grayscale to 2D float |
| `ToComplex(img *Ximg) []complex64` | Grayscale to flat complex64 |
| `ToComplex_2D(img *Ximg) [][]complex64` | Grayscale to 2D complex64 |
| `ToComplex128(img *Ximg) []complex128` | Grayscale to flat complex128 |
| `ToComplex128_2D(img *Ximg) [][]complex128` | Grayscale to 2D complex128 |

### Conversion (from data to image)

| Function | Description |
| --- | --- |
| `FromFloat[F](data []F, w, h int) *Ximg` | Flat float to grayscale image |
| `FromFloat2D[F](data [][]F) *Ximg` | 2D float to grayscale image |
| `FromComplex(data []complex64, w, h int) *Ximg` | Flat complex64 magnitude to grayscale |
| `FromComplex2D(data [][]complex64) *Ximg` | 2D complex64 magnitude to grayscale |
| `FromComplex128(data []complex128, w, h int) *Ximg` | Flat complex128 magnitude to grayscale |
| `FromComplex128_2D(data [][]complex128) *Ximg` | 2D complex128 magnitude to grayscale |

### FFT

| Function | Description |
| --- | --- |
| `FFT[F](input []F) []complex64` | 1D FFT (Cooley-Tukey radix-2, zero-padded to next power of 2) |
| `FFT2D[F](data [][]F) [][]complex64` | 2D FFT (row then column, padded per axis) |
| `FFTCenter[C](data [][]C) [][]C` | Shift DC to center (fftshift) |
| `LogMagnitude(data []complex64) []float64` | Compute log(1 + \|z\|) for spectrum display |

### Normalization

| Function | Description |
| --- | --- |
| `NormalizeFloat[F](data []F) []F` | Min-max scale to [0, 1] |
| `NormalizeComplex(data []complex64) []complex64` | Scale by max magnitude to [0, 1] |
| `NormalizeComplex2D(data [][]complex64) [][]complex64` | 2D scale by max magnitude |

### Utilities

| Function | Description |
| --- | --- |
| `ConcatV(width int, images ...*Ximg) *Ximg` | Stack images vertically at given width |
| `ConcatH(height int, images ...*Ximg) *Ximg` | Stack images horizontally at given height |

### Advanced

| Type / Function | Description |
| --- | --- |
| `ApplyFunc` | `func(src, dst *Ximg, x, y int)` |
| `Options` | `Workers int` — parallel worker count |
| `Apply(stride int, fn ApplyFunc, opts *Options)` | Run a pixel function in parallel |

### Filter kernels (exported variables)

| Variable | Description |
| --- | --- |
| `BlurFilter` | 3×3 averaging kernel |
| `SharpenFilter` | 3×3 sharpening kernel |
| `Sharpen2Filter` | 3×3 stronger sharpening kernel |
| `LaplacianFilter` | 3×3 Laplacian kernel |
| `Gaussian5Filter` | 5×5 Gaussian kernel |
| `Gaussian7Filter` | 7×7 Gaussian kernel |
| `EdgeDetectFilter` | 3×3 Laplacian-like kernel |
| `EdgeDetectHFilter` | 3×3 horizontal edge kernel |
| `EdgeDetectVFilter` | 3×3 vertical edge kernel |
| `EdgeDetectD1Filter` | 3×3 diagonal (/) edge kernel |
| `EdgeDetectD2Filter` | 3×3 diagonal (\) edge kernel |
| `SobelHFilter` | 3×3 horizontal Sobel kernel |
| `SobelVFilter` | 3×3 vertical Sobel kernel |
| `Sobel5HFilter` | 5×5 horizontal Sobel kernel |
| `Sobel5VFilter` | 5×5 vertical Sobel kernel |
