package main

import (
	"fmt"
	"strconv"
	"strings"
)


func main() {
	input := "b inc 5 if a > 1\na inc 1 if b < 5\nc dec -10 if a >= 1\nc inc -20 if c == 10" //Here put the input
	tempInput := strings.Split(input, "\n")
	registers := make(map[string]int)
	for _, item := range tempInput {
		splittedItem := strings.Split(item, " if ")
		operation := strings.Split(splittedItem[0], " ")
		condition := strings.Split(splittedItem[1], " ")
		if operation[1] == "inc" && checkCondition(condition, registers){
			numberToAdd, _ := strconv.Atoi(operation[2])
			registers[operation[0]] += numberToAdd
		} else if operation[1] == "dec" && checkCondition(condition, registers){
			numberToSubstract, _ := strconv.Atoi(operation[2])
			registers[operation[0]] -= numberToSubstract
		}
	}
	fmt.Println(registers)
}

func checkCondition (condition []string, valuesMap map[string]int) bool{
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