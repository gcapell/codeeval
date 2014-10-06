package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var direction = map[int]string{
	-1: "/",
	0:  "|",
	1:  `\`,
}

func steer(line string, last int) (string, int) {
	gate, checkpoint := strings.IndexByte(line, '_'), strings.IndexByte(line, 'C')
	next := choose(last, gate, checkpoint)
	if last == -1 {
		last = next
	}
	reply := line[:next] + direction[next-last] + line[next+1:]
	return reply, next
}

func choose(last, gate, checkpoint int) int {
	if last == -1 {
		if checkpoint != -1 {
			return checkpoint
		}
		return gate
	}
	if near(last, checkpoint) {
		return checkpoint
	}
	return gate
}

func near(a, check int) bool {
	if check == -1 {
		return false
	}
	if a > check {
		return (a - check) < 2
	}
	return (check - a) < 2
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
	lastRow := -1
	out := ""
	for scanner.Scan() {
		out, lastRow = steer(scanner.Text(), lastRow)
		fmt.Println(out)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
