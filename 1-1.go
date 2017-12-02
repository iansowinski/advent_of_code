package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "" //Here put the input
	inputArray := strings.Split(input, "")
	lastElementIndex := len(inputArray) - 1
	result := 0
	for index, number := range inputArray {
		if index != lastElementIndex {
			if number == inputArray[index+1] {
				numberInt, _ := strconv.Atoi(number)
				result += numberInt
			}
		} else {
			if number == inputArray[0] {
				numberInt, _ := strconv.Atoi(number)
				result += numberInt
			}
		}
	}
	fmt.Println(result)
}
