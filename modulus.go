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
	for p := range fileToProblems() {
		fmt.Println(modulus(p.a, p.b))
	}
}

func modulus(a, b int) int {
	return a - b*(a/b)
}

type problem struct{ a, b int }

func newProblem(s string) (problem, error) {
	sides := strings.Split(s, ",")
	var p problem
	if len(sides) != 2 {
		return p, fmt.Errorf("expected two ,-delimited strings, got %q", s)
	}

	p = problem{atoi(sides[0]), atoi(sides[1])}
	return p, nil
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func fileToProblems() chan problem {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan problem)

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
			p, err := newProblem(strings.TrimSpace(line))
			if err != nil {
				log.Fatal(err)
			}
			c <- p
		}
		close(c)
	}()
	return c
}
