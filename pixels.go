package main

var pixels = make([]byte, windowSize * windowSize * 4)

func setPixel(x, y int, color color) {
	index := (y * windowSize + x) * 4

	if index < len(pixels) - 4 && index >= 0 {
		pixels[index] = color.r
		pixels[index + 1] = color.g
		pixels[index + 2] = color.b
		pixels[index + 3] = color.a
	}
}