package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord struct{ x, y, z uint8 }

type cell bool // is hole

type maze struct {
	cells                       map[coord]cell // is hole
	entrance, exit, depthMarker coord
}

func (m maze) String() string {
	b := &bytes.Buffer{}
	for z := uint8(0); z < size; z++ {
		fmt.Fprintln(b, "Level", z+1)
		for y := uint8(0); y < size; y++ {
			for x := uint8(0); x < size; x++ {
				r := '*'
				here := coord{x, y, z}
				if c, ok := m.cells[here]; ok {
					switch {
					case here == m.entrance:
						r = 'e'
					case here == m.exit:
						r = 'E'
					case bool(c):
						r = 'o'
					default:
						r = ' '
					}
				}
				b.WriteRune(r)
			}
			b.WriteRune('\n')
		}
	}
	return b.String()
}

var size uint8

func (c coord) isEdge() bool {
	if !(c.z == 0 || c.z == size-1) {
		return false
	}
	return c.x == 0 || c.y == 0 || c.x == size-1 || c.y == size-1
}

func (c coord) neighbours(cells map[coord]cell) []coord {
	var p []coord
	if c.x > 0 {
		p = append(p, coord{c.x - 1, c.y, c.z})
	}
	if c.x+1 < size {
		p = append(p, coord{c.x + 1, c.y, c.z})
	}
	if c.y > 0 {
		p = append(p, coord{c.x, c.y - 1, c.z})
	}
	if c.y+1 < size {
		p = append(p, coord{c.x, c.y + 1, c.z})
	}
	if c.z > 0 && cells[c] {
		p = append(p, coord{c.x, c.y, c.z - 1})
	}

	var reply []coord
	for _, d := range p {
		if _, ok := cells[d]; ok {
			reply = append(reply, d)
		}
	}

	if c.z+1 < size {
		d := coord{c.x, c.y, c.z + 1}
		if cells[d] {
			reply = append(reply, d)
		}
	}
	return reply
}

func iToCoord(n int) coord {
	x, n := uint8(n%int(size)), n/int(size)
	y, z := uint8(n%int(size)), uint8(n/int(size))
	return coord{x, y, z}
}

func parse(line string) maze {
	chunks := strings.Split(line, ";")
	iSize, err := strconv.Atoi(chunks[0])
	if err != nil {
		log.Fatalf("%s bad size %q", err, chunks[0])
	}
	size = uint8(iSize)
	var m maze
	m.cells = make(map[coord]cell, size*size*size)
	m.depthMarker = coord{size, size, size}

	for p, r := range chunks[1] {
		if r == '*' {
			continue
		}
		here := iToCoord(p)
		m.cells[here] = cell(r == 'o')
		if here.isEdge() {
			if here.z == 0 {
				m.entrance = here
			} else {
				m.exit = here
			}
		}
	}
	return m
}

// solve returns length of shortest path.
func (m *maze) solve() int {
	// BFS
	// fmt.Println(*m)
	frontier := newQueue()
	visited := make(map[coord]bool)
	depth := 0
	frontier.push(m.entrance)
	frontier.push(m.depthMarker)
	for !frontier.empty() {
		c := frontier.pop()
		if c == m.depthMarker {
			if frontier.empty() {
				return 0
			}
			depth++
			frontier.push(m.depthMarker)
			continue
		}
		for _, n := range c.neighbours(m.cells) {
			if n == m.exit {
				return depth + 2
			}
			if !visited[n] {
				visited[n] = true
				frontier.push(n)
			}
		}
	}
	return 0
}

type queue struct {
	q          []coord
	head, tail int
}

func (q queue) String() string {
	return fmt.Sprintf("%v###%v", q.q[:q.head], q.q[q.head:q.tail])
}
func newQueue() queue {
	s := uint32(size)
	return queue{make([]coord, s*s*s), 0, 0}
}

func (q *queue) push(c coord) {
	q.q[q.tail] = c
	q.tail++
	if q.tail == len(q.q) {
		log.Fatal("queue full", len(q.q), *q)
	}
}
func (q *queue) pop() coord {
	if q.empty() {
		log.Fatal("empty pop")
	}
	c := q.q[q.head]
	q.head++
	return c
}

func (q *queue) empty() bool {
	return q.tail == q.head
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
		m := parse(scanner.Text())
		fmt.Println(m.solve())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
