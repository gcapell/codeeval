package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	for line := range linesFromFilename() {
		fmt.Println(search(parse(line)))
	}
}

// A parsed pattern is a slice of literal sub-patterns,
// separated by '*' operators.  Leading and trailing *
// operators are trivially satisfied, therefore ignored.
type pattern []string

func parse(line string) (string, pattern) {
	byComma := strings.Split(line, ",")

	textToSearch, patternText := byComma[0], byComma[1]
	buf := ""
	patterns := make([]string, 0)
	for {
		pos, ok := searchLiteral(patternText, "*")
		if !ok {
			patterns = append(patterns, buf+patternText[:])
			break
		}
		if pos > 0 && patternText[pos-1] == '\\' {
			buf += patternText[:pos-1] + "*"
			patternText = patternText[pos+1:]
			continue
		}
		patterns = append(patterns, buf+patternText[:pos])
		patternText = patternText[pos+1:]
		buf = ""
	}

	return textToSearch, patterns
}

func search(text string, literals pattern) bool {
	for _, p := range literals {
		loc, ok := searchLiteral(text, p)
		if !ok {
			return false
		}
		text = text[loc+len(p):]
	}
	return true
}

func searchLiteral(text, p string) (int, bool) {
baseLoop:
	for base := 0; base <= len(text)-len(p); base++ {
		for j := 0; j < len(p); j++ {
			if text[base+j] != p[j] {
				continue baseLoop
			}
		}
		return base, true
	}
	return 0, false
}

func linesFromFilename() chan string {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan string)

	go func() {
		reader := bufio.NewReader(f)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					log.Fatal(err)
				}
				break
			}
			c <- strings.TrimSpace(line)
		}
		close(c)
	}()
	return c
}
