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
	for p := range pairsFromFilename() {
		fmt.Println(multiple(p.a, p.b))
	}
}

func multiple(a, b int) int {
	m := a & ^(b - 1)
	if m == a {
		return m
	}
	return m + b
}

type pair struct{ a, b int }

func pairsFromFilename() chan pair {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan pair)

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
			ns := strings.Split(strings.TrimSpace(line), ",")
			if len(ns) != 2 {
				log.Fatalf("expected two ints, got %#v", ns)
			}
			a, err := strconv.Atoi(ns[0])
			if err != nil {
				log.Fatal(err)
			}
			b, err := strconv.Atoi(ns[1])
			if err != nil {
				log.Fatal(err)
			}
			c <- pair{a, b}
		}
		close(c)
	}()
	return c
}
