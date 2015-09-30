package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func possible(line string) int {
	left, right := line[:len(line)/2], line[len(line)/2:]

	variants := 1
	for j := range left {
		if left[j] == '*' || right[j] == '*' {
			if left[j] == right[j] {
				variants *= 2
			}
		} else {
			if left[j] != right[j] {
				return 0
			}
		}
	}
	return variants
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
		fmt.Println(possible(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
