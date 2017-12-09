package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "" //Here put the input
	tempInput := strings.Split(input, "\n")
	registers := make(map[string]int)
	maxEver := 0
	for _, item := range tempInput {
		splittedItem := strings.Split(item, " if ")
		operation := strings.Split(splittedItem[0], " ")
		condition := strings.Split(splittedItem[1], " ")
		if operation[1] == "inc" && checkCondition(condition, registers) {
			numberToAdd, _ := strconv.Atoi(operation[2])
			registers[operation[0]] += numberToAdd
		} else if operation[1] == "dec" && checkCondition(condition, registers) {
			numberToSubstract, _ := strconv.Atoi(operation[2])
			registers[operation[0]] -= numberToSubstract
		}
		if maxEver < registers[operation[0]] {
			maxEver = registers[operation[0]]
		}
	}
	fmt.Println(maxEver)
}

func maxMin(registers map[string]int) (int, int) {
	firstTime := true
	var max int
	var min int
	for _, item := range registers {
		if firstTime {
			max = item
			min = item
			firstTime = false
		}
		if item > max {
			max = item
		}
		if item < min {
			min = item
		}
	}
	return max, min
}

func checkCondition(condition []string, valuesMap map[string]int) bool {
	leftSide := valuesMap[condition[0]]
	rightSide, _ := strconv.Atoi(condition[2])
	switch condition[1] {
	case "==":
		return leftSide == rightSide

	case "!=":
		return leftSide != rightSide

	case "<=":
		return leftSide <= rightSide

	case ">=":
		return leftSide >= rightSide

	case "<":
		return leftSide < rightSide

	case ">":
		return leftSide > rightSide
	default:
		return false
	}
	return false
}
