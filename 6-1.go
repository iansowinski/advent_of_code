package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "4	1	15	12	0	9	9	5	5	8	7	3	14	5	12	3" //Here put the input
	// input := "0	2	7	0"
	tempBanks := strings.Split(input, "\t")
	banks := make([]int, len(tempBanks))
	for index, item := range tempBanks {
		banks[index], _ = strconv.Atoi(item)
	}
	banksHistory := make([]string, 0)
	i := 0

	for !checkDuplicates(banksHistory, join(banks), i) {
		banksHistory = append(banksHistory, join(banks))
		i += 1
		maxIndex, maxNumber := getMax(banks)
		banks[maxIndex] = 0
		for j := 0; j < maxNumber; j++ {
			banks[(maxIndex+j+1)%len(banks)] += 1
		}
	}
	fmt.Println(i)
}

func getMax(array []int) (int, int) {
	var max int = array[0]
	var maxIndex int = 0
	for index, number := range array {
		if number > max {
			max = number
			maxIndex = index
		}
	}
	return maxIndex, max
}

func checkDuplicates(array []string, itemToCheck string, itemToCheckIndex int) bool {
	for index, number := range array {
		if index != itemToCheckIndex && number == itemToCheck && number != "" {
			return true
		}
	}
	return false
}

func join(array []int) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(array), " "), ","), "[]")
}
