package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	for m := range matricesFromFilename() {
		fmt.Println(minPathSum(m))
	}
}

func minPathSum(m [][]int) int {
	size := len(m)
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if r == 0 && c == 0 {
				continue
			}
			alts := []int{}
			if r != 0 {
				alts = append(alts, m[r-1][c])
			}
			if c != 0 {
				alts = append(alts, m[r][c-1])
			}
			sort.Ints(alts) // eggshell, meet sledgehammer
			m[r][c] += alts[0]
		}
	}

	return m[size-1][size-1]
}

func matricesFromFilename() chan [][]int {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan [][]int)

	go func() {
		reader := bufio.NewReader(f)
		for {
			size := 0
			_, err := fmt.Fscanln(reader, &size)
			if err != nil {
				break
			}
			m := make([][]int, size)
			for j := 0; j < size; j++ {
				row := make([]int, size)
				m[j] = row
				var line string
				_, err := fmt.Fscanln(reader, &line)
				if err != nil {
					log.Fatal(err)
				}
				fields := strings.Split(line, ",")
				for pos, f := range fields {
					var err error
					row[pos], err = strconv.Atoi(f)
					if err != nil {
						log.Fatal(err)
					}
				}
			}

			c <- m
		}
		close(c)
	}()
	return c
}
