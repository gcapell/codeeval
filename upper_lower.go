package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func upper_lower(line string) {
	var upper, lower int
	for _, r := range line {
		if !unicode.IsLetter(r) {
			continue
		}
		if unicode.IsUpper(r) {
			upper++
		} else {
			lower++
		}
	}
	percent := float64(100) / float64(upper+lower)
	fmt.Printf("lowercase: %.2f uppercase: %.2f\n", float64(lower)*percent, float64(upper)*percent)
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
		upper_lower(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
