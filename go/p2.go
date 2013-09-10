package main

import "fmt"

func main() {
	first_fib, second_fib, sum, new_fib := 1, 2, 2, 0
	for {
		new_fib = first_fib + second_fib
		if new_fib > 4000000 {
			break
		}
		if new_fib % 2 == 0 {
			sum += new_fib
		}
		first_fib, second_fib = second_fib, new_fib
	}
	fmt.Println(sum)
}