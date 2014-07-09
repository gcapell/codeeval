package main

import (
	"fmt"
	"math"
)

func main() {
	for p := range palindromes() {
		if isPrime(p) {
			fmt.Println(p)
			break
		}
	}
}

func isPrime(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	for j := 2; j < sqrt; j++ {
		if n%j == 0 {
			return false
		}
	}
	return true
}

func palindromes() chan int {
	c := make(chan int)
	go func() {
		for a := 9; a > 0; a-- {
			for b := 9; b > 0; b-- {
				c <- a*100 + b*10 + a
			}
		}
	}()
	return c
}
