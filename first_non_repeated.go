package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	for line := range linesFromFilename() {
		fmt.Println(firstNonRepeated([]byte(line)))
	}
}

func firstNonRepeated(line []byte) string {
	pos := make(map[byte]int)

	for p, r := range line {
		if pos[r] == 0 {
			// first time
			pos[r] = p + 1
		} else {
			// eliminate
			pos[r] = -1
		}
	}

	sortable := []byte(line)
	sort.Sort(ByPos{pos, sortable})
	return string(sortable[:1])
}

type ByPos struct {
	pos  map[byte]int
	line []byte
}

func (a ByPos) Len() int      { return len(a.line) }
func (a ByPos) Swap(i, j int) { a.line[i], a.line[j] = a.line[j], a.line[i] }
func (a ByPos) Less(i, j int) bool {
	if a.pos[a.line[i]] == -1 {
		return false
	}
	if a.pos[a.line[j]] == -1 {
		return true
	}
	return a.pos[a.line[i]] < a.pos[a.line[j]]
}

func linesFromFilename() chan string {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan string)

	go func() {
		reader := bufio.NewReader(f)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					log.Fatal(err)
				}
				break
			}
			c <- strings.TrimSpace(line)
		}
		close(c)
	}()
	return c
}
