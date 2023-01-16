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
	fileName := fmt.Sprintf("o%d_%s.png", orderOfFractalExpression, time.Now().Format("2006_01_02_15_04_05"))

	switch currentFractal {
	case 0:
		fileName = "mandelbrot_" + fileName
	case 1:
		fileName = "julia_" + fileName
	default:
		panic("unknown fractal type")
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

	drawExportText(fmt.Sprintf("Exporting... Please don't press any keys"), 0)
	drawExportText(fmt.Sprintf("Calculating %dx%d picture with %d precision", w, h, maxSetCheckIterations), 1)

	var iterations int
	totalPixels := w * h
	prevPercent := 0
	currPixel := 0
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			currPixel++
			currentPercent := 100 * currPixel / totalPixels
			if currentPercent != prevPercent {
				drawExportText(fmt.Sprintf("Calculating: %d%%\n", currentPercent), 2)
				prevPercent = currentPercent
			}

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
	drawExportText(fmt.Sprintf("Saving output file."), 3)
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
	drawExportText(fmt.Sprintf("Saved to %s", fileName), 0)
	drawExportText(fmt.Sprintf("Export completed."), 4)
}

func drawExportText(text string, lineNum int32) {
	textSize := int32(RenderHeight / 20)
	rl.BeginTextureMode(middleware.TargetTexture)
	rl.DrawRectangle(0, lineNum*(textSize+2), rl.MeasureText(text, textSize+1), textSize+2, rl.Black)
	rl.DrawText(text,
		0, (textSize+2)*lineNum, textSize, color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		})
	rl.EndTextureMode()
	middleware.Flush()
}
