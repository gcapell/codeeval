package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var directions = map[string]struct {
	origin, rowFirst bool
}{
	"UP":    {true, false},
	"DOWN":  {false, false},
	"LEFT":  {true, true},
	"RIGHT": {false, true},
}

func puzzle(line string) {
	parts := strings.Split(line, ";")
	direction, size, board := parts[0], atoi(parts[1]), parseBoard(parts[2])

	d := directions[direction]
	start, delta := 0, 1
	if !d.origin {
		start, delta = size-1, -1
	}

	ns := make([]int, size)
	for x, j := start, 0; j < size; x, j = x+delta, j+1 {
		for y, k := start, 0; k < size; y, k = y+delta, k+1 {
			if d.rowFirst {
				ns[k] = board[x][y]
			} else {
				ns[k] = board[y][x]
			}
		}
		// fmt.Printf("%v -> ", ns)
		compress(ns)
		// fmt.Println(ns)
		for y, k := start, 0; k < size; y, k = y+delta, k+1 {
			if d.rowFirst {
				board[x][y] = ns[k]
			} else {
				board[y][x] = ns[k]
			}
		}
	}
	out(board)
}

func parseBoard(line string) [][]int {
	var reply [][]int
	for _, row := range strings.Split(line, "|") {
		var r []int
		for _, ns := range strings.Fields(row) {
			r = append(r, atoi(ns))
		}
		reply = append(reply, r)
	}
	return reply
}

func out(board [][]int) {
	for m, row := range board {
		if m != 0 {
			fmt.Printf("|")
		}
		for n, col := range row {
			if n != 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", col)
		}
	}
	fmt.Println()
}

func compress(n []int) {
	canDouble := false
	dst := 0
	for src := 0; src < len(n); src++ {
		if n[src] == 0 {
			continue
		}
		if canDouble {
			if n[dst] == n[src] {
				n[dst] *= 2
				dst++
				canDouble = false
				continue
			}
			dst++
		}
		n[dst] = n[src]
		canDouble = true
	}
	if canDouble {
		dst++
	}
	for j := dst; j < len(n); j++ {
		n[j] = 0
	}
}

func atoi(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Fatal(err)
	}
	return n
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
		puzzle(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
