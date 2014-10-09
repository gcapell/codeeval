package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var romanBase = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func aromatic(line string) int {
	total := 0
	prevRVal := 0
	for len(line) > 0 {
		a, r := line[len(line)-2], line[len(line)-1]
		line = line[:len(line)-2]
		aVal := a - '0'
		rVal := romanBase[r]
		val := int(aVal) * rVal
		if rVal < prevRVal {
			val = -val
		}
		prevRVal = rVal
		total += val
	}
	return total
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
		fmt.Println(aromatic(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
