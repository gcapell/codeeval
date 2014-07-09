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
	for p := range fileToProblems() {
		fmt.Println(strings.LastIndex(p.haystack, p.needle))
	}
}

type problem struct{ haystack, needle string }

func newProblem(s string) (problem, error) {
	sides := strings.Split(s, ",")
	var p problem
	if len(sides) != 2 {
		return p, fmt.Errorf("expected two ,-delimited strings, got %q", s)
	}
	if len(sides[1]) != 1 {
		return p, fmt.Errorf("expected single char, got %q", sides[1])
	}
	p = problem{sides[0], sides[1]}
	return p, nil
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
