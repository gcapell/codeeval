package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func solveMaze(line string) int {
	return parse(line).solve()
}

type coord [3]uint8

type queue str
type maze struct {
	cells [][][]rune
	entrance, exit coord
	depthMarker coord
}

func parse(line) maze {
	
}

// solve returns length of shortest path.
func (m*maze)solve() int {
	// BFS
	var frontier queue
	visited := make(map[coord]bool)
	depth := 0
	frontier.push(m.entrance)
	frontier.push(m.depthMarker)
	for {
		c := frontier.pop()
		if c == m.depthMarker {
			depth++
			frontier.push(m.depthMarker)
			continue
		}
		for n := m.neighbours(c) {
			if n == m.exit {
				return depth
			}
			if !visited[n] {
				visited[n] = true
				frontier.push(n)
			}
		}
	}
	
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
		fmt.Println(solveMaze(line))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
