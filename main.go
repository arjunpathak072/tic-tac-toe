package main

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	framesPerSecond = 60
	windowSize      = 900
)

type color struct {
	r, g, b, a uint8
}

var winner winners

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println("initializing SDL: ", err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"Tic Tac Toe",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		windowSize,
		windowSize,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		fmt.Println("creating window: ", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("creating a renderer: ", err)
		return
	}
	defer renderer.Destroy()

	texture, err := renderer.CreateTexture(
		sdl.PIXELFORMAT_ABGR8888,
		sdl.TEXTUREACCESS_STATIC,
		windowSize,
		windowSize,
	)
	if err != nil {
		fmt.Println("initializing texture: ", err)
		return
	}
	defer texture.Destroy()

	var frameStart time.Time
	var elapsedTime float32
	var turn bool = false
	winner = nobody

	drawGameBoard(color{128, 128, 128, 255})

	for {
		var boxX, boxY, centreX, centreY int
		frameStart = time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.MouseButtonEvent:
				if event.GetType() == sdl.MOUSEBUTTONDOWN && event.Button == sdl.BUTTON_LEFT {
					turn = !turn
					boxX, boxY, centreX, centreY = getQuardrantCentre(event)
					if turn {
						drawO(centreX, centreY, color{255, 255, 255, 255})
						winner = updateAndValidate(boxX-1, boxY-1, -1)
					} else {
						drawX(centreX, centreY, color{255, 255, 255, 255})
						winner = updateAndValidate(boxX-1, boxY-1, 1)
					}
				}
			}
		}

		if winner != nobody {
			texture.Update(nil, unsafe.Pointer(&pixels[0]), windowSize*4)
			renderer.Copy(texture, nil, nil)
			renderer.Present()	
			sdl.Delay(2000)
			break
		}

		texture.Update(nil, unsafe.Pointer(&pixels[0]), windowSize*4)
		renderer.Copy(texture, nil, nil)
		renderer.Present()

		elapsedTime = float32(time.Since(frameStart).Seconds())
		if elapsedTime < 0.005 {
			sdl.Delay(5 - uint32(elapsedTime*1000.0))
		}
	}
}
