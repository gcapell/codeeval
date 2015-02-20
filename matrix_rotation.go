package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var sizes = map[int]int {100:10, 81:9, 64:8, 49:7, 36:6, 25:5, 16:4, 9:3, 4:2, 1:1,}

func rotate(line string) string {
	letters := strings.Fields(line)
	size := sizes[len(letters)]
	
	var reply []string
	for c := 0; c<size; c++ {
		for r := size-1; r >= 0; r-- {
			reply = append(reply, letters[r*size+c])
		}
	}
	return strings.Join(reply," ")
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
		fmt.Println(rotate(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
