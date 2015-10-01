package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func mapBools(rows []string) [][]bool {
	var reply [][]bool
	for _, row := range rows {
		reply = append(reply, mapBool(strings.TrimSpace(row)))
	}
	return reply
}

func mapBool(s string) []bool {
	reply := make([]bool, len(s))
	for n, r := range s {
		reply[n] = r == '1'
	}
	return reply
}

func blackOrWhite(line string) string {
	rows := mapBools(strings.Split(line, "|"))
	n, c := submatrix(rows)
	return fmt.Sprintf("%dx%d, %d", n, n, c)
}

func submatrix(m [][]bool) (int, int) {
	for n := 1; ; n++ {
		if c, ok := count(m, n); ok {
			return n, c
		}
	}
}

func count(m [][]bool, sub int) (int, bool) {
	size := len(m)
	first := true
	var count int
	for x := 0; x+sub <= size; x++ {
		for y := 0; y+sub <= size; y++ {
			n := 0
			for j := 0; j < sub; j++ {
				for k := 0; k < sub; k++ {
					if m[x+j][y+k] {
						n++
					}
				}
			}
			switch {
			case first:
				count = n
				first = false
			case n != count:
				return -1, false
			}
		}
	}
	return count, true
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
		fmt.Println(blackOrWhite(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
