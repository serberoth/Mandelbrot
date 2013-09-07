package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/cmplx"
)

type Mandelbrot struct {
	X, Y, Zoom float64
	Iterations int
}

func (m Mandelbrot) colorAt(i int) color.Color {
	r, g, b := HSVToRGB(1.0-math.Mod(float64(i)/float64(m.Iterations)+0.3333, 1.0), 1.0, 1.0)
	return color.RGBA{r, g, b, 255}
}

func (m Mandelbrot) calculate(c complex128) color.Color {
	limit := 2.0 * m.Zoom
	z := complex(0.0, 0.0)
	for i := 0; i < m.Iterations; i++ {
		z = z*z + c
		if cmplx.Abs(z) > limit {
			return m.colorAt(i)
		}
	}

	return color.Black
}

func (m Mandelbrot) Plot(img *image.RGBA) {
	bounds := img.Bounds()

	ix, iy := LinearSpacing(m.X-m.Zoom, m.X+m.Zoom, bounds.Dx()),
		LinearSpacing(m.Y-m.Zoom, m.Y+m.Zoom, bounds.Dy())

	for y, j := bounds.Min.Y, iy.Start; y < bounds.Max.Y; y, j = y+1, j+iy.Step {
		for x, i := bounds.Min.X, ix.Start; x < bounds.Max.X; x, i = x+1, i+ix.Step {
			img.Set(x, y, m.calculate(complex(i, j)))
		}
	}
}

func mandelbrot(c complex128) string {
	a := []string{"-", "+", "=", "#", "@", "$", "%", "&", "?", "!"}
	z := complex(0.0, 0.0)
	for i := 0; i < 10; i++ {
		z = z*z + c
		if cmplx.Abs(z) > 2.0 {
			// return "+"
			return a[i]
		}
	}

	return "*"
}

func plot(x, y Interval) (s string) {
	for j := range y.Range() {
		for i := range x.Range() {
			s += mandelbrot(complex(i, j))
		}
		s += "\n"
	}

	return
}

func main() {
	s := plot(Interval{-2.05, 1.05, 0.03}, Interval{-1.2, 1.2, 0.05})

	fmt.Printf("%s", s)

	m := Mandelbrot{-0.5, 0.0, 2.0, 256}

	img := image.NewRGBA(image.Rect(0, 0, 1440, 900))

	m.Plot(img)

	WritePNG(img, "output.png")
}
