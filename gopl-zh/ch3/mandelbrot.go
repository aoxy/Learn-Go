// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		xin, err := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
		if err != nil {
			xin = 2.0
		}
		yin, err := strconv.ParseFloat(r.URL.Query().Get("y"), 64)
		if err != nil {
			yin = 2.0
		}
		zoom, err := strconv.ParseFloat(r.URL.Query().Get("zoom"), 64)
		if err != nil {
			zoom = 100.0
		}
		var (
			xmin, ymin, xmax, ymax = -xin, -xin, +yin, +yin
			width, height          = 10.24 * zoom, 10.24 * zoom
		)

		img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
		for py := 0; py < int(height); py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < int(width); px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.
				img.Set(px, py, mandelbrot(z))
			}
		}
		png.Encode(w, img) // NOTE: ignoring errors
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
