package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func matchingNames(line string) string {
	split := strings.Split(line, "|")
	wines, pattern := strings.Fields(split[0]), letterCount(strings.Fields(split[1])[0])
	var matches []string
	for _, w := range wines {
		if match(letterCount(w), pattern) {
			matches = append(matches, w)
		}
	}
	if len(matches) == 0 {
		return "False"
	}
	return strings.Join(matches, " ")
}

func letterCount(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}
	return counts
}

func match(w, pattern map[rune]int) bool {
	for r, n := range pattern {
		if w[r] < n {
			return false
		}
	}
	return true
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
		fmt.Println(matchingNames(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
