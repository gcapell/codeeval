package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const letters = 26

func columnName(line string) string {
	n, err := strconv.Atoi(line)
	if err != nil {
		log.Fatal(err)
	}
	n--
	decade := letters
	digits := 1
	for n >= decade {
		n -= decade
		decade *= letters
		digits++
	}

	reply := make([]byte, digits)
	for j := 0; j < digits; j++ {
		reply[digits-j-1] = byte('A' + n%letters)
		n /= letters
	}
	return string(reply)
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
		fmt.Println(columnName(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
