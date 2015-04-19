package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func comparePoints(line string) string {
	fields := strings.Fields(line)
	x0, y0, x1, y1 := atoi(fields[0]), atoi(fields[1]), atoi(fields[2]), atoi(fields[3])

	lat, lon := "", ""
	switch {
	case y0 < y1:
		lat = "N"
	case y0 > y1:
		lat = "S"
	}
	switch {
	case x0 < x1:
		lon = "E"
	case x0 > x1:
		lon = "W"
	}
	dir := lat + lon
	if dir == "" {
		dir = "here"
	}
	return dir
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
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
		fmt.Println(comparePoints(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
