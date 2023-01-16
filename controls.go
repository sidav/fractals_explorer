package main

import (
	"fractals_explorer/middleware"
	rl "github.com/gen2brain/raylib-go/raylib"
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
		currentFractal = (currentFractal + 1) % TOTAL_FRACTALS
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
		if rl.IsKeyDown(rl.KeyLeftShift) {
			maxSetCheckIterations -= 5
		} else {
			maxSetCheckIterations--
		}
		if maxSetCheckIterations < 1 {
			maxSetCheckIterations = 1
			return false
		}
		return true
	}
	if rl.IsKeyDown(rl.KeyPeriod) {
		if rl.IsKeyDown(rl.KeyLeftShift) {
			maxSetCheckIterations += 5
		} else {
			maxSetCheckIterations++
		}
		return true
	}
	// handle resolution
	if rl.IsKeyDown(rl.KeySemicolon) {
		if rl.IsKeyDown(rl.KeyLeftShift) {
			PIXEL_FACTOR -= 2
		} else {
			PIXEL_FACTOR -= 1
		}
		if PIXEL_FACTOR < 1 {
			PIXEL_FACTOR = 1
		}
		handleResize()
		time.Sleep(200 * time.Millisecond)
		return true
	}
	if rl.IsKeyDown(rl.KeyApostrophe) {
		if rl.IsKeyDown(rl.KeyLeftShift) {
			PIXEL_FACTOR += 2
		} else {
			PIXEL_FACTOR += 1
		}
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
		exportToPng(rl.IsKeyDown(rl.KeyLeftShift))
		return false
	}
	// get new julia parameter
	if rl.IsKeyDown(rl.KeySpace) {
		generateNewJuliaParameter()
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
		zoomTimes := 1
		if rl.IsKeyDown(rl.KeyLeftShift) {
			zoomTimes = 5
		}
		for i := 0; i < zoomTimes; i++ {
			surface.zoomIn()
		}
		rl.SetMousePosition(WINDOW_W/2, int(WINDOW_H)/2)
		return true
	}
	if mwm < 0 {
		mx, my := float64(rl.GetMouseX()), float64(rl.GetMouseY())
		comp := surface.pixelToComplex(mx/float64(PIXEL_FACTOR), my/float64(PIXEL_FACTOR))
		surface.setCenterAt(comp)
		zoomTimes := 1
		if rl.IsKeyDown(rl.KeyLeftShift) {
			zoomTimes = 5
		}
		for i := 0; i < zoomTimes; i++ {
			surface.zoomOut()
		}
		rl.SetMousePosition(WINDOW_W/2, int(WINDOW_H)/2)
		return true
	}
	return false
}
