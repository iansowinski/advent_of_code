package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	input := "" //Here put the input
	stringArray := strings.Split(input, "\n")
	instructionsArrayLength := len(stringArray)
	instructionsArray := make([]int, instructionsArrayLength)
	pointer := 0
	result := 0
	for index, item := range stringArray {
		instructionsArray[index], _ = strconv.Atoi(item)
	}
	for true {
		if pointer >= 0 && pointer <= instructionsArrayLength - 1 {
			if instructionsArray[pointer] > 2 {
				result += 1
				instructionsArray[pointer] -= 1
				pointer += instructionsArray[pointer] + 1
			} else {
				result += 1
				instructionsArray[pointer] += 1
				pointer += instructionsArray[pointer] - 1
			}
		} else {
			break
		}
	}
	fmt.Println(result)
}