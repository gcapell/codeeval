package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func minNew(line string) int {
	chunks := strings.Split(line, " | ")
	c, v, current := atoi(chunks[0]), atoi(chunks[1]), mapAtoi(strings.Fields(chunks[2]))

	maxPossible, added := 0, 0
	for maxPossible < v {
		if len(current) > 0 && current[0] <= maxPossible+1 {
			maxPossible += current[0] * c
			current = current[1:]
		} else {
			maxPossible += (maxPossible + 1) * c
			added++
		}
	}
	return added
}

func atoi(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func mapAtoi(ss []string) []int {
	reply := make([]int, 0, len(ss))
	for _, s := range ss {
		reply = append(reply, atoi(s))
	}
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
		fmt.Println(minNew(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
