package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var primes = []float64{2, 3, 5, 7, 11, 13}

func mersenne(n int) {
	first := true
	for _, p := range primes {
		m := int(math.Exp2(p)) - 1
		if m >= n {
			break
		}
		if !first {
			fmt.Print(", ")
		}
		first = false
		fmt.Print(m)
	}
	fmt.Println()
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expected filename")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		mersenne(atoi(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
