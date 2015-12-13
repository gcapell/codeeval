package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func errors(line string) int {
	parts := strings.Split(line, " | ")
	left, right := parts[0], parts[1]
	errors := 0
	for pos := 0; pos < len(left); pos++ {
		if right[pos] != left[pos] {
			errors++
		}
	}
	return errors
}

var statusCounts = []struct {
	n   int
	msg string
}{
	{0, "Done"},
	{2, "Low"},
	{4, "Medium"},
	{6, "High"},
	{50, "Critical"}}

func status(n int) string {
	for _, s := range statusCounts {
		if n <= s.n {
			return s.msg
		}
	}
	return "ERROR"
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
		fmt.Println(status(errors(scanner.Text())))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
