package main

import (
	"fmt"
	"image"
)

// The main method for the Mandelbrot plotting application.  Thie function
// reads the 'plot.json' file and plots the Mandelbrot set defined within
// then outputs the resulting plot to the 'output.png' file.  Additionally
// this function plots a textual version of the Mandelbrot set and writes
// that out to the console.
func main() {
	s := plot(Interval{-2.5, 1.0, 0.03}, Interval{-1.5, 1.5, 0.05})
	fmt.Printf("%s", s)

	// m := Mandelbrot{-0.5, 0.0, 2.0, 256}
	m, err := ReadMandelbrot("plot.json")
	if err != nil {
		fmt.Printf("Failed reading 'plot.json': %v\n", err)
		return
	}

	img := image.NewRGBA(image.Rect(0, 0, 1440, 900))

	m.Plot(img)

	WritePNG(img, "output.png")
}
