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
		if overlapping(line) {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}

func overlapping(line string) bool {
	n := atoi(strings.Split(line, ","))
	return overlap(n[0], n[2], n[4], n[6]) && overlap(n[7], n[5], n[3], n[1])
}

func overlap(a0, a1, b0, b1 int) bool {
	return (a0 <= b0 && b0 <= a1) || (b0 <= a0 && a0 <= b1)
}

func atoi(ss []string) []int {
	ns := make([]int, len(ss))

	for pos, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		ns[pos] = n
	}
	return ns
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
			line = strings.TrimSpace(line)
			if len(line) > 0 {
				c <- line
			}
		}
		close(c)
	}()
	return c
}
