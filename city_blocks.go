package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func blocks(line string) int {
	chunks := strings.Fields(line)
	streets, avenues := parse(chunks[0]), parse(chunks[1])
	fmt.Println(streets, avenues)
	return 0
}

func parse(line string)[]int {
	ns := strings.Split(line[1:len(line)-1], ",")
	reply := make([]int,0, len(ns))
	for _, n := range ns {
		nn, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		reply = append(reply, nn)
	}
	return reply
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
		fmt.Println(blocks(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
