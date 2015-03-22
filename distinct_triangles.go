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

func distinct(s string) int {
	chunks := strings.Split(s, ";")
	edges := make(map[int][]int)
	edgeStrings := chunks[1]
	for _, edge := range strings.Split(edgeStrings, ",") {
		chunks := strings.Fields(edge)
		src, dst := atoi(chunks[0]), atoi(chunks[1])
		if src > dst {
			src, dst = dst, src
		}
		edges[src] = append(edges[src], dst)
	}

	var keys []int
	for k, v := range edges {
		keys = append(keys, k)
		sort.Ints(v)
	}
	sort.Ints(keys)
	total := 0
	for _, a := range keys {
		for pos, b := range edges[a] {
			total += common(edges[a][pos+1:], edges[b])
		}
	}

	return total
}

// common returns the number of common elements in two sorted slices
// FIXME - this could be optimised (e.g. binary search to find next pos)
func common(a, b []int) int {
	count := 0
	for len(a) > 0 && len(b) > 0 {
		switch {
		case a[0] == b[0]:
			count++
			a, b = a[1:], b[1:]
		case a[0] < b[0]:
			a = a[1:]
		case a[0] > b[0]:
			b = b[1:]
		}
	}
	return count
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
		fmt.Println(distinct(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
