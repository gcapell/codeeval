package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func highest(line string) string {
	var row [][]int
	for _, r := range strings.Split(line, "|") {
		row = append(row, toInt(r))
	}
	var reply []string
	for col := 0; col < len(row[0]); col++ {
		var max int
		for pos, r := range row {
			if pos == 0 || r[col] > max {
				max = r[col]
			}
		}
		reply = append(reply, strconv.Itoa(max))
	}
	return strings.Join(reply, " ")
}

func toInt(row string) []int {
	fields := strings.Fields(row)
	var reply []int
	for _, f := range fields {
		n, err := strconv.Atoi(f)
		if err != nil {
			log.Fatal(err)
		}
		reply = append(reply, n)
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
		fmt.Println(highest(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
