package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func mandelbrot(c complex128) color.Color {
	const iter = 200 // Number of iterations
	var z complex128
	for n := uint8(0); n < iter; n++ {
		z = z*z + c
		if cmplx.Abs(z) > 2 {
			return colorScheme(n)
		}
	}
	return color.Black
}

// colorScheme maps iteration count to a color
func colorScheme(n uint8) color.Color {
	// Customize these RGB values for different color schemes
	red := 61 - 0*n
	green := 219 - 5*n
	blue := 217 - 5*n
	return color.RGBA{red, green, blue, 255}
}

func main() {
	// Set the size of the image (width, height)
	const width, height = 1920, 1080

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*2 - 1 // Scale y to -1 to +1
		for px := 0; px < width; px++ {
			x := float64(px)/width*3.5 - 2.5 // Scale x to -2.5 to +1
			color := mandelbrot(complex(x, y))
			img.Set(px, py, color)
		}
	}

	// Create the output file
	file, err := os.Create("fractal.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Encode the image to PNG format
	png.Encode(file, img)
}
