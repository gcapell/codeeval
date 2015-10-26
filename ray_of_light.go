package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

const (
	size  = 10
	slash = '/'
	slosh = '\\'
)

type room [size][size]byte

func (r *room) parse(line string) {
	for pos, ch := range line {
		r[pos/size][pos%size] = byte(ch)
	}
}

func (r *room) String() string {
	var b bytes.Buffer
	for _, row := range r {
		b.Write(row[:])
		b.WriteByte('\n')
	}
	return b.String()
}

type point struct{ row, col int }

func (p point) String() string {
	return fmt.Sprintf("{%d,%d}", p.row, p.col)
}

type direction struct {
	name       string
	dRow, dCol int
	symbol     byte
}

var (
	ne = direction{"ne", -1, 1, slash}
	se = direction{"se", 1, 1, slosh}
	sw = direction{"sw", 1, -1, slash}
	nw = direction{"nw", -1, 1, slosh}
)

const (
	ttl = 3 // 19? 20?
)

type cursor struct {
	point
	d         direction
	travelled int
	r         *room
}

func die(args ...interface{}) {
	log.Print(args)
	os.Exit(1)
}

func (c *cursor) move() {
	c.travelled++
	c.row += c.d.dRow
	c.col += c.d.dCol
	switch c.r[c.row][c.col] {
	case ' ':
		c.r[c.row][c.col] = c.d.symbol
	case '*':
		die("* unimplemented")
	case '#':
		die("# unimplemented")
	case 'o':
		die("o unimplemented")
	case c.d.symbol, 'x':
	default:
		c.r[c.row][c.col] = 'x'

	}
}

func (c cursor) String() string {
	return fmt.Sprintf("%s,%s,%d\n%s", c.point, c.d.name, c.travelled, c.r)
}
func startDirection(p point, c byte) direction {
	switch {
	case p.col == 0:
		return ifSlash(c, ne, se)
	case p.col == size-1:
		return ifSlash(c, sw, nw)
	case p.row == 0:
		return ifSlash(c, sw, se)
	case p.row == size-1:
		return ifSlash(c, ne, nw)
	}
	log.Fatal(p, c)
	return ne
}

func ifSlash(c byte, a, b direction) direction {
	if c == slash {
		return a
	}
	return b
}

func (r *room) start() cursor {
	for j := 0; j < size; j++ {
		for _, p := range []point{{0, j}, {j, 0}, {j, size - 1}, {size - 1, j}} {
			if c := r[p.row][p.col]; c == slash || c == slosh {
				return cursor{point: p, d: startDirection(p, c), r: r}
			}
		}
	}
	log.Fatal(*r)
	return cursor{}
}

func (r *room) sim() {
	c := r.start()
	fmt.Println(c)
	for c.travelled < ttl {
		c.move()
		fmt.Println(c)
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

	var r room
	for scanner.Scan() {
		r.parse(scanner.Text())
		r.sim()
		fmt.Println(&r)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
