package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "" //Here put the input
	inputTempArray := strings.Split(input, "\n")
	lengthTempArray := len(inputTempArray)
	inputArray := make([][]int, lengthTempArray)
	for index, numbers := range inputTempArray {
		tempStringsArray := strings.Split(numbers, "\t")
		for _, stringNumber := range tempStringsArray {
			intNumber, _ := strconv.Atoi(stringNumber)
			inputArray[index] = append(inputArray[index], intNumber)
		}
	}
	result := 0
	for _, array := range inputArray {
		firstNumber, secondNumber := findEvenlyDivide(array)
		result = result + (firstNumber / secondNumber)
	}
	fmt.Println(result)
}

func findEvenlyDivide(array []int) (int, int) {
	for _, firstNumber := range array {
		for _, secondNumber := range array {
			if firstNumber%secondNumber == 0 && firstNumber != secondNumber {
				return firstNumber, secondNumber
			}
		}
	}
	return 0, 0
}
