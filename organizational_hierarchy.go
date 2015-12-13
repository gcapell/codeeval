package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type node struct {
	name      byte
	children  []*node
	hasParent bool
}

type byName []*node

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].name < a[j].name }

func (n node) pprint(b *bytes.Buffer) {
	b.WriteByte(n.name)
	if len(n.children) > 0 {
		sort.Sort(byName(n.children))
		b.Write([]byte(" ["))
		for i, c := range n.children {
			if i != 0 {
				b.Write([]byte(", "))
			}
			c.pprint(b)
		}
		b.WriteByte(']')
	}
}

func chart(line string) string {
	var buf bytes.Buffer
	parse(line).pprint(&buf)
	return buf.String()
}

func parse(line string) *node {
	nodes := make(map[byte]*node)

	for _, edge := range strings.Split(line, " | ") {
		src, dst := find(nodes, edge[0]), find(nodes, edge[1])
		src.children = append(src.children, dst)
		dst.hasParent = true
	}
	for _, n := range nodes {
		if !n.hasParent {
			return n
		}
	}
	panic(nodes)
}

func find(nodes map[byte]*node, b byte) *node {
	if n, ok := nodes[b]; ok {
		return n
	}
	n := &node{name: b}
	nodes[b] = n
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
		fmt.Println(chart(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
