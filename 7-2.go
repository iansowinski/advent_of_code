package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Program struct {
	name            string
	weight          int
	children        []Program
	childrenStrings []string
}

func main() {
	input := "" //Here put the input
	tempInput := strings.Split(input, "\n")
	allPrograms := make([]Program, 0)
	for _, item := range tempInput {
		var newProgram Program
		firstSplit := strings.Split(item, ") -> ")
		secondSplit := strings.Split(firstSplit[0], " (")
		newProgram.name = secondSplit[0]
		if len(firstSplit) > 1 {
			newProgram.weight, _ = strconv.Atoi(secondSplit[1])
			newProgram.childrenStrings = strings.Split(firstSplit[1], ", ")
		} else {
			number := strings.Split(secondSplit[1], ")")
			newProgram.weight, _ = strconv.Atoi(number[0])
		}
		allPrograms = append(allPrograms, newProgram)
	}
	firstParent := findFirstParent(allPrograms)
	allPrograms = deleteByItem(firstParent, allPrograms)
	firstParent, _ = growTree(firstParent, allPrograms)
	fmt.Println(findDefect(firstParent).weight + findDifference(firstParent))
}

func findDefect(program Program) Program {
	if len(program.children) == 0 {
		return program
	}
	rightWeight := recurrentSum(program.children[0])
	defect := program.children[0]
	for index, item := range program.children {
		if index != 0 {
			if rightWeight == recurrentSum(item) {
				break
			} else {
				rightWeight = recurrentSum(item)
			}
		}
	}
	for _, item := range program.children {
		if rightWeight != recurrentSum(item) {
			defect = item
			break
		}
	}
	if recurrentSum(defect) != rightWeight {
		return findDefect(defect)
	}
	return program
}

func findDifference(program Program) int {
	rightWeight := recurrentSum(program.children[0])
	defect := program.children[0]
	for index, item := range program.children {
		if index != 0 {
			if rightWeight == recurrentSum(item) {
				break
			} else {
				rightWeight = recurrentSum(item)
			}
		}
	}
	for _, item := range program.children {
		if rightWeight != recurrentSum(item) {
			defect = item
			break
		}
	}
	return rightWeight - recurrentSum(defect)
}

func recurrentSum(program Program) int {
	sum := program.weight
	if len(program.children) != 0 {
		for _, child := range program.children {
			sum += recurrentSum(child)
		}
	}
	return sum
}

func growTree(parent Program, allPrograms []Program) (Program, []Program) {
	if len(parent.childrenStrings) == 0 {
		return parent, allPrograms
	}
	for _, itemParent := range parent.childrenStrings {
		for _, item := range allPrograms {
			if item.name == itemParent {
				clearedAllPrograms := deleteByItem(item, allPrograms)
				var childToAppend Program
				childToAppend, allPrograms = growTree(item, clearedAllPrograms)
				parent.children = append(parent.children, childToAppend)
			}
		}
	}
	return parent, allPrograms
}

func deleteByItem(itemToDelete Program, array []Program) []Program {
	indexToDelete := 0
	for index, item := range array {
		if itemToDelete.name == item.name {
			indexToDelete = index
			break
		}
	}
	array = array[:indexToDelete+copy(array[indexToDelete:], array[indexToDelete+1:])]
	return array
}

func findFirstParent(allPrograms []Program) Program {
	var lastParent Program
	currentParent := allPrograms[0]
	for !(lastParent.name == currentParent.name) {
		lastParent = currentParent
		currentParent = findParent(allPrograms, currentParent)
	}
	return currentParent
}

func findParent(allPrograms []Program, currentParent Program) Program {
	for _, item := range allPrograms {
		if currentParent.name != item.name && len(item.childrenStrings) > 0 {
			for _, currentChild := range item.childrenStrings {
				if currentChild == currentParent.name {
					return item
				}
			}
		}
	}
	return currentParent
}
