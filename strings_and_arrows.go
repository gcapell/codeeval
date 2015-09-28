package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func arrows(line string) int {
	return countOverlapping(line, ">>-->") + countOverlapping(line, "<--<<")
}

func countOverlapping(line, pattern string) int {
	c := 0
	for {
		pos := strings.Index(line, pattern)
		if pos == -1 {
			return c
		}
		c++
		line = line[pos+1:]
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
		fmt.Println(arrows(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
