package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countZeroes(line string) int {
	var zeroes, max int
	if _, err := fmt.Sscanf(line, "%d %d", &zeroes, &max); err != nil {
		log.Fatal(err)
	}
	count := 0
	for j := 1; j <= max; j++ {
		if nZeroes(j, zeroes) {
			count++
		}
	}
	return count
}

func nZeroes(n, z int) bool {
	for n > 1 {
		if n%2 == 0 {
			z--
			if z < 0 {
				return false
			}
		}
		n /= 2
	}
	return z == 0
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
		fmt.Println(countZeroes(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
