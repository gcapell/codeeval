package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	size, totalCells int
	neighbourSpace   [8]square
	neighbourDelta   = []struct{ dx, dy int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1} /* {0,0} */, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	boardA, boardB board
)

type square struct{ x, y int }

func (s square) neighbours() []square {
	reply := neighbourSpace[:0]
	for _, d := range neighbourDelta {
		p := square{s.x + d.dx, s.y + d.dy}
		if p.valid() {
			reply = append(reply, p)
		}
	}
	return reply
}

func (s square) valid() bool {
	return s.x >= 0 && s.x < size && s.y >= 0 && s.y < size
}

type board struct {
	other *board
	live  map[square]bool
}

var rows int

func addRow(line string) {
	for x, r := range line {
		if r == '*' {
			boardA.live[square{x, rows}] = true
		}
	}
	rows++
}

func run(iterations int) *board {
	a, b := &boardA, &boardB
	for j := 0; j < iterations; j++ {
		b.step()
		a, b = b, a
	}
	return a
}

func (b *board) step() {
	b.live = make(map[square]bool, totalCells)
	dead := make(map[square]int, totalCells)
	for cell, _ := range b.other.live {
		live := 0
		for _, neighbour := range cell.neighbours() {
			if b.other.live[neighbour] {
				live++
			} else {
				dead[neighbour]++
			}
		}
		if live == 2 || live == 3 {
			b.live[cell] = true
		}
	}
	for cell, count := range dead {
		if count == 3 {
			b.live[cell] = true
		}
	}
}

var glyph = map[bool]rune{
	false: '.',
	true:  '*',
}

func (b *board) out() {
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			fmt.Printf("%c", glyph[b.live[square{col, row}]])
		}
		fmt.Println()
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
	boardA.other = &boardB
	boardB.other = &boardA
	boardA.live = make(map[square]bool)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		addRow(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	size = rows
	fmt.Println("size", size)
	totalCells = size * size
	run(10).out()
}
