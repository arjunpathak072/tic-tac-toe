package main

import "fmt"

var gameState = [3][3]int{}

type winners int
const (
	nobody = iota
	cross = iota
	circle = iota
)

func updateAndValidate(x, y, val int) (winner winners) {
	gameState[x][y] = val

	fmt.Println(gameState)

	if checkRow(x, val) ||
	checkColumn(y, val) ||
	checkDiagonal(val) {
		if val == 1 {
			fmt.Println("Cross won!")
			return cross
		} else {
			fmt.Println("Circle won!")
			return circle
		}
	}
	return nobody
}

func checkRow(x, val int) (flag bool) {
	for i := 0; i < 3; i++ {
		if gameState[x][i] != val {
			flag = true
		}
	}
	return !flag
}

func checkColumn(y, val int) (flag bool) {
	for i := 0; i < 3; i++ {
		if gameState[i][y] != val {
			flag = true
		}
	}
	return !flag
}

func checkDiagonal(val int) (flag bool) {
	for i := 0; i < 3; i++ {
		if gameState[i][i] != val {
			flag = true
		}
	}
	if !flag {
		return true
	}

	flag = false
	y := 2
	for x := 0; x < 3; x++ {
		if gameState[x][y] != val {
			flag = true
		}
		y--
	}
	return !flag
}