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
		b := parse(line)
		if b.valid() {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}

type board struct {
	n, side int
	b       [][]int
}

type sequenceT struct {
	name string
	off  int
	s    chan int
}

func parse(line string) board {
	f1 := strings.Split(line, ";")
	n := atoi(f1[0])
	f2 := strings.Split(f1[1], ",")
	b := make([][]int, n)
	for j := 0; j < n; j++ {
		b[j] = make([]int, n)
		for k := 0; k < n; k++ {
			b[j][k] = atoi(f2[j*n+k])
		}
	}
	if n == 4 {
		return board{n, 2, b}
	}
	return board{n, 3, b}
}

func (b board) print() {
	for j := 0; j < b.n; j++ {
		for k := 0; k < b.n; k++ {
			fmt.Print(b.b[j][k])
		}
		fmt.Println()
	}
}

func (b board) valid() bool {
	for s := range b.sequences() {
		if !uniques(s.s, b.n) {
			return false
		}
	}
	return true
}

func uniques(ch chan int, max int) bool {
	seen := make(map[int]bool)
	for n := range ch {
		if n < 0 || n > max {
			return false
		}
		if seen[n] {
			return false
		}
		seen[n] = true
	}
	return true
}

func (b board) sequences() chan sequenceT {
	ch := make(chan sequenceT)
	go func() {
		for j := 0; j < b.n; j++ {
			ch <- sequenceT{"row", j, b.row(j)}
			ch <- sequenceT{"col", j, b.col(j)}
			ch <- sequenceT{"square", j, b.square(j)}
		}
		close(ch)
	}()
	return ch
}

func (b board) row(r int) chan int {
	ch := make(chan int)
	go func() {
		for j := 0; j < b.n; j++ {
			ch <- b.b[r][j]
		}
		close(ch)
	}()
	return ch
}

func (b board) col(c int) chan int {
	ch := make(chan int)
	go func() {
		for j := 0; j < b.n; j++ {
			ch <- b.b[j][c]
		}
		close(ch)
	}()
	return ch
}

func (b board) square(s int) chan int {
	rowOff := s / b.side
	colOff := s % b.side

	ch := make(chan int)
	go func() {
		for j := 0; j < b.side; j++ {
			for k := 0; k < b.side; k++ {
				ch <- b.b[rowOff*b.side+j][colOff*b.side+k]
			}
		}
		close(ch)
	}()
	return ch
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
