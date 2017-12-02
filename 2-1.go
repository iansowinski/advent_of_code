package main

import (
  "fmt"
  "strings"
  "strconv"
)

func main () {
	input := "" //Here put the input
  inputTempArray := strings.Split(input, "\n")
  lengthTempArray := len(inputTempArray)
  inputArray := make([][]int, lengthTempArray)
  for index, numbers := range inputTempArray {
  	tempStringsArray := strings.Split(numbers, "\t")
  	for _, stringNumber := range tempStringsArray{
  		intNumber, _ := strconv.Atoi(stringNumber)
  		inputArray[index] = append(inputArray[index], intNumber)
  	}
  }
  result := 0
  for _, array := range inputArray {
  	min, max := minMax(array)
  	result = result + (max - min)
  }
  fmt.Println(result)
}


func minMax(array []int)(int, int){
	var max int = array[0]
	var min int = array[0]
	for _, number := range array {
		if number > max {
			max = number
		}
		if number < min {
			min = number
		}
	}
	return min, max
}