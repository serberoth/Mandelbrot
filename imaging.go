package main

import (
	"image"
	"image/png"
	"math"
	"os"
)

func HSVToRGB(h, s, v float64) (uint8, uint8, uint8) {
	r, g, b := v, v, v
	if s != 0.0 {
		h = math.Mod(h*6.0, 6.0)
		s = s * v
		v = s * (1.0 - math.Abs(math.Mod(h, 2.0)-1.0))

		f := [][]float64{
			{s, v, 0.0},
			{v, s, 0.0},
			{0.0, s, v},
			{0.0, v, s},
			{v, 0.0, s},
			{s, 0.0, v},
		}

		r, g, b = f[int(math.Floor(h))][0], f[int(math.Floor(h))][1], f[int(math.Floor(h))][2]
	}

	return uint8(r * 255.0), uint8(g * 255.0), uint8(b * 255.0)
}

func WritePNG(img image.Image, filename string) error {
	writer, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer writer.Close()

	return png.Encode(writer, img)
}
