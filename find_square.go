package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	for line := range linesFromFilename() {
		fmt.Println(isSquare(parse(line)))
	}
}

type point struct{ x, y float64 }

func (p point) equals(p2 point) bool {
	dist := math.Sqrt(
		math.Pow(p.x-p2.x, 2) +
			math.Pow(p.y-p2.y, 2))
	return dist < 0.01
}

func parse(line string) []point {
	pRE := regexp.MustCompile(`\(([0-9]+),([0-9]+)\)`)
	matches := pRE.FindAllStringSubmatch(line, -1)
	if len(matches) != 4 {
		log.Fatal("expected 4 points, got %q", matches)
	}
	reply := make([]point, len(matches))
	for pos, match := range matches {
		reply[pos] = point{atof(match[1]), atof(match[2])}
	}
	return reply
}

func isSquare(p []point) bool {
	c := point{
		(p[0].x + p[1].x + p[2].x + p[3].x) / 4,
		(p[0].y + p[1].y + p[2].y + p[3].y) / 4,
	}
	dx := p[0].x - c.x
	dy := p[0].y - c.y

	if dx*dx+dy*dy < 0.01 {
		return false
	}

	predicted := []point{
		point{c.x + dy, c.y - dx},
		point{c.x - dx, c.y - dy},
		point{c.x - dy, c.y + dx},
	}
	return samePoints(predicted, p[1:])
}

func samePoints(a, b []point) bool {
	if len(a) != len(b) {
		return false
	}
	used := make(map[int]bool)
search:
	for _, p := range a {
		for pos, p2 := range b {
			if used[pos] {
				continue
			}
			if p.equals(p2) {
				used[pos] = true
				continue search
			}
		}

		return false
	}
	return true
}

func atof(s string) float64 {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return float64(n)
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
