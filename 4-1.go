package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "" //Here put the input

	validAmount := 0
	for _, passPraseValid := range checkPassphrases(input) {
		if passPraseValid == true {
			validAmount += 1
		}
	}
	fmt.Println(validAmount)
}

func checkPassphrases(input string) []bool {
	inputTempArray := strings.Split(input, "\n")
	lengthTempArray := len(inputTempArray)
	resultsArray := make([]bool, lengthTempArray)
	for index, passPrase := range inputTempArray {
		resultsArray[index] = passphraseCorrect(passPrase)
	}
	return resultsArray
}

func passphraseCorrect(input string) bool {
	inputArray := strings.Split(input, " ")
	for index, word := range inputArray {
		for indexSecond, wordSecond := range inputArray {
			if indexSecond != index && word == wordSecond {
				return false
			}
		}
	}
	return true
}
