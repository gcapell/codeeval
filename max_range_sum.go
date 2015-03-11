package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func maxRangeSum(line string) int {
	chunks := strings.Split(line, ";")
	days := atoi(chunks[0])
	deltas := strings.Fields(chunks[1])
	totals := make([]int, 0, len(deltas))
	total := 0
	for _, d := range deltas {
		total += atoi(d)
		totals = append(totals, total)
	}

	maxDelta := 0
	for s, e := -1, days-1; e < len(totals); s, e = s+1, e+1 {
		delta := totals[e]
		if s >= 0 {
			delta -= totals[s]
		}
		if delta > maxDelta {
			maxDelta = delta
		}
	}
	return maxDelta
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(s, err)
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
		fmt.Println(maxRangeSum(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
