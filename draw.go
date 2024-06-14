package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

func drawGameBoard(color color) {
	lineWidth := 5
	startX := windowSize/3 - lineWidth
	startY := windowSize/3 - lineWidth

	for y := 0; y < windowSize; y++ {
		for x := startX; x < startX+lineWidth; x++ {
			setPixel(x, y, color)
		}
	}

	startX += windowSize / 3
	for y := 0; y < windowSize; y++ {
		for x := startX; x < startX+lineWidth; x++ {
			setPixel(x, y, color)
		}
	}

	for y := startY; y < startY+lineWidth; y++ {
		for x := 0; x < windowSize; x++ {
			setPixel(x, y, color)
		}
	}

	startY += windowSize / 3
	for y := startY; y < startY+lineWidth; y++ {
		for x := 0; x < windowSize; x++ {
			setPixel(x, y, color)
		}
	}
}

func getQuardrantCentre(event *sdl.MouseButtonEvent) (boxX, boxY, centreX, centreY int) {
	boxX = int((event.X / 300) + 1)
	boxY = int((event.Y / 300) + 1)

	xOffset := 150 * (boxX - 1)
	yOffset := 150 * (boxY - 1)

	return boxX, boxY, (boxX * 150 + xOffset), (boxY * 150 + yOffset)
}

func drawX(centreX, centreY int, color color) {
	padding := 20
	startX := centreX - windowSize/6
	startY := centreY - windowSize/6 + padding
	endX := centreX + windowSize/6
	endY := centreY + windowSize/6 - padding

	x1 := startX + padding
	x2 := endX - padding
	for y := startY; y < endY; y++ {
		setPixel(x1, y, color)
		setPixel(x2, y, color)
		x1++
		x2--
	}
}

func drawO(centreX, centreY int, color color) {
	padding := 20
	delta := 1.0
	radius := windowSize/6 - padding
	startX := centreX - windowSize/6
	startY := centreY - windowSize/6
	endX := centreX + windowSize/6
	endY := centreY + windowSize/6

	for y:= startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			distance := math.Sqrt(math.Pow(float64(centreX) - float64(x), 2) + math.Pow(float64(centreY) - float64(y), 2))
			if distance <= float64(radius) + delta && distance >= float64(radius) - delta {
				setPixel(x, y, color)
			}
		}
	}
}