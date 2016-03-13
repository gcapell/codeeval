package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func compare(line string) string {
	chunks := strings.Split(line, "|")
	cards := strings.Fields(chunks[0])
	l, r := cards[0], cards[1]
	trump := strings.TrimSpace(chunks[1])
	ls, rs := score(l, trump), score(r, trump)
	switch {
	case ls > rs:
		return l
	case rs > ls:
		return r
	default:
		return strings.Join(cards, " ")
	}
}

var values = map[byte]int{
	'2': 2, '3': 3, '4': 4, '5': 5, '6': 6,
	'7': 7, '8': 8, '9': 9, '1': 10,
	'J': 11, 'Q': 12, 'K': 13, 'A': 14,
}

func score(c string, trump string) int {
	n := values[c[0]]
	if strings.HasSuffix(c, trump) {
		n += 20
	}
	return n
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
		fmt.Println(compare(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
