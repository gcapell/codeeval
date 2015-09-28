package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func stepwise(word string) string {
	var reply []string
	for pos, c := range word {
		reply = append(reply, strings.Repeat("*", pos)+string(c))
	}
	return strings.Join(reply, " ")
}

func longest(line string) string {
	longest := ""
	for _, f := range strings.Fields(line) {
		if len(f) > len(longest) {
			longest = f
		}
	}
	return longest
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
		fmt.Println(stepwise(longest(scanner.Text())))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
