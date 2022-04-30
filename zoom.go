package main

import (
    "fmt"
    "image"
    "image/color"
    "image/png"
    "math"
    "math/cmplx"
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
    if err != nil {
        panic(err)
    }
    s := START
    for i := 0; i < steps; i++ {

        fmt.Printf("Frame %d\n", i+1);
        w, h := float64(width), float64(height)
        var xmin, ymin, xmax, ymax float64
        if width > height {
            xmin = float64(-TRESHOLD/float64(height) * float64(width) * 1 / s)
            ymin = float64(-TRESHOLD * 1 / s)
            xmax = float64(TRESHOLD/float64(height) * float64(width) * 1 / s)
            ymax = float64(TRESHOLD * 1 / s)
        } else {
            xmin = float64(-TRESHOLD * 1 / s)
            ymin = float64(-TRESHOLD/float64(width)*float64(height) * 1 / s)
            xmax = float64(TRESHOLD * 1 / s)
            ymax = float64(TRESHOLD/float64(width)*float64(height) * 1 / s)
        }

        fmt.Printf("%f %f %f %f\n", xmin, ymin, xmax, ymax)
        img := image.NewRGBA(image.Rect(0, 0, width, height))
        for py := 0; py < height; py++ {
            y := float64(py)/h*(ymax-ymin) + ymin + yoffset
            for px := 0; px < width; px++ {
                x := float64(px)/w*(xmax-xmin) + xmin + xoffset
                z := complex(x, y)
                img.Set(px, py, mandelbrot(z))
            }
        }
        f, err := os.Create(fmt.Sprintf("./out/frame%d.png", i+1))
        if err != nil {
            panic(err)
        }
        png.Encode(f, img)
        s *= math.Exp(math.Log(scale / START) / float64(steps))
    }
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
