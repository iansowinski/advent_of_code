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
	// fmt.Println(allPrograms)

	fmt.Println(findFirstParent(allPrograms))
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
