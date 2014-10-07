package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var (
	profiling      = false
	size           int8
	neighbourSpace [8]square
	neighbourDelta = []struct{ dx, dy int8 }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1} /* {0,0} */, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	boardA, boardB board
)

type square struct{ x, y int8 }

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

var rows int8

func addRow(line string) {
	for x, r := range line {
		if r == '*' {
			boardA.live[square{int8(x), rows}] = true
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
	b.live = make(map[square]bool, 0)
	dead := make(map[square]int, 0)
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
	for row := int8(0); row < size; row++ {
		for col := int8(0); col < size; col++ {
			fmt.Printf("%c", glyph[b.live[square{col, row}]])
		}
		fmt.Println()
	}
}

func init() {
	boardA.other = &boardB
	boardB.other = &boardA
	boardA.live = make(map[square]bool, 0)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expected filename")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if profiling {
		cf, err := os.Create("cpu.out")
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(cf)
		defer pprof.StopCPUProfile()
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		addRow(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	size = rows
	run(10).out()

	if profiling {
		pf, err := os.Create("mem.out")
		if err != nil {
			log.Fatal(err)
		}
		defer pf.Close()
		pprof.WriteHeapProfile(pf)
	}
}
