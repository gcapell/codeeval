package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// http://stackoverflow.com/questions/10177389/generalised-two-egg-puzzle
// The highest ﬂoor that can be tested with
// d drops and e eggs is as follows :
// f[d,e] = f[d-1,e] + f[d-1,e-1] + 1

type pair struct{ a, b int }

var cache = make(map[pair]int)

func eggDrop(eggs, floors int) int {
	for d := 1; ; d++ {
		if maxDrop(d, eggs) >= floors {
			return d
		}
	}
}

func maxDrop(d, e int) int {
	if e == 1 {
		return d
	}
	if d == 0 {
		return 0
	}
	p := pair{d, e}
	if f, ok := cache[p]; ok {
		return f
	}
	f := maxDrop(d-1, e) + maxDrop(d-1, e-1) + 1
	cache[p] = f
	return f
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expected filename")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fs := strings.Fields(scanner.Text())
		fmt.Println(eggDrop(atoi(fs[0]), atoi(fs[1])))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
