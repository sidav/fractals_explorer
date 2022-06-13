package main

import (
	"fmt"
	"fractals_explorer/middleware"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"
)

func exportToPng(useHighPrecision bool) {
	w, h := config.Export.Width, config.Export.Height
	initialPrecision := maxSetCheckIterations
	if useHighPrecision {
		maxSetCheckIterations = config.Export.PreciseIterations
	}
	fileName := fmt.Sprintf("o%d_%d.png", orderOfFractalExpression, time.Now().UnixNano())
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

	drawExportText(fmt.Sprintf("Exporting..."), 2)

	var iterations int
	for x := 0; x < w; x++ {
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

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, img)
	if useHighPrecision {
		maxSetCheckIterations = initialPrecision
	}

	rl.BeginTextureMode(middleware.TargetTexture)
	drawExportText(fmt.Sprintf("Export done."), 3)
}

func drawExportText(text string, lineNum int32) {
	textSize := int32(RenderHeight / 30)
	rl.BeginTextureMode(middleware.TargetTexture)
	rl.DrawText(text,
		0, (textSize+2) * lineNum, textSize, color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 255,
		})
	rl.EndTextureMode()
	middleware.Flush()
}
