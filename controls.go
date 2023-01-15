package main

import (
	"fractals_explorer/middleware"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"time"
)

func workKeys() bool { // true if redraw/recalculation needed
	middleware.Flush()
	if rl.IsKeyDown(rl.KeyEscape) {
		exit = true
		return false
	}
	// move "camera"
	if rl.IsKeyDown(rl.KeyLeft) {
		surface.moveBy(-1, 0)
		return true
	}
	if rl.IsKeyDown(rl.KeyRight) {
		surface.moveBy(1, 0)
		return true
	}
	if rl.IsKeyDown(rl.KeyUp) {
		surface.moveBy(0, -1)
		return true
	}
	if rl.IsKeyDown(rl.KeyDown) {
		surface.moveBy(0, 1)
		return true
	}
	// zoom "camera"
	if rl.IsKeyDown(rl.KeyMinus) {
		surface.zoomOut()
		return true
	}
	if rl.IsKeyDown(rl.KeyEqual) {
		surface.zoomIn()
		return true
	}
	// reset "camera"
	if rl.IsKeyDown(rl.KeyBackspace) {
		surface.init(RenderWidth, RenderHeight)
		return true
	}
	// change fractal
	if rl.IsKeyDown(rl.KeyTab) {
		surface.init(RenderWidth, RenderHeight)
		currentFractal = (currentFractal + 1) % 2
		time.Sleep(200 * time.Millisecond)
		return true
	}
	if rl.IsKeyDown(rl.KeyLeftBracket) {
		orderOfFractalExpression--
		if orderOfFractalExpression < 2 {
			orderOfFractalExpression = 2
		}
		time.Sleep(200 * time.Millisecond)
		return true
	}
	if rl.IsKeyDown(rl.KeyRightBracket) {
		orderOfFractalExpression++
		time.Sleep(200 * time.Millisecond)
		return true
	}
	// change iterations
	if rl.IsKeyDown(rl.KeyComma) {
		maxSetCheckIterations--
		if maxSetCheckIterations < 1 {
			maxSetCheckIterations = 1
			return false
		}
		return true
	}
	if rl.IsKeyDown(rl.KeyPeriod) {
		maxSetCheckIterations++
		return true
	}
	// handle resolution
	if rl.IsKeyDown(rl.KeySemicolon) {
		PIXEL_FACTOR -= 1
		if PIXEL_FACTOR < 1 {
			PIXEL_FACTOR = 1
		}
		handleResize()
		time.Sleep(200 * time.Millisecond)
		return true
	}
	if rl.IsKeyDown(rl.KeyApostrophe) {
		PIXEL_FACTOR += 1
		//if PIXEL_FACTOR < 1 {
		//	PIXEL_FACTOR = 1
		//}
		handleResize()
		time.Sleep(200 * time.Millisecond)
		return true
	}
	// change hue
	if rl.IsKeyDown(rl.KeyC) {
		baseHue += 0.025
		return true
	}
	// export image
	if rl.IsKeyDown(rl.KeyE) {
		exportToPng(false)
		return false
	}
	// export image with high precision
	if rl.IsKeyDown(rl.KeyR) {
		exportToPng(true)
		return false
	}
	// get new julia parameter
	if rl.IsKeyDown(rl.KeySpace) {
		parameter := randomComplex(-2, 2, -2, 2)
		juliaParameter.setEqualTo(parameter)
		for !isPartOfMandelbrotForPrecision(parameter, rand.Intn(50)+1) {
			parameter = randomComplex(-2, 2, -2, 2)
			juliaParameter.setEqualTo(parameter)
		}
		surface.init(RenderWidth, RenderHeight)
		time.Sleep(100 * time.Millisecond)
		return true
	}
	// mouse wheel
	mwm := rl.GetMouseWheelMove()
	if mwm > 0 {
		mx, my := float64(rl.GetMouseX()), float64(rl.GetMouseY())
		comp := surface.pixelToComplex(mx/float64(PIXEL_FACTOR), my/float64(PIXEL_FACTOR))
		surface.setCenterAt(comp)
		surface.zoomIn()
		rl.SetMousePosition(WINDOW_W/2, int(WINDOW_H)/2)
		return true
	}
	if mwm < 0 {
		mx, my := float64(rl.GetMouseX()), float64(rl.GetMouseY())
		comp := surface.pixelToComplex(mx/float64(PIXEL_FACTOR), my/float64(PIXEL_FACTOR))
		surface.setCenterAt(comp)
		surface.zoomOut()
		rl.SetMousePosition(WINDOW_W/2, int(WINDOW_H)/2)
		return true
	}
	return false
}
