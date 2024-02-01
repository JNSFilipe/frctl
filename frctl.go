package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
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
	return color.RGBA{22, 22, 22, 255}
}

// colorScheme maps iteration count to a color
func colorScheme(n uint8) color.Color {
	// Customize these RGB values for different color schemes
	red := 61 - 0*n
	green := 219 - 10*n
	blue := 217 - 10*n
	return color.RGBA{red, green, blue, 255}
}

func main() {
	// Set the size of the image (width, height)
	const width, height = 1920, 1080

	// Fractal parameters
	xOffset, yOffset := -1.2, -0.65 // Initial position
	scale := 0.5                    // Scale factor
	rotation := -29.33              // Rotation angle in degrees

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	sin, cos := math.Sincos(rotation * math.Pi / 180)

	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			x := (float64(px)-float64(width)/2)*(scale*0.001) + xOffset
			y := (float64(py)-float64(height)/2)*(scale*0.001) + yOffset
			xNew := cos*x - sin*y
			yNew := sin*x + cos*y
			color := mandelbrot(complex(xNew, yNew))
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
