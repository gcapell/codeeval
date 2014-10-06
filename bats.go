package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func bats(line string) int {
	numbers := mapInt(line)
	length, distance, nBats, positions := numbers[0], numbers[1], numbers[2], numbers[3:]
	if nBats != len(positions) {
		log.Fatalf("%d bats but %v in %q", nBats, positions, line)
	}
	if nBats == 0 {
		return (length-12)/distance + 1
	}
	sort.Ints(positions)
	total := (positions[0] - 6) / distance                         // left
	total += (length - positions[len(positions)-1] - 6) / distance // right
	for j := 1; j < len(positions); j++ {
		divs := (positions[j] - positions[j-1]) / distance
		if divs > 1 {
			total += divs - 1
		}
	}
	return total
}

func mapInt(line string) []int {
	fields := strings.Fields(line)
	reply := make([]int, 0, len(fields))
	for _, f := range fields {

		n, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}

		reply = append(reply, n)
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
		fmt.Println(bats(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
