package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func glueShredded(line string) string {
	line = strings.Trim(strings.TrimSpace(line), "|")
	chunks := strings.Split(line, "|")

	edges := make(map[int][]int)
	for i := 0; i < len(chunks)-1; i++ {
		for j := i + 1; j < len(chunks); j++ {
			if link(chunks[i], chunks[j]) {
				edges[i] = append(edges[i], j)
			}
			if link(chunks[j], chunks[i]) {
				edges[j] = append(edges[j], i)
			}
		}
	}
	
	fmt.Printf("edges %#v\n", edges)
	return ""

	path := hamiltonian(edges, len(chunks))
	var buf bytes.Buffer

	for _, p := range path {
		buf.WriteByte(chunks[p][0])
	}
	buf.WriteString(chunks[path[len(path)-1]][1:])

	return buf.String()
}

func origin(edges map[int][]int) (int, bool) {
	isDst := make(map[int]bool)
	for _, dsts := range edges {
		for _, dst := range dsts {
			isDst[dst] = true
		}
	}

	var o, count int
	for src := range edges {
		if !isDst[src] {
			o = src
			count++
		}
	}
	if count > 1 {
		log.Fatal("multiple origins", edges)
	}
	return o, count == 1
}

func hamiltonian(edges map[int][]int, nodes int) []int {
	path := make([]int, 0, nodes)
	seen := make(map[int]bool)
	if o, ok := origin(edges); ok {
		return traverse(append(path, o), edges, seen)
	}

	for e := range edges {
		if p := traverse(append(path, e), edges, seen); p!= nil {
			return p
		}
	}
	return nil
}

func traverse(path []int, edges map[int][]int, visited map[int]bool) []int {
	if len(path) == cap(path) {
		return path
	}
	p := path[len(path)-1]
	visited[p] = true
	defer func() { visited[p] = false }()
	for _, e := range edges[p] {
		if visited[e] {
			continue
		}
		if pp := traverse(append(path, e), edges, visited); pp != nil {
			return pp
		}
	}
	return nil
}

func link(a, b string) bool {
	return a[1:] == b[:len(b)-1]
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
		fmt.Println(glueShredded(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
