package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func count(m []string) int {
	reply := 0
	for r := 0; r < len(m)-1; r++ {
		for c := 0; c < len(m[0])-1; c++ {
			if code(m[r][c], m[r+1][c], m[r][c+1], m[r+1][c+1]) {
				reply++
			}
		}
	}
	return reply
}

func code(cs ...byte) bool {
	a := map[byte]bool{'c': true, 'o': true, 'd': true, 'e': true}
	for _, c := range cs {
		delete(a, c)
	}
	return len(a) == 0
}

func parse(line string) []string {
	return strings.Split(line, " | ")
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
		fmt.Println(count(parse(scanner.Text())))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
