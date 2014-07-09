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
		if goodChain(line) {
			fmt.Println("GOOD")
		} else {
			fmt.Println("BAD")
		}
	}
}

func goodChain(line string) bool {
	edges := strings.Split(line, ";")

	nodeMap := make(map[string]string)
	for _, edge := range edges {
		nodes := strings.Split(edge, "-")

		src, dst := nodes[0], nodes[1]
		if _, ok := nodeMap[src]; ok {
			return false
		}
		nodeMap[src] = dst
	}

	visited := make(map[string]bool)
	n := "BEGIN"
	for n != "END" {
		var ok bool
		n, ok = nodeMap[n]
		if !ok {
			return false
		}
		if visited[n] {
			return false
		}
		visited[n] = true
	}

	return len(visited) == len(nodeMap)
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
			line = strings.TrimSpace(line)
			if len(line) > 0 {
				c <- line
			}
		}
		close(c)
	}()
	return c
}
