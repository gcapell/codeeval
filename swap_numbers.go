package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var pattern = regexp.MustCompile("([0-9])([A-Za-z]+)([0-9])")

func swap(s string) string {
	return string(pattern.ReplaceAll([]byte(s), []byte("$3$2$1")))
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
		fmt.Println(swap(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
