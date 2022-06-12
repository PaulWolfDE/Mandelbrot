# Mandelbrot

## Install

```sh
$ git clone https://github.com/PaulWolfDE/Mandelbrot.git
$ go get -v github.com/gerow/go-color
```

## Use

### Rendering image
```sh
$ go run image.go [image width] [image height] [scale factor] [offset x] [offset y] [formula iterations] [color scheme] > mandelbrot.png
```

### Rendering zoom with multiple images
```sh
$ mkdir out
$ go run zoom.go [image width] [image height] [x offset] [y offset] [zoom iterations] [scale factor] [formula iterations] [color scheme]
```

### Make video from zoom files

[<img src="https://img.shields.io/badge/Rquirement-FFMPEG-blue.svg">](https://ffmpeg.org/)

```sh
$ mkdir out
$ go run zoom.go 1920 1080 -.68 -.3 300 305 600 5
$ ffmpeg -framerate 24 -pattern_type glob -i '*.png' video.mp4
$ open video.mp4
```

## Credits

- Initial code derived from [The Go Programming Language](https://github.com/adonovan/gopl.io/) by Alan A. A. Donovan and Brian W. Kernighan
- Used library [HSL Color](https://github.com/gerow/go-color) by Gerow

## License

[<img src="https://img.shields.io/badge/License-GPL 3-important.svg">](https://gnu.org/licenses/gpl-3.0.html)
(Derived from gopl's [CC BY-NC-SA 4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/))