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
		manhattan(line)
	}
}

type event struct {
	up bool
	y  int
}

func manhattan(line string) {
	triples := strings.Split(line, ";")
	events := make(map[int][]event)
	for _, t := range triples {
		var x0, y, x1 int
		_, err := fmt.Sscanf(strings.TrimSpace(t), "(%d,%d,%d)", &x0, &y, &x1)
		if err != nil {
			log.Fatal(err, line)
		}
		events[x0] = append(events[x0], event{true, y})
		events[x1] = append(events[x1], event{false, y})
	}

	eventLocations := make([]int, 0, len(events))
	for x := range events {
		eventLocations = append(eventLocations, x)
	}
	sort.Ints(eventLocations)

	heights := make([]int, 0)
	lastY := 0
	for pos, x := range eventLocations {
		for _, e := range events[x] {
			if e.up {
				insert(&heights, e.y)
			} else {
				remove(&heights, e.y)
			}
		}
		y := biggest(heights)
		if y != lastY {
			if pos == 0 {
				fmt.Printf("%d %d", x, y)
			} else {
				fmt.Printf(" %d %d", x, y)
			}
			lastY = y
		}
	}
	fmt.Println()
}

func insert(s *[]int, x int) {
	i := sort.SearchInts(*s, x)
	if i == len(*s) {
		*s = append(*s, x)
		return
	}
	*s = append(*s, 0)
	copy((*s)[i+1:], (*s)[i:])
	(*s)[i] = x
}

func remove(s *[]int, x int) {
	i := sort.SearchInts(*s, x)
	*s = append((*s)[:i], (*s)[i+1:]...)
}

func biggest(s []int) int {
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return 0
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
