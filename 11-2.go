package main

import (
	"fmt"
	"strings"
)

type hexField struct {
	x int
	y int
	z int
}

func main() {
	input := "" //Here put the input
	directionsArray := strings.Split(input, ",")
	_, max := createBoardWithMax(directionsArray)
	fmt.Println(max)
}

func createBoardWithMax (directionsArray []string) ([]hexField, int) {
	var currentPosition hexField
	currentPosition.x = 0
	currentPosition.y = 0
	currentPosition.z = 0
	max := 0
	hexBoard := make([]hexField, 0)
	hexBoard = append(hexBoard, currentPosition)
	for _, item := range directionsArray {
		switch item {
		case "n":
			currentPosition.y += 1
			currentPosition.z -= 1
		case "s":
			currentPosition.y -= 1
			currentPosition.z += 1
		case "ne":
			currentPosition.x += 1
			currentPosition.z -= 1
		case "sw":
			currentPosition.x -= 1
			currentPosition.z += 1
		case "nw":
			currentPosition.y += 1
			currentPosition.x -= 1
		case "se":
			currentPosition.y -= 1
			currentPosition.x += 1
		}
		hexBoard = append(hexBoard, currentPosition)
		if cubeDistance(hexBoard[0], currentPosition) > max {
			max = cubeDistance(hexBoard[0], currentPosition)
		}
	}
	return hexBoard, max
}

func cubeDistance(firstHex hexField, secondHex hexField) int {
	return (abs(firstHex.x - secondHex.x) + abs(firstHex.y - secondHex.y) + abs(firstHex.z - secondHex.z)) / 2
}

func abs(number int) int {
	if number < 0 {
	   number = -number
	}
	return number
}

