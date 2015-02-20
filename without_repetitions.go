package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func withoutRepetitions(line string) string {
	var prev rune
	return strings.Map(func(r rune) rune {
		if r == prev {
			return -1
		}
		prev=r
		return r
	}, line)
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
		fmt.Println(withoutRepetitions(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
