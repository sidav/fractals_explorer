package main

type complexSurface struct {
	topLeftPixelValue     complex
	BottomRightPixelValue complex
	horizSize, vertSize   float64
}

func (cs *complexSurface) init(screenW, screenH int) {
	const initDist = 4.0
	var rSize, iSize float64
	if screenW > screenH {
		rSize = initDist
		iSize = float64(screenH) / float64(screenW) * initDist
	} else {
		rSize = float64(screenW) / float64(screenH) * initDist
		iSize = initDist
	}
	cs.topLeftPixelValue = complex{
		real:      -rSize/2,
		imaginary: -iSize/2,
	}
	cs.BottomRightPixelValue = complex{
		real:      rSize/2,
		imaginary: iSize/2,
	}
	cs.refresh()
}

func (cs *complexSurface) refresh() {
	cs.vertSize = cs.BottomRightPixelValue.imaginary - cs.topLeftPixelValue.imaginary
	cs.horizSize = cs.BottomRightPixelValue.real - cs.topLeftPixelValue.real
}

func (cs *complexSurface) setCenterAt(c *complex) {
	cs.topLeftPixelValue = complex{
		real:      c.real - cs.horizSize/2,
		imaginary: c.imaginary - cs.vertSize/2,
	}
	cs.BottomRightPixelValue = complex{
		real:      c.real + cs.horizSize/2,
		imaginary: c.imaginary + cs.vertSize/2,
	}
}

func (cs *complexSurface) moveBy(x, y float64) {
	// x-axis
	cs.topLeftPixelValue.real += x * (cs.horizSize / 10)
	cs.BottomRightPixelValue.real += x * (cs.horizSize / 10)
	// y-axis
	cs.topLeftPixelValue.imaginary += y * (cs.vertSize / 10)
	cs.BottomRightPixelValue.imaginary += y * (cs.vertSize / 10)
}

func (cs *complexSurface) zoomIn(factor float64) {
	// x-axis
	cs.topLeftPixelValue.real += cs.horizSize / factor
	cs.BottomRightPixelValue.real -= cs.horizSize / factor
	// y-axis
	cs.topLeftPixelValue.imaginary += cs.vertSize / factor
	cs.BottomRightPixelValue.imaginary -= cs.vertSize / factor

	cs.refresh()
}

func (cs *complexSurface) zoomOut(factor float64) {
	// x-axis
	cs.topLeftPixelValue.real -= cs.horizSize / factor
	cs.BottomRightPixelValue.real += cs.horizSize / factor
	// y-axis
	cs.topLeftPixelValue.imaginary -= cs.vertSize / factor
	cs.BottomRightPixelValue.imaginary += cs.vertSize / factor

	cs.refresh()
}


func (cs *complexSurface) pixelToComplex(x, y float64) *complex {
	// pixel 0, 0 is (-2 - 2i)
	// bottom-right pixel is 2+2i
	verticalFactor := cs.vertSize / float64(RenderHeight)
	horizontalFactor := cs.horizSize / float64(RenderWidth)
	return newComplex(cs.topLeftPixelValue.real+x*horizontalFactor, cs.topLeftPixelValue.imaginary+y*verticalFactor)
}
