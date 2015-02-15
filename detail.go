package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func detail(line string) int {
	rows := strings.Split(line, ",")
	first := true
	minDots := 0
	for _, r := range rows {
		dots := len(strings.Trim(r, "XY"))
		if first || dots < minDots {
			first = false
			minDots = dots
		}
	}
	return minDots
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
		fmt.Println(detail(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
