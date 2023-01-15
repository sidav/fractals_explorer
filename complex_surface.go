package main

import "math/big"

const zoomFactor = 10

type complexSurface struct {
	topLeftPixelValue     *complex
	BottomRightPixelValue *complex
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
	cs.topLeftPixelValue = newComplex(
		-rSize/2,
		-iSize/2,
	)
	cs.BottomRightPixelValue = newComplex(
		rSize/2,
		iSize/2,
	)
	cs.refresh()
}

func (cs *complexSurface) getCenter() *complex {
	TopLeftRealFloat, _ := cs.topLeftPixelValue.real.Float64()
	TopLeftImFloat, _ := cs.topLeftPixelValue.imaginary.Float64()
	return newComplex(
		TopLeftRealFloat+cs.horizSize/2,
		TopLeftImFloat+cs.vertSize/2,
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
	cs.topLeftPixelValue = newComplex(
		-rSize/2,
		-iSize/2,
	)
	cs.BottomRightPixelValue = newComplex(
		rSize/2,
		iSize/2,
	)
	cs.refresh()
	cs.setCenterAt(currCenter)
}

func (cs *complexSurface) refresh() {
	result := big.NewFloat(0)
	cs.vertSize, _ = result.Sub(cs.BottomRightPixelValue.imaginary, cs.topLeftPixelValue.imaginary).Float64()
	cs.horizSize, _ = result.Sub(cs.BottomRightPixelValue.real, cs.topLeftPixelValue.real).Float64()
}

func (cs *complexSurface) setCenterAt(c *complex) {
	cs.topLeftPixelValue = subFloat(c, cs.horizSize/2, cs.vertSize/2)
	cs.BottomRightPixelValue = addFloat(c, cs.horizSize/2, cs.vertSize/2)
}

func (cs *complexSurface) moveBy(x, y float64) {
	// x-axis
	cs.topLeftPixelValue = addFloat(cs.topLeftPixelValue, x*cs.horizSize/zoomFactor, y*cs.vertSize/zoomFactor)
	cs.BottomRightPixelValue = addFloat(cs.BottomRightPixelValue, x*cs.horizSize/zoomFactor, y*cs.vertSize/zoomFactor)
}

func (cs *complexSurface) zoomIn() {
	cs.topLeftPixelValue = addFloat(cs.topLeftPixelValue, cs.horizSize/zoomFactor, cs.vertSize/zoomFactor)
	cs.BottomRightPixelValue = subFloat(cs.BottomRightPixelValue, cs.horizSize/zoomFactor, cs.vertSize/zoomFactor)
	cs.refresh()
}

func (cs *complexSurface) zoomOut() {
	cs.topLeftPixelValue = subFloat(cs.topLeftPixelValue, cs.horizSize/zoomFactor, cs.vertSize/zoomFactor)
	cs.BottomRightPixelValue = addFloat(cs.BottomRightPixelValue, cs.horizSize/zoomFactor, cs.vertSize/zoomFactor)
	cs.refresh()
}

func (cs *complexSurface) pixelToComplex(x, y float64) *complex {
	// pixel 0, 0 is (-2 - 2i)
	// bottom-right pixel is 2+2i
	verticalFactor := cs.vertSize / float64(cs.screenH)
	horizontalFactor := cs.horizSize / float64(cs.screenW)
	return addFloat(cs.topLeftPixelValue, x*horizontalFactor, y*verticalFactor)
	// return newComplex(cs.topLeftPixelValue.real+x*horizontalFactor, cs.topLeftPixelValue.imaginary+y*verticalFactor)
}
