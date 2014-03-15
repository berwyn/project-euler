package main

import (
	"fmt"
	"math"
)

func main() {
	sumOfSquares, squareOfSum := 0, 0
	for i := 0; i <= 100; i++ {
		sumOfSquares += int(math.Pow(float64(i), 2))
		squareOfSum += i
	}
	ans := int(math.Pow(float64(squareOfSum), 2)) - sumOfSquares
	fmt.Println(ans)
}
