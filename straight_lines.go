package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type point struct{ x, y int }

func parse(line string) []point {
	var reply []point
	for _, s := range strings.Split(line, " | ") {
		var p point
		if _, err := fmt.Sscanf(s, "%d %d", &p.x, &p.y); err != nil {
			log.Fatal(err, line)
		}
		reply = append(reply, p)
	}
	return reply
}

type pair struct{ a, b int }

func lines(p []point) int {
	lines := 0

	seen := make(map[pair]bool)

	for i := 0; i < len(p); i++ {
		for j := i + 1; j < len(p); j++ {
			for k := j + 1; k < len(p); k++ {
				if colinear(p[i], p[j], p[k]) {
					if !seen[pair{i, j}] {
						lines++
						seen[pair{i, j}] = true
					}
					seen[pair{j, k}] = true
				}
			}
		}
	}
	return lines
}

func colinear(a, b, c point) bool {
	return (a.y-b.y)*(a.x-c.x) == (a.y-c.y)*(a.x-b.x)
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
		fmt.Println(lines(parse(scanner.Text())))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
