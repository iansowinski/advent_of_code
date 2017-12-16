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
	input := `` //Here put the input
	scanner, directions := createScanner(input)
	max, _ := maxMin(scanner)
	index := 0
	caught := make([]int, 0)
	for index <= max {
		_, contains := scanner[index]
		if contains {
			if scanner[index][0] == 1 {
				caught = append(caught, index)
			}
		}
		scanner, directions = moveScanner(scanner, directions)
		index += 1
	}
	severity := 0
	for _, item := range caught {
		severity += item * len(scanner[item])
	}
	fmt.Println(severity)
}

func createScanner(input string) (map[int][]int, map[int]string) {
	inputListTemp := strings.Split(input, "\n")
	inputMap := make(map[int][]int, 0)
	directions := make(map[int]string, 0)
	for _, item := range inputListTemp {
		firstSplit := strings.Split(item, ": ")
		index, _ := strconv.Atoi(firstSplit[0])
		layers, _ := strconv.Atoi(firstSplit[1])
		tempArray := make([]int, layers)
		inputMap[index] = tempArray
	}
	for index, _ := range inputMap {
		inputMap[index][0] = 1
		directions[index] = "up"
	}
	return inputMap, directions
}

func moveScanner(scanner map[int][]int, directions map[int]string) (map[int][]int, map[int]string) {
	for index, item := range scanner {
		maxIndex := len(item) - 1
		for indexChild, itemChild := range item {
			if itemChild == 1 && directions[index] == "up" {
				if indexChild == maxIndex {
					scanner[index][indexChild] = 0
					scanner[index][indexChild-1] = 1
					directions[index] = "down"
					break
				} else {
					scanner[index][indexChild] = 0
					scanner[index][indexChild+1] = 1
					break
				}
			} else if itemChild == 1 && directions[index] == "down" {
				if indexChild == 0 {
					scanner[index][indexChild] = 0
					scanner[index][indexChild+1] = 1
					directions[index] = "up"
					break
				} else {
					scanner[index][indexChild] = 0
					scanner[index][indexChild-1] = 1
					break
				}
			}
		}
	}
	return scanner, directions
}

func maxMin(scanner map[int][]int) (int, int) {
	firstTime := true
	var max int
	var min int
	for index, _ := range scanner {
		if firstTime {
			max = index
			min = index
			firstTime = false
		}
		if index > max {
			max = index
		}
		if index < min {
			min = index
		}
	}
	return max, min
}
