package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

type edgelist map[int][]int
type pair struct{ a, b int }

func maybeLink(a, b int, chunks []string, forward, backward edgelist) {
	if link(chunks[a], chunks[b]) {
		forward[a] = append(forward[a], b)
		backward[b] = append(backward[b], a)
	}
}

func reduce(forward, backward edgelist) (edgelist, map[pair]int) {
	skipped := make(map[int]bool)
	longcuts := make(map[pair]int)
	for k := range forward {
		if skipped[k] {
			continue
		}
		if !(len(forward[k]) == 1 && len(backward[k]) == 1) {
			continue
		}
		s := k
		next := k
		for len(forward[s]) == 1 && len(backward[s]) == 1 {
			skipped[s] = true
			next = s
			s = backward[s][0]
		}
		e := forward[k][0]
		for len(forward[e]) == 1 && len(backward[e]) == 1 {
			skipped[e] = true
			e = forward[e][0]
		}
		replaceLink(forward[s], next, e)
		longcuts[pair{s, e}] = next
	}
	replaced := make(map[int][]int)
	for k, v := range forward {
		if skipped[k] {
			continue
		}
		replaced[k] = v
	}
	return replaced, longcuts
}

func replaceLink(a []int, orig, replacement int) {
	for k, v := range a {
		if v == orig {
			a[k] = replacement
			return
		}
	}
	log.Fatal("couldn't find %v in %v", orig, a)
}

func digraph(forward edgelist) {
	fmt.Printf("digraph T {")
	for k, vs := range forward {
		for _, v := range vs {
			fmt.Printf("%d->%d;", k, v)
		}
	}
	fmt.Println("}")
}

func expandPath(shortPath []int, edges edgelist, longcuts map[pair]int) []int {
	var path []int
	prev := -1
	for _, n := range shortPath {
		if n2, ok := longcuts[pair{prev, n}]; ok {
			for n2 != n {
				path = append(path, n2)
				if len(edges[n2]) != 1 {
					log.Fatal("bad longcut", n2, n)
				}
				n2 = edges[n2][0]
			}
		}
		path = append(path, n)
		prev = n
	}
	return path
}

func glueShredded(line string) string {
	line = strings.Trim(strings.TrimSpace(line), "|")
	chunks := strings.Split(line, "|")

	forward, backward := make(map[int][]int), make(map[int][]int)
	for i := 0; i < len(chunks)-1; i++ {
		for j := i + 1; j < len(chunks); j++ {
			maybeLink(i, j, chunks, forward, backward)
			maybeLink(j, i, chunks, forward, backward)
		}
	}

	replaced, longcuts := reduce(forward, backward)

	shortPath := hamiltonian(replaced)

	path := expandPath(shortPath, forward, longcuts)

	var buf bytes.Buffer

	for _, p := range path {
		buf.WriteByte(chunks[p][0])
	}
	buf.WriteString(chunks[path[len(path)-1]][1:])

	return buf.String()
}

func origin(edges map[int][]int) (origin, nodes int) {
	isDst := make(map[int]bool)
	for src, dsts := range edges {
		if _, ok := isDst[src]; !ok {
			isDst[src] = false
		}
		for _, dst := range dsts {
			isDst[dst] = true
		}
	}

	count := 0
	for src := range edges {
		if !isDst[src] {
			origin = src
			count++
		}
	}
	if count != 1 {
		log.Fatal("origins", count, edges)
	}
	return origin, len(isDst)
}

func hamiltonian(edges map[int][]int) []int {
	o, nodes := origin(edges)
	path := make([]int, 0, nodes)
	seen := make(map[int]bool)
	return traverse(append(path, o), edges, seen)
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
