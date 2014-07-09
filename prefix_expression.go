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
		fields := strings.Fields(line)
		reverse(fields)
		fmt.Println(postfix(fields))
	}
}

func postfix(fields []string) int {
	s := stack(make([]int, 0))

	for _, f := range fields {
		var v int
		switch f {
		case "*":
			v = s.pop() * s.pop()
		case "+":
			v = s.pop() + s.pop()
		case "/":
			a := s.pop()
			b := s.pop()
			v = a / b
		default:
			v = atoi(f)
		}
		s.push(v)
	}
	return s.pop()
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func reverse(a []string) {
	for j, k := 0, len(a)-1; j < k; j, k = j+1, k-1 {
		a[j], a[k] = a[k], a[j]
	}
}

type stack []int

func (s *stack) push(n int) {
	*s = append(*s, n)
}
func (s *stack) pop() int {
	pos := len(*s) - 1
	n := (*s)[pos]
	*s = (*s)[:pos]
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
