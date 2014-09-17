package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func track(line string, lastpos int) (string, int) {
	next := access(line, lastpos)
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
	pos := -1
	for scanner.Scan() {
		line, pos := track(scanner.Text(), pos)
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
