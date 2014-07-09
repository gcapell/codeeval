package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func major_element(line string) string {
	counts := make(map[string]int)
	for _, n := range strings.Split(line, ",") {
		counts[n] += 1
	}

	var maxElem string
	var maxCount int

	total := 0
	for elem, n := range counts {
		total += n
		if n > maxCount {
			maxCount = n
			maxElem = elem
		}
	}
	if maxCount >= total/2 {
		return maxElem
	} else {
		return "None"
	}
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
		fmt.Println(major_element(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
