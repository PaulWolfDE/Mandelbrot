package main

import (
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "os"
    "strconv"
)

func main() {

    width, err := strconv.Atoi(os.Args[1])
    height, err := strconv.Atoi(os.Args[2])
    scale, err := strconv.ParseFloat(os.Args[3], 64)
    xoffset, err := strconv.ParseFloat(os.Args[4], 64)
    yoffset, err := strconv.ParseFloat(os.Args[5], 64)
    if err != nil {
        panic(err)
    }

    xmin, ymin, xmax, ymax := float64(-2*1/scale), float64(-2*1/scale), float64(2*1/scale), float64(2*1/scale)
    
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    
    for py := 0; py < height; py++ {
        y := float64(py)/float64(height)*(ymax-ymin) + ymin + yoffset
        for px := 0; px < width; px++ {
            x := float64(px)/float64(width)*(xmax-xmin) + xmin + xoffset
            z := complex(x, y)
            img.Set(px, py, mandelbrot(z))
        }
    }
    png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.RGBA {
    
    const iterations = 200
    const contrast = 20

    var v complex128
    for n := uint8(0); n < iterations; n++ {
        v = v*v + z
        if cmplx.Abs(v) > 2 {
            // Values should be edited later
            return color.RGBA{255 - contrast*1/2*n, 255 - contrast*1/3*n, 255 - contrast*n, 255}
        }
    }
    return color.RGBA{0, 0, 0, 255}
}