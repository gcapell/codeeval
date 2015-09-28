package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func mask(line string) string {
	fields := strings.Fields(line)
	word, mask := []byte(fields[0]), fields[1]
	for pos, c := range word {
		if mask[pos] == '1' {
			word[pos] = byte(unicode.ToUpper(rune(c)))
		}
	}
	return string(word)
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
		fmt.Println(mask(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
