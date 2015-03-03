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
type pair struct {a,b int}

func maybeLink(a,b int, chunks []string, forward, backward edgelist) {
	if link(chunks[a], chunks[b]) {
		forward[a] = append(forward[a], b)
		backward[b] = append(backward[b], a)
	}
}

func reduce(forward, backward edgelist) (map[int]bool, map[pair]int) {
	skipped := make(map[int]bool)
	replacements := make(map[pair]int)
	for k := range forward {
		if skipped[k] {
			continue
		}
		if ! (len(forward[k]) == 1 && len(backward[k]) ==1) {
			continue
		}
		s := k
		next := k
		for len(forward[s]) == 1 && len(backward[s]) ==1 {
			skipped[s]=true
			next =s
			s = backward[s][0]
		}
		e := forward[k][0]
		for len(forward[e]) == 1 && len(backward[e]) ==1 {
			skipped[e] = true
			e = forward[e][0]
		}
		replaceLink(forward[s], next, e)
		replacements[pair{s,e}] = next
	}
	return skipped, replacements
}

func replaceLink(a []int, orig, replacement int) {
	for k,v := range a{
		if v == orig {
			a[k] = replacement
			return
		}
	}
	log.Fatal("couldn't find %v in %v", orig, a)
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
	
	skipped, _ := reduce(forward, backward)
	// fmt.Printf("skipped %v\n", skipped)
	// fmt.Printf("expansions %v\n", expansions)

	fmt.Printf("digraph T {")
	for k,vs := range forward {
		if skipped[k] {
			continue
		}
		for _, v := range vs {
			fmt.Printf("%d->%d;", k, v)
		}
	}
	fmt.Println("}")

	return ""
	path := hamiltonian(forward, len(chunks))
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
