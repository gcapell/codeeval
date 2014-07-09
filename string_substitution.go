package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	for line := range linesFromFilename() {
		fmt.Println(replace(parse(line)))
	}
}

func replace(orig string, patterns []pattern) string {
	root := &node{orig, false, nil}
	for _, p := range patterns {
		root.findAndReplace(p)
	}
	return root.flat()
}

type node struct {
	s        string
	replaced bool
	next     *node
}

func (n *node) String() string {
	if n == nil {
		return ""
	}
	var left, right string
	if n.replaced {
		left, right = "<", ">"

	} else {
		left, right = `"`, `"`
	}
	return left + n.s + right + n.next.String()
}

func (n *node) findAndReplace(p pattern) {
	// fmt.Printf("findAndReplace(%q,%q,%s)\n", p.pattern, p.replacement, n)
	for ; n != nil; n = n.next {
		if !n.replaced {
			if pos := strings.Index(n.s, p.pattern); pos != -1 {
				n.replace(pos, len(p.pattern), p.replacement)
			}
		}
	}
}

func (n *node) replace(pos, length int, replacement string) {
	// fmt.Printf("replace %q (%d:%d) with %q\n", n.s, pos, length, replacement)
	// replace start
	if pos == 0 {
		if length == len(n.s) {
			// replace whole
			n.s = replacement
			n.replaced = true
			return
		}
		next := &node{n.s[length:], false, n.next}
		n.s = replacement
		n.replaced = true
		n.next = next
		return
	}

	// replace end
	if pos+length == len(n.s) {
		next := &node{replacement, true, n.next}
		n.s = n.s[:pos]
		n.next = next
		return
	}

	// replace middle
	end := &node{n.s[pos+length:], false, n.next}
	middle := &node{replacement, true, end}
	n.s = n.s[:pos]
	n.next = middle
}

func (n *node) flat() string {
	if n == nil {
		return ""
	}
	return n.s + n.next.flat()
}

type pattern struct {
	pattern, replacement string
}

func parse(line string) (string, []pattern) {
	f := strings.Split(line, ";")
	orig := f[0]
	f = strings.Split(f[1], ",")
	patterns := make([]pattern, len(f)/2)

	for j := 0; j < len(f); j += 2 {
		patterns[j/2] = pattern{f[j], f[j+1]}
	}
	return orig, patterns
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
