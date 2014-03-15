package main

import (
	"fmt"
	"math"
)

func main() {
	n := float64(600851475143)
	fmt.Println(trialDivision(n))
}

func trialDivision(n float64) []float64 {
	if n == 1 {
		return []float64{n}
	}

	primeFactors := []float64{}
	primeChan := make(chan int)
	primes := primeSieve(primeChan)

	primeChan <- int(math.Pow(n, 0.5) + 1)
	close(primeChan)

	for i := range primes {
		p := float64(i)
		if math.Pow(p, 2) > n {
			break
		}
		for math.Mod(n, p) == 0 {
			primeFactors = append(primeFactors, n)
			n /= p
		}
	}

	if n > 1 {
		primeFactors = append(primeFactors, n)
	}

	return primeFactors
}

func primeSieve(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			A := make([]bool, n+1)
			for i := 2; i < len(A); i++ {
				A[i] = true
			}

			for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
				if A[i] {
					for j := int(math.Pow(float64(i), 2)); j <= n; j += i {
						A[j] = false
					}
				}
			}

			for i := 2; i < len(A); i++ {
				if A[i] {
					out <- i
				}
			}
		}
		close(out)
	}()

	return out
}
