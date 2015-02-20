package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func guess(line string) int {
	fields := strings.Fields(line)
	high, _ := strconv.Atoi(fields[0])
	low := 0

	for _, m := range fields[1:] {
		g := (low + high + 1) / 2
		switch m {
		case "Lower":
			high = g - 1
		case "Higher":
			low = g + 1
		case "Yay!":
			return g
		}
	}
	log.Fatal("no yay?")
	return -1
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
		fmt.Println(guess(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
