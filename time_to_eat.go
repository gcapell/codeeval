package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func sortedTimes(line string) string {
	fields := strings.Fields(line)
	sort.Strings(fields)
	for l, r := 0, len(fields)-1; l < r; l, r = l+1, r-1 {
		fields[l], fields[r] = fields[r], fields[l]
	}
	return strings.Join(fields, " ")
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
		fmt.Println(sortedTimes(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
