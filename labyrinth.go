package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type (
	coord struct{ x, y uint8 }
	cell  bool // is hole
	maze  struct {
		cells          map[coord]cell
		entrance, exit coord
	}
	queue struct {
		q          []coord
		head, tail int
	}
)

var height, width uint8

func (m maze) String() string {
	b := &bytes.Buffer{}
	for y := uint8(0); y < height; y++ {
		for x := uint8(0); x < width; x++ {
			offTrack, ok := m.cells[coord{x, y}]
			var r rune
			switch {
			case !ok:
				r = '*'
			case bool(offTrack):
				r = ' '
			default:
				r = '+'
			}
			b.WriteRune(r)
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (c coord) neighbours(cells map[coord]cell) []coord {
	var reply []coord
	for _, p := range []coord{{c.x - 1, c.y}, {c.x + 1, c.y}, {c.x, c.y - 1}, {c.x, c.y + 1}} {
		if cells[p] {
			reply = append(reply, p)
		}
	}
	return reply
}

func parse(txt []byte) maze {
	txt = bytes.TrimSpace(txt)
	lines := bytes.Split(txt, []byte("\n"))
	height, width = uint8(len(lines)), uint8(len(lines[0]))
	var m maze
	m.cells = make(map[coord]cell, height*width)

	for y, line := range lines {
		for x, r := range line {
			if r == '*' {
				continue
			}
			here := coord{uint8(x), uint8(y)}
			m.cells[here] = true
			switch y {
			case 0:
				m.entrance = here
			case int(height - 1):
				m.exit = here
			}
		}
	}
	return m
}

// solve returns length of shortest path.
func (m *maze) solve() bool {
	// BFS
	back := make(map[coord]coord)
	frontier := newQueue()
	frontier.push(m.entrance)
	back[m.entrance] = m.entrance
	for !frontier.empty() {
		c := frontier.pop()
		for _, n := range c.neighbours(m.cells) {
			if _, ok := back[n]; ok {
				continue
			}
			back[n] = c
			if n == m.exit {
				m.backtrace(back, n)
				return true
			}
			frontier.push(n)
		}
	}
	return false
}

func (m maze) backtrace(back map[coord]coord, c coord) {
	for {
		m.cells[c] = false
		if d := back[c]; d == c {
			break
		} else {
			c = d
		}
	}
}

func (q queue) String() string {
	return fmt.Sprintf("%v###%v", q.q[:q.head], q.q[q.head:q.tail])
}
func newQueue() queue {
	return queue{make([]coord, int(height)*int(width)), 0, 0}
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
	txt, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	m := parse(txt)
	if m.solve() {
		fmt.Println(m)
	} else {
		fmt.Println("no")
	}
}
