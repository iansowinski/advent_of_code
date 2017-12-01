package main

import (
  "fmt"
  "strings"
  "strconv"
)

func main () {
	input := "" //Here put the input
  inputArray := strings.Split(input, "")
  result := 0
  halfWay := (len(inputArray)) / 2
  for index, number := range inputArray {
    if index + halfWay <= len(inputArray) - 1 {
      if number == inputArray[index + halfWay] {
        numberInt, _ := strconv.Atoi(number)
        result += numberInt
      }
    } else {
    	tempIndex := index + halfWay - len(inputArray)
      if number == inputArray[tempIndex] {
        numberInt, _ := strconv.Atoi(number)
        result += numberInt
      }
    }
  }
  fmt.Println(result)
}
