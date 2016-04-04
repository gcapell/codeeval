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

func huffman(line string) string {
	return report(treeOf(count(line)).codes(""))
}

func report(codes []code) string {
	sort.Sort(byRune(codes))

	var reply bytes.Buffer
	for pos, c := range codes {
		if pos != 0 {
			fmt.Fprint(&reply, " ")
		}
		fmt.Fprintf(&reply, "%s: %s;", string(c.r), c.code)
	}
	return reply.String()
}

type code struct {
	r    rune
	code string
}

type byRune []code

func (a byRune) Len() int           { return len(a) }
func (a byRune) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byRune) Less(i, j int) bool { return a[i].r < a[j].r }

func treeOf(leaf []*node) *node {
	var internals []*node
	for len(leaf)+len(internals) > 1 {
		a := dequeue(&leaf, &internals)
		b := dequeue(&leaf, &internals)
		internals = append(internals, &node{n: a.n + b.n, l: a, r: b})
	}
	return internals[0]
}

// Return node with smallest scores from head of a or b
func dequeue(a, b *[]*node) *node {
	s := smaller(a, b)
	n := (*s)[0]
	*s = (*s)[1:]
	return n
}

func smaller(a, b *[]*node) *[]*node {
	if len(*a) == 0 {
		return b
	}
	if len(*b) == 0 {
		return a
	}
	if (*a)[0].lessThan((*b)[0]) {
		return a
	}
	return b
}

type node struct {
	c    rune
	n    int
	l, r *node
}

func (n node) String() string {
	if n.l == nil {
		return fmt.Sprintf("%s:%d", string(n.c), n.n)
	}
	return fmt.Sprintf("internal:%d", n.n)
}

func (n *node) prettyPrint(indent int) {
	fmt.Print(strings.Repeat(" ", indent))
	if n.l == nil {
		fmt.Println(string(n.c), n.n)
		return
	}
	fmt.Println(n.n)
	n.l.prettyPrint(indent + 4)
	n.r.prettyPrint(indent + 4)
}

func (n node) codes(prefix string) []code {
	if n.l != nil {
		return append(n.l.codes(prefix+"0"), n.r.codes(prefix+"1")...)
	}
	return []code{code{n.c, prefix}}
}

func (n *node) lessThan(o *node) bool {
	// different priority?
	if n.n != o.n {
		return n.n < o.n
	}
	// node vs symbol?
	if (n.l == nil) != (o.l == nil) {
		return n.l != nil
	}
	// compare symbols by character
	if n.l == nil {
		return n.c < o.c
	}
	// compare nodes by left child
	return n.l.lessThan(o.l)
}

type byCount []*node

func (a byCount) Len() int           { return len(a) }
func (a byCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byCount) Less(i, j int) bool { return a[i].lessThan(a[j]) }

func count(line string) []*node {
	counter := make(map[rune]int)
	for _, r := range line {
		counter[r] += 1
	}
	reply := make([]*node, 0, len(counter))
	for r, n := range counter {
		reply = append(reply, &node{c: r, n: n})
	}
	sort.Sort(byCount(reply))
	return reply
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
		fmt.Println(huffman(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
