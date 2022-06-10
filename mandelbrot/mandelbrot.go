package mandelbrot

import (
	"image/color"
	"math/cmplx"
)

var (
	ColorScheme = 0
	Iterations  = 256
)

func Mandelbrot(z complex128) color.RGBA {

	const contrast = 20

	var v complex128
	for n := 0; n < Iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {

			// 0 -> black and white
			if ColorScheme == 0 {
				return color.RGBA{R: uint8(contrast * n), G: uint8(contrast * n), B: uint8(contrast * n), A: 255}
			}
			// 1 -> rgb scheme
			if ColorScheme == 1 {
				if int(n/256)%3 == 0 {
					if n/256 < 1 {
						return color.RGBA{R: uint8(n), A: 255}
					} else {
						return color.RGBA{R: uint8(n), B: 255 - uint8(n-768*int(n/256)), A: 255}
					}
				}
				if int(n/256)%3 == 1 {
					return color.RGBA{R: 255 - uint8(n-256*int(n/256)), G: uint8(n % 256), A: 255}
				}
				if int(n/256)%3 == 2 {
					return color.RGBA{G: 255 - uint8(n-512*int(n/256)), B: uint8(n % 256), A: 255}
				}
			}
			// 2 -> rgb advanced
			if ColorScheme == 2 || ColorScheme == 3 {
				if int(contrast*n/256)%3 == 0 {
					if contrast*n/256 < 1 {
						return color.RGBA{R: uint8(contrast * n), A: 255}
					} else {
						return color.RGBA{R: uint8(contrast * n), B: 255 - uint8(contrast*n-768*int(contrast*n/256)), A: 255}
					}
				}
				if int(contrast*n/256)%3 == 1 {
					return color.RGBA{R: 255 - uint8(contrast*n-256*int(contrast*n/256)), G: uint8(contrast * n % 256), A: 255}
				}
				if int(contrast*n/256)%3 == 2 {
					return color.RGBA{G: 255 - uint8(contrast*n-512*int(contrast*n/256)), B: uint8(contrast * n % 256), A: 255}
				}
			}

			// HSL
			if ColorScheme == 4 {
				c := HSL{float64((contrast/4*n)%360) / 360, .5, .5}.ToRGB()
				return color.RGBA{uint8(c.R * 255), uint8(c.G * 255), uint8(c.B * 255), 255}
			}
			// return color.RGBA{uint8(255 - contrast*1/2*n), uint8(255 - contrast*1/3*n), uint8(255 - contrast*n), 255}
			return color.RGBA{}
		}
	}
	if ColorScheme == 0 || ColorScheme == 3 {
		return color.RGBA{R: 255, G: 255, B: 255, A: 255}
	}
	return color.RGBA{A: 255}
}
