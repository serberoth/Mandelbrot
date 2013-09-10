package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"math"
	"math/cmplx"
)

type Mandelbrot struct {
	X, Y, Zoom float64
	Iterations, Colors int
}

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

func (m Mandelbrot) colorAt(i int) color.Color {
	r, g, b := HSVToRGB(1.0-math.Mod(float64(i % m.Colors)/float64(m.Colors)+0.3333, 1.0), 1.0, 1.0)
	return color.RGBA{r, g, b, 255}
}

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
