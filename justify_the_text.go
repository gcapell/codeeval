package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func justify(line string) {
	for len(line) > 80 {
		var before string
		before, line = split(line)
		fmt.Println(justified(before))
	}
	fmt.Println(line)
}

func split(line string) (before, after string) {
	pos := strings.LastIndex(line[:81], " ")
	if pos == -1 {
		log.Fatal(line, "no space in first 80 chars")
	}
	before = line[:pos]
	after = line[pos+1:]
	if after[0] == ' ' {
		log.Fatal(after)
	}
	return line[:pos], line[pos+1:]
}

func justified(line string) string {
	words := strings.Fields(line)
	if len(words) == 1 {
		return words[0]
	}
	space := 80
	for _, w := range words {
		space -= len(w)
	}
	slots := len(words) - 1
	average := space / slots
	spaces := string(bytes.Repeat([]byte(" "), average))
	extras := space - (average * slots)

	var reply string
	for j, w := range words {
		if j != 0 {
			reply += spaces
			if j <= extras {
				reply += " "
			}
		}
		reply += w
	}
	return reply
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
		justify(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
