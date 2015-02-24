package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func doubleThink(line string) string {
	chunks := strings.Split(line, ";")
	if len(chunks) != 2 {
		log.Fatal(len(chunks), line)
	}
	orig, want := strings.Fields(chunks[0]), strings.Fields(chunks[1])

	o := 0
	var reply []string
wantLoop:
	for _, w := range want {
		for o < len(orig) {
			chunks := strings.SplitN(orig[o], w, 2)
			if len(chunks) == 1 {
				reply = append(reply, underscore(orig[o]))
				o++
				continue
			}
			reply = append(reply, underscore(chunks[0])+w+underscore(chunks[1]))
			o++
			continue wantLoop
		}
		return "I cannot fix history"
	}
	for _, w := range orig[o:] {
		reply = append(reply, underscore(w))
	}
	return strings.Join(reply, " ")
}

const maxWord = 512

var underscores = string(bytes.Repeat([]byte{'_'}, maxWord))

func underscore(s string) string {
	if len(s) > maxWord {
		log.Fatal("maxWord", maxWord, s)
	}
	return underscores[:len(s)]
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
		fmt.Println(doubleThink(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
