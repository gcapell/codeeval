package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var (
	table = []string{
		", yeah!",
		", this is crazy, I tell ya.",
		", can U believe this?",
		", eh?",
		", aw yea.",
		", yo.",
		"? No way!",
		". Awesome!",
	}
	pos     int
	replace bool
	re      = regexp.MustCompile("[.?!]")
)

func slang(line string) string {
	return re.ReplaceAllStringFunc(line, func(s string) string {
		if !replace {
			replace = true
			return s
		}
		reply := table[pos]
		pos = (pos + 1) % len(table)
		replace = false
		return reply
	})
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
		fmt.Println(slang(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
