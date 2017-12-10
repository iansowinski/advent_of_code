package main

import (
	"fmt"
	// "sort"
	// "strconv"
	// "strings"
)

func main() {
	// input := "" //Here put the input
	// lengths := []int{3, 4, 1, 5}
	// array := []int{0, 1, 2, 3, 4}
	lengths := []int{94,84,0,79,2,27,81,1,123,93,218,23,103,255,254,243}
	array := generateArray(256)
	array = generateHash(lengths, array)
	fmt.Println(array[0] * array[1])
}

func generateHash(lengths []int, array []int) []int {
	skip := 0
	marker := 0
	for _, length := range lengths {
		var tempArray []int
		if marker > len(array) {
			marker = marker - len(array)
		}
		if marker + length < len(array) {
			tempArray = reverse(array[marker:(marker + length)])
			array = replace(marker, array, tempArray)
			marker =  incresePosition(marker, length, skip, len(array))
			skip += 1
		} else {
			leftSide := array[(marker):]
			rightSide := array[:(length - (len(array) - marker ))]
			tempArray = append(leftSide, rightSide...)
			array = replace(marker, array, reverse(tempArray))
			marker = incresePosition(marker, length, skip, len(array))
			skip += 1
		} 
	}
	return array
}

func generateArray (size int) []int {
	array := make([]int, size)
	i := 0
	for i < size {
		array[i] = i
		i += 1
	}
	return array
}

func incresePosition (marker int, length int, skip int, arrayLen int) int {
	if marker + length + skip < arrayLen {
		return marker + length + skip
	}
	return  marker + length + skip - arrayLen
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