package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type vec struct{ x, y int }

var directions = []vec{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}

func nuts(line string) int {
	var box, target vec
	fmt.Sscanf(line, "%dx%d | %d %d", &box.x, &box.y, &target.x, &target.y)

	pos := vec{1, box.y}
	direction := 0
	score := 1
	delta := directions[direction]
	visited := make(map[vec]bool)
	visited[pos] = true
	for {
		pos.add(delta)
		if visited[pos] || pos.outside(box) {
			pos.sub(delta)
			direction = (direction + 1) % len(directions)
			delta = directions[direction]
			continue
		}
		score++
		if pos == target {
			return score
		}
		visited[pos] = true
	}
}

func (p vec) outside(b vec) bool {
	return p.x == 0 || p.y == 0 || p.x > b.x || p.y > b.y
}

func (p *vec) add(o vec) {
	p.x += o.x
	p.y += o.y
}

func (p *vec) sub(o vec) {
	p.x -= o.x
	p.y -= o.y
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
		fmt.Println(nuts(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
