package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	for line := range linesFromFilename() {
		fmt.Println(spiral(parse(line)))
	}
}

func parse(line string) [][]string {
	f := strings.Split(line, ";")
	rows, cols, values := atoi(f[0]), atoi(f[1]), strings.Fields(f[2])

	m := make([][]string, rows)
	for r := 0; r < rows; r++ {
		m[r] = make([]string, cols)
		for c := 0; c < cols; c++ {
			m[r][c] = values[r*cols+c]
		}
	}
	return m
}

type direction struct{ dc, dr int }

var directions = []direction{
	{1, 0}, {0, 1}, {-1, 0}, {0, -1},
}

type position struct{ r, c int }

func spiral(m [][]string) string {
	rows, cols := len(m), len(m[0])
	visited := map[position]bool{
		position{0, cols}:        true,
		position{rows, cols - 1}: true,
		position{rows - 1, -1}:   true,
	}

	reply := make([]string, 0, rows*cols)

	r, c, direction := 0, 0, 0
	for {
		reply = append(reply, m[r][c])
		visited[position{r, c}] = true

		d := directions[direction]
		nextR, nextC := r+d.dr, c+d.dc
		if visited[position{nextR, nextC}] {
			direction = (direction + 1) % 4
			d := directions[direction]
			nextR, nextC = r+d.dr, c+d.dc
			if visited[position{nextR, nextC}] {
				break
			}
		}
		r, c = nextR, nextC

	}

	return strings.Join(reply, " ")
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
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
