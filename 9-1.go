package main

import (
	"fmt"
	// "strconv"
	"strings"
)

func main() {
	input := "" //Here put the input
	inputSlice := strings.Split(input, "")
	inputSlice = clearCommentBrackets(clearExclamation(inputSlice))
	level := 0
	score := 0
	for _, item := range inputSlice {
		if item == "{" {
			level += 1
			score += level
		} else if item == "}" {
			level -= 1
		}
	}
	fmt.Println(score)
}




func deleteByIndexRange(indexOne int, indexTwo int, input []string) []string {
	return append(input[:indexOne], input[indexTwo+1:]...)
}

func clearExclamation(input []string) []string {
	for {
		var currentEnd int
		for index, item := range input{
			currentEnd = index
			if item == "!"	{
				input = deleteByIndexRange(index, index + 1, input)
				break
			}
		}
		if currentEnd == (len(input) - 1){
			break
		}
	}
	return input
}

func clearCommentBrackets(input []string) []string {
	var currentEnd int
	for {
		for index, item := range input {
			currentEnd = index
			if item == "<" {
				indexStart := index 
				indexEnd := len(input) - 1
				for indexChild, itemChild := range input[indexStart:] {
					if itemChild == ">" {
						indexEnd = indexStart + indexChild
						break
					}
				}
				input = deleteByIndexRange(indexStart, indexEnd, input)
				break
			}
		}
		if currentEnd == (len(input) - 1) || len(input) == 0{
			break
		}
	}
	return input
}