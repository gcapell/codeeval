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
		fmt.Println(maxBeauty(line))
	}
}

func maxBeauty(s string) int {
	counts := make(map[rune]int)

	for _, c := range strings.ToLower(s) {
		if c >= 'a' && c <= 'z' {
			counts[c] += 1
		}
	}

	totals := make([]int, 0, len(counts))
	for _, n := range counts {
		totals = append(totals, n)
	}
	sort.Sort(Reversed(totals))

	score := 26
	total := 0
	for _, n := range totals {
		total += n * score
		score -= 1
	}

	return total

}

type Reversed []int

func (a Reversed) Len() int           { return len(a) }
func (a Reversed) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Reversed) Less(i, j int) bool { return a[i] > a[j] }

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
