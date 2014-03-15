package main

import (
	"fmt"
)

func main() {
	i := 20
	for {
		matchFound := true
		for j := 2; j <= 20; j++ {
			matchFound = matchFound && i%j == 0
		}
		if matchFound {
			break
		} else {
			i++
		}
	}

	fmt.Println(i)
}
