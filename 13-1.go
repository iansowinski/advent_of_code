package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	id           int
	children     []node
	childrenTemp []int
}

func main() {
	input := `0: 3
1: 2
4: 4
6: 4` //Here put the input
	scanner := createScanner(input)
	fmt.Println(scanner[6])
	scanner = moveScanner(scanner)
	fmt.Println(scanner[6])
	scanner = moveScanner(scanner)
	fmt.Println(scanner[6])
	scanner = moveScanner(scanner)
	fmt.Println(scanner[6])
}

func createScanner (input string) map[int][]int {
	inputListTemp := strings.Split(input, "\n")
	inputMap := make(map[int][]int, 0)
	for _, item := range inputListTemp {
		firstSplit := strings.Split(item, ": ")
		index, _ := strconv.Atoi(firstSplit[0])
		layers, _ := strconv.Atoi(firstSplit[1])
		tempArray := make([]int, layers)
		inputMap[index] = tempArray
	}
	for index, _ := range inputMap {
		inputMap[index][0] = 1
	}
	return inputMap
}

func moveScanner(scanner map[int][]int) map[int][]int {
	for index, item := range scanner {
		maxIndex := len(item) - 1
		for indexChild, itemChild := range item {
			if itemChild == 1 {
				if indexChild == maxIndex {
					scanner[index][indexChild] = 0
					scanner[index][0] = 1
					break
				} else {
					scanner[index][indexChild] = 0
					scanner[index][indexChild + 1] = 1
					break
				}
			}
		}
	}
	return scanner
}