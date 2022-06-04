package main

import (
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "os"
    "strconv"
)

var iterations, colorScheme int

func main() {

    width, err := strconv.Atoi(os.Args[1])
    height, err := strconv.Atoi(os.Args[2])
    scale, err := strconv.ParseFloat(os.Args[3], 64)
    xoffset, err := strconv.ParseFloat(os.Args[4], 64)
    yoffset, err := strconv.ParseFloat(os.Args[5], 64)
    iterations, err = strconv.Atoi(os.Args[6])
    colorScheme, err = strconv.Atoi(os.Args[7])

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
    
    const contrast = 10

    var v complex128
    for n := 0; n < iterations; n++ {
        v = v*v + z
        if cmplx.Abs(v) > 2 {

            // 0 -> black and white
            if (colorScheme == 0) {
                return color.RGBA{uint8(contrast*n), uint8(contrast*n), uint8(contrast*n), 255}
            }
            // 1 -> rgb scheme
            if (colorScheme == 1) {
                if (int(n/256) % 3 == 0) {
                    if (n/256 < 1) {
                        return color.RGBA{uint8(n), 0, 0, 255}
                    } else {
                        return color.RGBA{uint8(n), 0, 255 - uint8(n - 768*int(n/256)), 255}
                    }
                }
                if (int(n/256) % 3 == 1) {
                    return color.RGBA{255-uint8(n - 256*int(n/256)), uint8(n%256), 0, 255}
                }
                if (int(n/256) % 3 == 2) {
                    return color.RGBA{0, 255-uint8(n - 512*int(n/256)), uint8(n%256), 255}
                }
            }
            // 2 -> rgb advanced
            if (colorScheme == 2) {
                if (int(contrast*n/256) % 3 == 0) {
                    if (contrast*n/256 < 1) {
                        return color.RGBA{uint8(contrast*n), 0, 0, 255}
                    } else {
                        return color.RGBA{uint8(contrast*n), 0, 255 - uint8(contrast*n - 768*int(contrast*n/256)), 255}
                    }
                }
                if (int(contrast*n/256) % 3 == 1) {
                    return color.RGBA{255-uint8(contrast*n - 256*int(contrast*n/256)), uint8(contrast*n%256), 0, 255}
                }
                if (int(contrast*n/256) % 3 == 2) {
                    return color.RGBA{0, 255-uint8(contrast*n - 512*int(contrast*n/256)), uint8(contrast*n%256), 255}
                }   
            }
            // return color.RGBA{uint8(255 - contrast*1/2*n), uint8(255 - contrast*1/3*n), uint8(255 - contrast*n), 255}
            return color.RGBA{0, 0, 0, 0}
        }
    }
    if (colorScheme == 0) {
        return color.RGBA{255, 255, 255, 255}
    }
    return color.RGBA{0, 0, 0, 255}
}