package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"
)

func exportToPng() {
	w, h := config.Export.Width, config.Export.Height
	fileName := fmt.Sprintf("o%d_%d_%d.png", orderOfFractalExpression, time.Now().Hour(),
		time.Now().Minute())
	if currentFractal == 0 {
		fileName = "mandelbrot_" + fileName
	}
	if currentFractal == 1 {
		fileName = "julia_" + fileName
	}

	imageCmpSurf := &complexSurface{
		topLeftPixelValue:     surface.topLeftPixelValue,
		BottomRightPixelValue: surface.BottomRightPixelValue,
		horizSize:             surface.horizSize,
		vertSize:              surface.vertSize,
	}
	imageCmpSurf.reinit(w, h)

	img := image.NewRGBA(image.Rect(0, 0, w, h))
	currColor := color.RGBA{}

	fmt.Println("Calculating for export...")
	var iterations int
	for x := 0; x < w; x++ {
		fmt.Printf(" x: %d\n", x)
		for y := 0; y < h; y++ {
			complexPixel := imageCmpSurf.pixelToComplex(float64(x), float64(y))

			if currentFractal == 0 {
				iterations = getMandelbrotIterations(complexPixel)
			}
			if currentFractal == 1 {
				iterations = getJuliaIterations(complexPixel, juliaParameter)
			}

			if iterations == -1 {
				currColor = color.RGBA{
					R: 0,
					G: 0,
					B: 0,
					A: 255,
				}
			} else {
				r, g, b := getSpectrumColorFor(iterations, 0, maxSetCheckIterations)
				currColor = color.RGBA{
					R: r,
					G: g,
					B: b,
					A: 255,
				}
			}
			img.Set(x, y, currColor)
		}
	}

	fmt.Println("Exporting to file...")
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, img)
	fmt.Println("Export done.")
}
