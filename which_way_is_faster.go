package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type (
	point   struct{ x, y int }
	journey struct {
		point
		voyaging bool
	}

	board struct {
		w, h           int
		port, mountain map[point]bool
		start, finish  point
		voyageComplete bool
	}
)

func (j journey) String() string {
	s := fmt.Sprintf("%d,%d", j.x, j.y)
	if j.voyaging {
		return s + "!"
	}
	return s
}

func (b board) tripLength() int {
	distance := 0
	start := journey{b.start, false}
	d0 := []journey{start}
	d1 := []journey{}
	d2 := []journey{}
	visited := map[journey]bool{}
	for {
		if len(d0) == 0 && len(d1) == 0 {
			return -1
		}
		d0 = unvisited(visited, d0)
		//log.Printf("%v", d0)
		for _, j := range d0 {
			if j.point == b.finish && !j.voyaging {
				return distance
			}
			n1, n2 := b.next(j, visited)
			d1 = append(d1, n1...)
			d2 = append(d2, n2...)
		}
		d0, d1, d2 = d1, d2, nil
		distance++
	}
}

func unvisited(visited map[journey]bool, js []journey) []journey {
	var reply []journey
	dup := map[journey]bool{}
	for _, j := range js {
		if visited[j] || dup[j] {
			continue
		}

		visited[j] = true
		dup[j] = true
		reply = append(reply, j)

	}
	return reply
}

func (b *board) next(j journey, visited map[journey]bool) ([]journey, []journey) {
	var n1, n2 []journey

	if b.port[j.point] {
		if j.voyaging {
			n1 = append(n1, journey{j.point, false})
			b.voyageComplete = true
		} else {
			n1 = append(n1, journey{j.point, true})
		}
		b.port[j.point] = false
	}
	if j.voyaging {
		if !b.voyageComplete {
			n1 = append(n1, b.neighbours(j)...)
		}
	} else {
		n2 = append(n2, b.neighbours(j)...)
	}
	return n1, n2
}

func (b board) neighbours(o journey) []journey {
	var reply []journey
	for _, p := range []point{{o.x - 1, o.y}, {o.x + 1, o.y}, {o.x, o.y - 1}, {o.x, o.y + 1}} {
		if b.mountain[p] || p.x < 0 || p.y < 0 || p.x >= b.w || p.y >= b.h {
			continue
		}
		reply = append(reply, journey{p, o.voyaging})
	}
	return reply
}

func parse(line string) board {
	rows := strings.Split(line, " | ")
	letter := make(map[point]rune)
	b := board{w: len(rows[0]), h: len(rows),
		port:     make(map[point]bool),
		mountain: make(map[point]bool),
	}
	for r, row := range rows {
		for c, v := range row {
			p := point{r, c}
			switch v {
			case 'S':
				b.start = p
			case 'F':
				b.finish = p
			case 'P':
				b.port[p] = true
			case '^':
				b.mountain[p] = true
			}
			letter[p] = v
		}
	}
	return b
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
		fmt.Println(parse(scanner.Text()).tripLength())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
