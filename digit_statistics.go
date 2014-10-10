package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func dumb(a, n int) map[int]int {
	digit := 1
	reply := make(map[int]int)

	for j := 0; j < n; j++ {
		digit = (digit * a) % 10
		reply[digit]++
	}
	return reply
}

func stats(a, n int) map[int]int {
	digit := 1
	var sequence []int
	seen := make(map[int]bool)

	for j := 0; j < n; j++ {
		digit = (digit * a) % 10
		if seen[digit] {
			return extend(sequence, digit, n)
		}
		sequence = append(sequence, digit)
		seen[digit] = true
	}
	// No dups
	reply := make(map[int]int)
	for _, digit := range sequence {
		reply[digit] = 1
	}
	return reply
}

func extend(sequence []int, digit, n int) map[int]int {
	reply := make(map[int]int)
	for pos, d := range sequence {
		reply[d] = n / len(sequence)
		if n%len(sequence) > pos {
			reply[d]++
		}
	}
	return reply
}

func pretty(counts map[int]int) string {
	var chunks []string
	for j := 0; j < 10; j++ {
		chunks = append(chunks, fmt.Sprintf("%d: %d", j, counts[j]))
	}
	return strings.Join(chunks, ", ")
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
		var a, n int
		fmt.Sscanf(scanner.Text(), "%d %d", &a, &n)
		fmt.Println(pretty(stats(a, n)))

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
