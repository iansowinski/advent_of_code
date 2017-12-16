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
2: 4
4: 4
6: 5
8: 6
10: 8
12: 8
14: 6
16: 6
18: 9
20: 8
22: 6
24: 10
26: 12
28: 8
30: 8
32: 14
34: 12
36: 8
38: 12
40: 12
42: 12
44: 12
46: 12
48: 14
50: 12
52: 12
54: 10
56: 14
58: 12
60: 14
62: 14
64: 14
66: 14
68: 14
70: 14
72: 14
74: 20
78: 14
80: 14
90: 17
96: 18` //Here put the input
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

func createScanner (input string) (map[int][]int, map[int]string) {
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
					scanner[index][indexChild - 1] = 1
					directions[index] = "down"
					break
				} else {
					scanner[index][indexChild] = 0
					scanner[index][indexChild + 1] = 1
					break
				}
			} else if itemChild == 1 && directions[index] == "down" {
				if indexChild == 0 {
					scanner[index][indexChild] = 0
					scanner[index][indexChild + 1] = 1
					directions[index] = "up"
					break
				} else {
					scanner[index][indexChild] = 0
					scanner[index][indexChild - 1] = 1
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