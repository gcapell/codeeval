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
	for sequence := range sequencesFromFile() {
		fmt.Println(strings.Join(cycle(sequence), " "))
	}
}

func sequencesFromFile() chan []string {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan []string)

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
			c <- strings.Fields(line)
		}
		close(c)
	}()
	return c
}

func same(a, b []string) bool {
	if len(a) != len(b) {
		log.Fatal(len(a), len(b))
	}
	for j := range a {
		if a[j] != b[j] {
			return false
		}
	}
	return true
}

func cycle(seq []string) []string {
	seen := make(map[string][]int)
	for pos, val := range seq {
		for _, start := range seen[val] {
			length := pos - start
			if same(seq[start:pos], seq[pos:pos+length]) {
				return seq[start:pos]
			}
		}
		seen[val] = append(seen[val], pos)
	}
	return nil
}
