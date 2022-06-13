package main

import (
	"fmt"
	"fractals_explorer/middleware"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"math/rand"
	"time"
)

var (
	WINDOW_W                  = 1920 / 2
	WINDOW_H                  = 1080 / 2
	PIXEL_FACTOR              = 2
	RenderWidth, RenderHeight int
	surface                   complexSurface
	juliaParameter            *complex
	currentFractal            int // 0 is Mandelbrot, 1 is Julia
	exit                      bool
	config                    Config
)

func main() {
	rl.InitWindow(int32(WINDOW_W), int32(WINDOW_H), "Fractal Explorer by Sidav")
	rl.SetWindowState(rl.FlagWindowResizable)
	rl.SetTargetFPS(60)
	rl.SetExitKey(rl.KeyEscape)

	RenderWidth, RenderHeight = WINDOW_W/PIXEL_FACTOR, int(WINDOW_H)/PIXEL_FACTOR
	middleware.SetInternalResolution(int32(RenderWidth), int32(RenderHeight))
	middleware.SetColor(0, 0, 0)
	middleware.FillRect(0, 0, RenderWidth, RenderHeight)

	rand.Seed(time.Now().UnixNano())

	config.initFromFile()

	showHelpWindow()

	surface.init(RenderWidth, RenderHeight)
	juliaParameter = newComplex(-0.66, -0.34)

	drawCurrentFractal()
	for !exit {
		forceRedraw := false
		if rl.IsWindowResized() {
			forceRedraw = true
			handleResize()
		}
		if workKeys() || forceRedraw {
			drawCurrentFractal()
		}
	}

	rl.CloseWindow()
}

func drawCurrentFractal() {
	textSize := int32(RenderHeight / 30)
	rl.BeginTextureMode(middleware.TargetTexture)
	rl.ClearBackground(rl.Black)
	duration := drawFractal(juliaParameter)
	if currentFractal == 0 {
		rl.DrawText(fmt.Sprintf("Mandelbrot^%d, iters %d, drawn in %dms", orderOfFractalExpression, maxSetCheckIterations, duration),
			0, 0, textSize, color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
			})
	} else {
		rl.DrawText(fmt.Sprintf("Julia^%d, param is %s, iters %d, drawn in %dms",
			orderOfFractalExpression, juliaParameter.toString(), maxSetCheckIterations, duration),

			0, 0, textSize, color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
			})
	}
	rl.EndTextureMode()
	middleware.Flush()
}

func handleResize() {
	WINDOW_W, WINDOW_H = rl.GetScreenWidth(), rl.GetScreenHeight()
	RenderWidth = WINDOW_W / PIXEL_FACTOR
	RenderHeight = WINDOW_H / PIXEL_FACTOR
	middleware.SetInternalResolution(int32(RenderWidth), int32(RenderHeight))
	surface.reinit(RenderWidth, RenderHeight)
}

func showHelpWindow() {
	lines := []string{
		"FRACTALS EXPLORER",
		"by sidav, 2022",
		"",
		"Arrows move the camera.",
		"[ and ] change fractal order",
		", and . change calculation precision",
		"; and ' change resolution",
		"- and + change zoom level",
		"c increases base hue",
		"e exports image to png",
		"r exports with high precision",
		"BACKSPACE resets the camera",
		"TAB changes mode",
		"SPACE generates new parameter for Julia set",
		"",
		"Press ENTER to start exploring!",
	}

	rl.BeginTextureMode(middleware.TargetTexture)
	rl.ClearBackground(rl.Black)

	textSize := int32(RenderHeight / 20)
	for i := range lines {
		rl.DrawText(lines[i], 0, int32(i)*(textSize+1), textSize, color.RGBA{
			R: 160,
			G: 64,
			B: 128,
			A: 255,
		})
	}
	rl.EndTextureMode()
	for {
		middleware.Flush()
		if rl.GetKeyPressed() == rl.KeyEnter {
			time.Sleep(250 * time.Millisecond)
			break
		}
	}
}
