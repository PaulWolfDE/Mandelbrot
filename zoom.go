package main

import (
	"fmt"
	"github.com/PaulWolfDE/Mandelbrot/mandelbrot"
	"image"
	"image/png"
	"math"
	"os"
	"strconv"
)

const START = 1.0
const TRESHOLD = 2.0

func main() {

	width, err := strconv.Atoi(os.Args[1])
	height, err := strconv.Atoi(os.Args[2])
	xoffset, err := strconv.ParseFloat(os.Args[3], 64)
	yoffset, err := strconv.ParseFloat(os.Args[4], 64)
	steps, err := strconv.Atoi(os.Args[5])
	scale, err := strconv.ParseFloat(os.Args[6], 64)
	mandelbrot.Iterations, err = strconv.Atoi(os.Args[7])
	mandelbrot.ColorScheme, err = strconv.Atoi(os.Args[8])

	if err != nil {
		panic(err)
	}
	s := START
	for i := 0; i < steps; i++ {

		fmt.Printf("Frame %d\n", i+1)
		w, h := float64(width), float64(height)
		var xmin, ymin, xmax, ymax float64
		if width > height {
			xmin = -TRESHOLD / float64(height) * float64(width) * 1 / s
			ymin = -TRESHOLD * 1 / s
			xmax = TRESHOLD / float64(height) * float64(width) * 1 / s
			ymax = TRESHOLD * 1 / s
		} else {
			xmin = -TRESHOLD * 1 / s
			ymin = -TRESHOLD / float64(width) * float64(height) * 1 / s
			xmax = TRESHOLD * 1 / s
			ymax = TRESHOLD / float64(width) * float64(height) * 1 / s
		}

		fmt.Printf("%f %f %f %f\n", xmin, ymin, xmax, ymax)
		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			y := float64(py)/h*(ymax-ymin) + ymin + yoffset
			for px := 0; px < width; px++ {
				x := float64(px)/w*(xmax-xmin) + xmin + xoffset
				z := complex(x, y)
				img.Set(px, py, mandelbrot.Mandelbrot(z))
			}
		}
		f, err := os.Create(fmt.Sprintf("./out/frame%06d.png", i+1))
		if err != nil {
			panic(err)
		}
		err = png.Encode(f, img)
		if err != nil {
			return
		}
		s *= math.Exp(math.Log(scale/START) / float64(steps))
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
