package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func blackSpot(line string) string {
	parts := strings.Split(line, " | ")
	n, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	names := strings.Fields(parts[0])
	pos := 0
	for len(names) > 1 {
		spot := (pos + n - 1) % len(names)
		names = append(names[:spot], names[spot+1:]...)
	}
	return names[0]
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
		fmt.Println(blackSpot(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
