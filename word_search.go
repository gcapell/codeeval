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
	b := parseBoard()
	for line := range linesFromFilename() {
		if b.find(line) {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}

type Board struct {
	start map[uint8][]*node
	nodes []*node
}

type node struct {
	pos  uint8
	r    uint8
	next map[uint8][]*node
}

func parseBoard() *Board {
	letters := "ABCESFCSADEE"
	neighbours := [][]int{
		{1, 4},
		{0, 2, 5},
		{1, 3, 6},
		{2, 7},
		{0, 8, 5},
		{1, 9, 4, 6},
		{5, 7, 2, 10},
		{3, 6, 11},
		{4, 9},
		{8, 5, 10},
		{9, 6, 11},
		{10, 7},
	}

	b := &Board{
		make(map[uint8][]*node),
		make([]*node, len(letters)),
	}

	for pos, r := range letters {
		ch := uint8(r)
		n := &node{uint8(pos), ch, make(map[uint8][]*node)}
		b.nodes[pos] = n
		b.start[ch] = append(b.start[ch], n)
	}

	for pos, ns := range neighbours {
		for _, nIndex := range ns {
			n := b.nodes[nIndex]
			m := b.nodes[pos].next
			m[n.r] = append(m[n.r], n)
		}
	}

	return b
}

func (b *Board) find(s string) bool {
	for _, n := range b.start[s[0]] {
		if n.find(s[1:], 0) {
			return true
		}
	}
	return false
}

func (n *node) find(s string, visited uint32) bool {
	if len(s) == 0 {
		return true
	}
	visited |= 1 << n.pos
	for _, o := range n.next[s[0]] {
		if visited&(1<<o.pos) == 0 {
			if o.find(s[1:], visited) {
				return true
			}
		}
	}
	return false
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
