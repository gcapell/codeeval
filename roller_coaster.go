package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func roller(line string) string {
	up := true

	runes := make([]rune, 0, len(line))
	for _, r := range line {
		if unicode.IsLetter(r) {
			if up {
				r = unicode.ToUpper(r)
			} else {
				r = unicode.ToLower(r)
			}
			up = !up
		}
		runes = append(runes, r)
	}
	return string(runes)
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
		fmt.Println(roller(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
