# Mandelbrot

```sh
$ go run mandelbrot.go [image width] [image height] [scale factor] [offset x] [offset y] [formula iterations] [color scheme] > mandelbrot.png
```

```sh
$ go run zoom.go [image width] [image height] [x offset] [y offset] [zoom iterations] [scale factor] [formula iterations] [color scheme]
```

## Make video from zoom files

### Requires
- [ffmpeg](https://ffmpeg.org/)

```sh
$ go run zoom.go 1920 1080 -.68 -.3 300 305 600 5
$ mkdir out
$ ffmpeg -framerate 24 -pattern_type glob -i '*.png' video.mp4
$ open video.mp4
```