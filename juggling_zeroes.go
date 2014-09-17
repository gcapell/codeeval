package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func zeroes(line string) int {
	words := strings.Fields(line)
	n := 0
	for len(words) > 0 {
		var flag, seq string
		flag, seq, words = words[0], words[1], words[2:]
		seqLen := uint(len(seq))
		n <<= seqLen
		if len(flag) == 2 {
			n += (1 << seqLen) - 1
		}
	}
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
		fmt.Println(zeroes(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
