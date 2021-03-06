package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readMore(line string) string {
	if len(line) <= 55 {
		return line
	}
	trimmed := line[:40]
	space := strings.LastIndex(trimmed, " ")
	if space == -1 {
		return trimmed
	}
	return trimmed[:space] + "... <Read More>"
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
		fmt.Println(readMore(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
