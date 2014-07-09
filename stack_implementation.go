package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	for line := range linesFromFilename() {
		printReversedAlternate(line)
	}
}

type stack []string

func (s *stack) Push(elem string) {
	*s = append(*s, elem)
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Pop() string {
	tail := len(*s) - 1
	elem := (*s)[tail]
	*s = (*s)[:tail]
	return elem
}

func printReversedAlternate(line string) {

	var s stack

	for _, f := range strings.Fields(line) {
		s.Push(f)
	}

	first := true
	for {
		if s.IsEmpty() {
			break
		}
		n := s.Pop()
		if !first {
			fmt.Printf(" ")
		}
		first = false
		fmt.Print(n)
		if s.IsEmpty() {
			break
		}
		s.Pop()
	}
	fmt.Println()

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
