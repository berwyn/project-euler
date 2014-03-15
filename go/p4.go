package main

import (
	"fmt"
	"strconv"
)

func main() {
	largestNum := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			number := i * j
			numString := strconv.Itoa(number)
			numberReversed := reverse(numString)
			if numString == numberReversed && number > largestNum {
				largestNum = number
			}
		}
	}
	fmt.Println(largestNum)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
