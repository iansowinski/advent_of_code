package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	input := "265149"
	target, _ := strconv.Atoi(input)

	steps := 0
	turns := 0
	x := 0.0
	y := 0.0

	for steps < target-1 {
		length := (turns / 2) + 1
		for i := 0; i < length; i++ {
			if steps == target-1 {
				break
			}
			steps++
			direction := turns % 4
			switch direction {
			case 0:
				x++
			case 1:
				y++
			case 2:
				x--
			default:
				y--
			}
		}
		turns++
	}
	fmt.Println(math.Abs(x) + math.Abs(y))
}
