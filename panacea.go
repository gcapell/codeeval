package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func moreBin(line string) bool {
	chunks := strings.Split(line, "|")
	return sumBase(chunks[1], 2) >= sumBase(chunks[0], 16)
}

func sumBase(line string, base int) int64 {
	var total int64
	for _, f := range strings.Fields(line) {
		n, err := strconv.ParseInt(f, base, 64)
		if err != nil {
			log.Fatal(err)
		}
		total += n
	}
	return total
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
		reply := "False"
		if moreBin(scanner.Text()) {
			reply = "True"
		}
		fmt.Println(reply)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
