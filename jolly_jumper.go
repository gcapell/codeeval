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
		if jolly(line) {
			fmt.Println("Jolly")
		} else {
			fmt.Println("Not jolly")
		}
	}
}

func jolly(line string) bool {
	f := strings.Fields(line)
	count := atoi(f[0])

	jumps := make([]bool, count)

	prev := atoi(f[1])
	for _, num := range f[1:] {
		n := atoi(num)
		delta := abs(n - prev)
		if delta >= count {
			return false
		}
		if jumps[delta] {
			return false
		}
		jumps[delta] = true
		prev = n
	}
	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
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
