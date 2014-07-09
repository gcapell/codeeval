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
		fields := strings.Split(line, ",")
		fmt.Println(subsequences(fields[0], fields[1]))
	}
}

func subsequences(seq, sub string) int {
	a := make([]int, len(seq))
	count := 0
	for pos, r := range seq {
		if r == rune(sub[0]) {
			count++
		}
		a[pos] = count
	}
	b := make([]int, len(a))

	for _, r := range sub[1:] {
		for pos, r2 := range seq {
			switch {
			case pos == 0:
				b[pos] = 0
			case r == r2:
				b[pos] = b[pos-1] + a[pos-1]
			default:
				b[pos] = b[pos-1]
			}
		}
		a, b = b, a
	}
	return a[len(a)-1]
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
