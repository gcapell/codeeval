package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func luhn(s string) bool {
	total := 0
	for pos, c := range strings.Join(strings.Fields(s), "") {
		n := int(c - '0')
		if pos%2 == 0 {
			n *= 2
		}
		total = (total + n) % 10
	}
	return total == 0
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
		s := "Fake"
		if luhn(scanner.Text()) {
			s = "Real"
		}
		fmt.Println(s)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
