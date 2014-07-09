package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	for line := range linesFromFilename() {
		fmt.Println(nSumZero(line))
	}
}

type Quads map[string]bool

func New() Quads {
	return Quads(make(map[string]bool))
}

func (q Quads) add(as, bs [][]int) {
	for _, a := range as {
		for _, b := range bs {
			c := []int{}
			c = append(c, a...)
			c = append(c, b...)
			sort.Ints(c)
			if !dups(c) {
				q[fmt.Sprintf("%v", c)] = true
			}
		}
	}
}

func dups(a []int) bool {
	var prev int
	for pos, val := range a {
		if pos != 0 && val == prev {
			return true
		}
		prev = val
	}
	return false
}

func (q Quads) count() int { return len(q) }

func nSumZero(line string) int {
	ns := atoi(line)

	// sum -> slice of pair of index
	sum2 := make(map[int][][]int)

	for pos, n1 := range ns {
		for off, n2 := range ns[pos+1:] {
			sum2[n1+n2] = append(sum2[n1+n2], []int{pos, pos + 1 + off})
		}
	}

	sumKeys := make([]int, 0, len(sum2))
	for k, _ := range sum2 {
		sumKeys = append(sumKeys, k)
	}
	sort.Ints(sumKeys)
	low, high := 0, len(sumKeys)-1

	if sumKeys[high] < 0 {
		return 0
	}

	quads := New()
	for ; low < high; high-- {
		for sumKeys[low]+sumKeys[high] < 0 {
			low++
		}
		if sumKeys[low]+sumKeys[high] == 0 {
			quads.add(sum2[sumKeys[low]], sum2[sumKeys[high]])
		}
	}
	// fmt.Printf("%v\n", quads)
	return quads.count()
}

func atoi(s string) []int {
	fields := strings.Split(s, ",")
	ns := make([]int, len(fields))
	for pos, s := range fields {
		var err error
		ns[pos], err = strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
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
			c <- strings.TrimSpace(line)
		}
		close(c)
	}()
	return c
}
