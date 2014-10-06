package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func recover(line string) string {
	chunks := strings.Split(line, ";")
	words, hints := strings.Fields(chunks[0]), strings.Fields(chunks[1])
	reply := make([]string, len(words))
	missing := len(words) * (len(words) + 1) / 2
	for j := range hints {
		n, _ := strconv.Atoi(hints[j])
		reply[n-1] = words[j]
		missing -= n
	}
	reply[missing-1] = words[len(words)-1]
	return strings.Join(reply, " ")
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
		fmt.Println(recover(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
