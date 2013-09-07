package main

import (
	"fmt"
	// "image"
	"image/color"
	// "image/png"
	"math"
	"math/cmplx"
	// "os"
)

func HSVToRGB(h, s, v float64) (uint8, uint8, uint8) {
  r, g, b := v, v, v
  if (s != 0.0) {
    h = math.Mod(h * 6.0, 6.0)
    s = s * v
    v = s * (1.0 - math.Abs(math.Mod(h, 2.0) - 1.0))

    f := [][]float64{
      { s, v, 0.0 },
      { v, s, 0.0 },
      { 0.0, s, v },
      { 0.0, v, s },
      { v, 0.0, s },
      { s, 0.0, v },
    }

    r, g, b = f[int32(math.Floor(h))][0], f[int32(math.Floor(h))][1], f[int32(math.Floor(h))][2]
  }

  return uint8(r * 255.0), uint8(g * 255.0), uint8(b * 255.0)
}

func getColor(i, s int32) color.Color {
  r, g, b := HSVToRGB(math.Mod(1.0 - (float64(i) / float64(s) + 0.3333), 1.0), 1.0, 1.0)
  return color.RGBA{ r, g, b, 1.0 }
}

func mandelbrot(c complex128) string {
	a := []string{ "-", "+", "=", "#", "@", "$", "%", "&", "?", "!" }
	z := complex(0.0, 0.0)
	for i := 0; i < 10; i++ {
		z = z * z + c
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
}
