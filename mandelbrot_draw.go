package main

import (
	"fractals_explorer/middleware"
	"time"
)

func drawMandelbrot() int {
	startTime := time.Now()
	for x := 0; x < RenderWidth; x++ {
		for y := 0; y < RenderHeight; y++ {
			complexPixel := surface.pixelToComplex(float64(x), float64(y))
			iterations := getMandelbrotIterations(complexPixel)
			if iterations == -1 {
				middleware.SetColor(0, 0, 0)
			} else {
				r, g, b := getSpectrumColorFor(iterations, 0, maxSetCheckIterations)
				middleware.SetColor(r, g, b)
			}
			middleware.DrawPoint(int32(x), int32(y))
		}
	}
	return int(time.Since(startTime).Milliseconds())
}
