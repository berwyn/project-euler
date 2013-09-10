package main

import "fmt"

func main() {
	sum, three_val, five_val := 0, 0, 0
	for i := 0; (three_val < 1000 || five_val < 1000); i++ {
		three_val, five_val = (i * 3), (i * 5)
		if three_val < 1000 { sum += three_val }
		if five_val < 1000 { sum += five_val }
	}
	fmt.Println(sum)
}