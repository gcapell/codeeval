package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var triangle [][]int

	for line := range linesFromFilename() {
		triangle = append(triangle, atoi(line))
	}

	fmt.Println(max_total(triangle))
}

func max_total(triangle [][]int) int {
	for row := len(triangle) - 2; row >= 0; row-- {
		up, down := triangle[row], triangle[row+1]
		for pos := range up {
			up[pos] += max(down[pos], down[pos+1])
		}
	}
	return triangle[0][0]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func atoi(s string) []int {
	f := strings.Fields(s)
	reply := make([]int, len(f))
	for pos, j := range f {
		var err error
		reply[pos], err = strconv.Atoi(j)
		if err != nil {
			log.Fatal(err)
		}
	}

	return reply
}

func linesFromFilename() chan string {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	chunks := strings.Split(string(data), "\n")
	c := make(chan string)

	go func() {
		for _, line := range chunks {
			line = strings.TrimSpace(line)
			if len(line) > 0 {
				c <- line
			}
		}
		close(c)
	}()
	return c
}
