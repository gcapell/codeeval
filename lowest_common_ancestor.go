package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Val         int
	Left, Right *Node
}

var tree = &Node{30,
	&Node{8,
		&Node{3, nil, nil},
		&Node{20,
			&Node{10, nil, nil},
			&Node{29, nil, nil}},
	},
	&Node{52, nil, nil},
}

func main() {
	for line := range linesFromFilename() {
		a, b := parse(line)
		fmt.Println(tree.lca(a, b))
	}
}

func (t *Node) lca(a, b int) int {
	if a < t.Val && b < t.Val {
		return t.Left.lca(a, b)
	}
	if a > t.Val && b > t.Val {
		return t.Right.lca(a, b)
	}
	return t.Val
}

func parse(s string) (int, int) {
	f := strings.Fields(s)
	if len(f) != 2 {
		log.Fatalf("expected 2 fields, got %q", f)
	}

	a, err := strconv.Atoi(f[0])
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.Atoi(f[1])
	if err != nil {
		log.Fatal(err)
	}
	return a, b
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
