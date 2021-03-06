package main

const zoomFactor = 10

type complexSurface struct {
	topLeftPixelValue     complex
	BottomRightPixelValue complex
	horizSize, vertSize   float64

	screenW, screenH int // in pixels
}

func (cs *complexSurface) init(screenW, screenH int) {
	cs.screenW, cs.screenH = screenW, screenH
	const baseSize = 4.0
	var rSize, iSize float64
	if screenW >= screenH {
		rSize = baseSize
		iSize = float64(screenH) / float64(screenW) * baseSize
	} else {
		rSize = float64(screenW) / float64(screenH) * baseSize
		iSize = baseSize
	}
	cs.topLeftPixelValue = complex{
		real:      -rSize / 2,
		imaginary: -iSize / 2,
	}
	cs.BottomRightPixelValue = complex{
		real:      rSize / 2,
		imaginary: iSize / 2,
	}
	cs.refresh()
}

func (cs *complexSurface) getCenter() *complex {
	return newComplex(
		cs.topLeftPixelValue.real+cs.horizSize/2,
		cs.topLeftPixelValue.imaginary+cs.vertSize/2,
	)
}

func (cs *complexSurface) reinit(screenW, screenH int) {
	cs.screenW, cs.screenH = screenW, screenH
	// save center for later recovery
	currCenter := cs.getCenter()

	var rSize, iSize = cs.horizSize, cs.vertSize
	if screenW >= screenH {
		iSize = float64(screenH) / float64(screenW) * cs.horizSize
	} else {
		rSize = float64(screenW) / float64(screenH) * cs.vertSize
	}
	cs.topLeftPixelValue = complex{
		real:      -rSize / 2,
		imaginary: -iSize / 2,
	}
	cs.BottomRightPixelValue = complex{
		real:      rSize / 2,
		imaginary: iSize / 2,
	}
	cs.refresh()
	cs.setCenterAt(currCenter)
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
	cs.topLeftPixelValue.real += x * (cs.horizSize / zoomFactor)
	cs.BottomRightPixelValue.real += x * (cs.horizSize / zoomFactor)
	// y-axis
	cs.topLeftPixelValue.imaginary += y * (cs.vertSize / zoomFactor)
	cs.BottomRightPixelValue.imaginary += y * (cs.vertSize / zoomFactor)
}

func (cs *complexSurface) zoomIn() {
	// x-axis
	cs.topLeftPixelValue.real += cs.horizSize / zoomFactor
	cs.BottomRightPixelValue.real -= cs.horizSize / zoomFactor
	// y-axis
	cs.topLeftPixelValue.imaginary += cs.vertSize / zoomFactor
	cs.BottomRightPixelValue.imaginary -= cs.vertSize / zoomFactor

	cs.refresh()
}

func (cs *complexSurface) zoomOut() {
	// x-axis
	cs.topLeftPixelValue.real -= cs.horizSize / zoomFactor
	cs.BottomRightPixelValue.real += cs.horizSize / zoomFactor
	// y-axis
	cs.topLeftPixelValue.imaginary -= cs.vertSize / zoomFactor
	cs.BottomRightPixelValue.imaginary += cs.vertSize / zoomFactor

	cs.refresh()
}

func (cs *complexSurface) pixelToComplex(x, y float64) *complex {
	// pixel 0, 0 is (-2 - 2i)
	// bottom-right pixel is 2+2i
	verticalFactor := cs.vertSize / float64(cs.screenH)
	horizontalFactor := cs.horizSize / float64(cs.screenW)
	return newComplex(cs.topLeftPixelValue.real+x*horizontalFactor, cs.topLeftPixelValue.imaginary+y*verticalFactor)
}
