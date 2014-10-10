package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type node uint

var (
	words []string
	start map[byte][]node
)

func chain(line string) string {
	words = strings.Split(line, ",")
	start = make(map[byte][]node)
	for pos, w := range words {
		s := w[0]
		start[s] = append(start[s], node(pos))
	}

	length := 0
	for pos := range words {
		l := lengthFrom(node(pos), 0, 0)
		if l > length {
			length = l
		}
	}
	if length == 0 {
		return "None"
	}
	return strconv.Itoa(length + 1)
}

type bitmap uint64

func (b bitmap) isSet(n node) bool {
	return b&(1<<n) != 0
}

func (b bitmap) set(n node) bitmap {
	return b | (1 << n)
}

func lengthFrom(n node, visited bitmap, depth int) int {
	w := words[n]
	visited = visited.set(n)
	longest := 0
	for _, next := range start[w[len(w)-1]] {
		if !visited.isSet(next) {
			l := 1 + lengthFrom(next, visited, depth+1)
			if l > longest {
				longest = l
			}
		}
	}
	return longest
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
		fmt.Println(chain(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
