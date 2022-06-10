package main

import (
	"github.com/PaulWolfDE/Mandelbrot/mandelbrot"
	"image"
	"image/png"
	"os"
	"strconv"
)

func main() {

	width, err := strconv.Atoi(os.Args[1])
	height, err := strconv.Atoi(os.Args[2])
	scale, err := strconv.ParseFloat(os.Args[3], 64)
	xoffset, err := strconv.ParseFloat(os.Args[4], 64)
	yoffset, err := strconv.ParseFloat(os.Args[5], 64)
	mandelbrot.Iterations, err = strconv.Atoi(os.Args[6])
	mandelbrot.ColorScheme, err = strconv.Atoi(os.Args[7])

	if err != nil {
		panic(err)
	}

	xmin, ymin, xmax, ymax := -2*1/scale, -2*1/scale, 2*1/scale, 2*1/scale

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin + yoffset
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin + xoffset
			z := complex(x, y)
			img.Set(px, py, mandelbrot.Mandelbrot(z))
		}
	}
	err = png.Encode(os.Stdout, img)
	if err != nil {
		panic(err)
	}
}
