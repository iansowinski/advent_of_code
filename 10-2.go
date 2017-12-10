package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "" //Here put the input
	array := generateArray(256)
	byteArrayTemp := []rune(input)
	toAdd := []rune{17, 31, 73, 47, 23}
	byteArrayTemp = append(byteArrayTemp, toAdd...)
	lengths := make([]int, len(byteArrayTemp))
	for index, item := range byteArrayTemp {
		lengths[index] = int(item)
	}
	marker := 0
	skip := 0
	i := 0
	for i < 64 {
		array, marker, skip = generateHash(lengths, array, marker, skip)
		i += 1
	}
	sparseHash := createSparseHash(array)
	denseHash := make([]string, 16)

	for index, item := range sparseHash {
		current := 0
		for _, childItem := range item {
			current = current ^ childItem
		}
		denseHash[index] = strconv.FormatInt(int64(current), 16)
	}
	//hack for disapearing 0 :(
	for index, item := range denseHash {
		if len(item) == 1 {
			item = fmt.Sprintf("0%s", item)
			denseHash[index] = item
		}
	}
	fmt.Println(strings.Join(denseHash, ""))
}

func createSparseHash(array []int) [][]int {
	sparseHash := make([][]int, 16)
	i := 0
	for _, item := range array {
		sparseHash[i/16] = append(sparseHash[i/16], item)
		i += 1
	}
	return sparseHash
}

func recursiveMarkerCheck(marker int, length int) int {
	if marker > length {
		marker = recursiveMarkerCheck(marker-length, length)
	}
	return marker
}

func generateHash(lengths []int, array []int, marker int, skip int) ([]int, int, int) {
	for _, length := range lengths {
		var tempArray []int
		marker = recursiveMarkerCheck(marker, len(array))
		if marker+length < len(array) {
			tempArray = reverse(array[marker:(marker + length)])
			array = replace(marker, array, tempArray)
			marker = incresePosition(marker, length, skip, len(array))
			skip += 1
		} else {
			leftSide := array[(marker):]
			rightSide := array[:(length - (len(array) - marker))]
			tempArray = append(leftSide, rightSide...)
			array = replace(marker, array, reverse(tempArray))
			marker = incresePosition(marker, length, skip, len(array))
			skip += 1
		}
	}
	return array, marker, skip
}

func generateArray(size int) []int {
	array := make([]int, size)
	i := 0
	for i < size {
		array[i] = i
		i += 1
	}
	return array
}

func incresePosition(marker int, length int, skip int, arrayLen int) int {
	if marker+length+skip < arrayLen {
		return marker + length + skip
	}
	return marker + length + skip - arrayLen
}

func reverse(array []int) []int {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func replace(start int, array []int, reversedFragment []int) []int {
	for index, item := range reversedFragment {
		if index+start < len(array) {
			array[index+start] = item
		} else {
			array[(index+start)-len(array)] = item
		}
	}
	return array
}
