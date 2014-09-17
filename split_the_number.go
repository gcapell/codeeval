package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func split(r *regexp.Regexp, line string) int {
	chunks := strings.Fields(line)
	n, pat := chunks[0], chunks[1]
	m := r.FindStringSubmatch(pat)
	left, op, right := m[1], m[2], m[3]
	nl, nr := extract(n, left), extract(n, right)
	if op == "+" {
		return nl + nr
	}
	return nl - nr
}

func extract(line, p string) int {
	total := 0
	for _, c := range p {
		i := c - 'a'
		n := line[i] - '0'
		total = total*10 + int(n)

	}
	return total
}

func main() {
	r := regexp.MustCompile("([a-z]+)([-+])([a-z]+)")
	if len(os.Args) != 2 {
		log.Fatal("expected filename")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(split(r, scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
