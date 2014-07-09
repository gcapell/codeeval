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
		fmt.Println(pangram(line))
	}
}

func pangram(line string) string {
	seen := make(map[rune]bool)
	line = strings.ToLower(line)
	for _, c := range line {
		seen[c] = true
	}
	reply := make([]byte, 0)
	for _, c := range "abcdefghijklmnopqrstuvwxyz" {
		if !seen[c] {
			reply = append(reply, byte(c))
		}
	}
	if len(reply) > 0 {
		return string(reply)
	}
	return "NULL"
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
