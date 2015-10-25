package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type edge struct {
	start      int
	horizontal bool
}

// Additional top and left edges of a square of 'size' with 'e' as top-left edge
// (relative to a square of 'size-1').
func (e edge) topLeft(size int) []edge {
	return []edge{
		{e.start + size - 1, true},
		{e.start + 5*(size-1), false},
	}
}

// Bottom and right edges of a square of 'size' with 'e' as top-left edge?
func (e edge) bottomRight(size int) []edge {
	var reply []edge
	// bottom
	bottom := e.start + 5*size
	right := e.start + size
	for j := 0; j < size; j++ {
		reply = append(reply, edge{bottom, true}, edge{right, false})
		bottom++
		right += 5
	}
	return reply
}

func parse(line string) map[edge]bool {
	edges := make(map[edge]bool)
	chunks := strings.Split(line, " | ")
	for _, c := range chunks {
		var a, b int
		if _, err := fmt.Sscanf(c, "%d %d", &a, &b); err != nil {
			log.Fatal(err, c, chunks, line)
		}
		if a > b {
			a, b = b, a
		}
		edges[edge{a, b == a+1}] = true
	}
	return edges
}

func all(list []edge, good map[edge]bool) bool {
	for _, e := range list {
		if !good[e] {
			return false
		}
	}
	return true
}

func squares(edges map[edge]bool) int {
	count := 0
	for e := range edges {
		if !e.horizontal {
			continue
		}
		for size := 1; ; size++ {
			if all(e.topLeft(size), edges) {
				if all(e.bottomRight(size), edges) {
					count++
				}
			} else {
				break
			}
		}
	}
	return count
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
		fmt.Println(squares(parse(scanner.Text())))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
