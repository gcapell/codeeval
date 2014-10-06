package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const font = `
-**----*--***--***---*---****--**--****--**---**--
*--*--**-----*----*-*--*-*----*-------*-*--*-*--*-
*--*---*---**---**--****-***--***----*---**---***-
*--*---*--*-------*----*----*-*--*--*---*--*----*-
-**---***-****-***-----*-***---**---*----**---**--
--------------------------------------------------
`

const cols, rows = 5, 6

var fontLine = strings.Fields(font)

func big(line string) string {
	digits := strings.Map(func(r rune) rune {
		if r < '0' || r > '9' {
			return -1
		}
		return r
	}, line)
	replyChunks := make([][]string, rows)
	for _, r := range digits {
		pos := (r - '0') * cols
		for j := 0; j < rows; j++ {
			replyChunks[j] = append(replyChunks[j], fontLine[j][pos:pos+cols])
		}
	}
	replyLines := make([]string, rows)
	for row, chunks := range replyChunks {
		replyLines[row] = strings.Join(chunks, "")
	}
	return strings.Join(replyLines, "\n")
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
		fmt.Println(big(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
