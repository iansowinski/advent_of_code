package main

import (
	"fmt"
	// "sort"
	// "strconv"
	// "strings"
)

func main() {
	// input := "" //Here put the input
	lengths := []int{3, 4, 1, 5}
	array := []int{0, 1, 2, 3, 4}
	skip := 0
	marker := 0
	for _, length := range lengths {
		var tempArray []int
		if marker + length + skip < len(array) {
			tempArray = reverse(array[marker:(marker + length + skip)])
			array = replace(marker, array, tempArray)
			marker = length + skip
			skip += 1
		} else {
			leftSide := array[(marker):]
			rightSide := array[:(length - (len(array) - marker))]
			tempArray = append(leftSide, rightSide...)
			// fmt.Println("tempArray")
			// fmt.Println(reverse(tempArray))
			// fmt.Println("/tempArray")
			array = replace((marker), array, reverse(tempArray))
			marker = length + skip
			skip += 1
		} 
		fmt.Println(array)
	}
}

func reverse (array []int) []int {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func replace (start int, array []int, reversedFragment []int) []int {
	for index, item := range reversedFragment {
		if index + start < len(array) {
			array[index + start] = item
		} else {
			array[(index + start) - len(array)] = item
		}
	}
	return array
}