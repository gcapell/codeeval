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
		words := strings.Fields(line)
		if len(words) == 0 {
			continue
		}
		reverse(words)
		fmt.Println(strings.Join(words, " "))
	}
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
			c <- line
		}
		close(c)
	}()
	return c
}

func reverse(s []string) {
	for a, b := 0, len(s)-1; a < b; a, b = a+1, b-1 {
		s[a], s[b] = s[b], s[a]
	}
}
