package main

import (
	"encoding/json"
	"image"
	"image/color"
	"io/ioutil"
	"math"
	"math/cmplx"
)

// The Mandelbrot data structure, (X, Y) the origin of the plot
// Zoom is the magnification factor.  Iterations determines the
// number of iterations prior to stopping the plot calculations
// and colors determines the number of colors in the color
// gradient.
type Mandelbrot struct {
	X, Y, Zoom float64
	Iterations, Colors int
}

// Read the Mandelbrot instance from the provided filename
// unmarshaling the structure from JSON.  Return the Mandelbrot
// and any error that may have occurred while reading the file.
func ReadMandelbrot(name string) (Mandelbrot, error) {
	var m Mandelbrot

	f, e := ioutil.ReadFile(name)
	if e != nil {
		return m, e
	}

	if e = json.Unmarshal(f, &m); e != nil {
		return m, e
	}

	return m, nil
}

// Determine the color at the given iteration converting the HSV color to the RGB colorspace
func (m Mandelbrot) colorAt(i int) color.Color {
	r, g, b := HSVToRGB(1.0-math.Mod(float64(i % m.Colors)/float64(m.Colors)+0.3333, 1.0), 1.0, 1.0)
	return color.RGBA{r, g, b, 255}
}

// Calculate the color at the provided complex point in the plot
func (m Mandelbrot) calculate(c complex128) color.Color {
	z := complex(0.0, 0.0)
	for i := 0; i < m.Iterations; i++ {
		z = z*z + c
		if cmplx.Abs(z) > 4.0 {
			return m.colorAt(i)
		}
	}

	return color.Black
}

// Plot the Mandelbrot set to the provided RGBA image instance
func (m Mandelbrot) Plot(img *image.RGBA) {
	bounds := img.Bounds()

	incrament := (4.0 / m.Zoom) / math.Max(float64(bounds.Dx()), float64(bounds.Dy()))
	ix, iy := -((incrament * float64(bounds.Dx())) / 2.0) + m.X, -((incrament * float64(bounds.Dy())) / 2.0) + m.Y

	for y, j := bounds.Min.Y, iy; y < bounds.Max.Y; y, j = y + 1, j + incrament {
		for x, i := bounds.Min.X, ix; x < bounds.Max.X; x, i = x + 1, i + incrament {
			img.Set(x, y, m.calculate(complex(i, j)))
		}
	}
}

// Calculate the point in the mandelbrot set returning a character value
// within the set [ -, +, =, #, @, $, %, &, ?, ! ].
func mandelbrot(c complex128) string {
	a := []string{"-", "+", "=", "#", "@", "$", "%", "&", "?", "!"}
	z := complex(0.0, 0.0)
	for i := 0; i < 50; i++ {
		z = z*z + c
		if cmplx.Abs(z) > 2.0 {
			// return "+"
			return a[i % len(a)]
		}
	}

	return "*"
}

// Plot the standard Mendelbrot plot to a string
func plot(x, y Interval) (s string) {
	for j := range y.Range() {
		for i := range x.Range() {
			s += mandelbrot(complex(i, j))
		}
		s += "\n"
	}

	return
}

