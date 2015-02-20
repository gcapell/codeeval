package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)


func knight(line string) string {
	col, row := int(line[0]), int(line[1])
	var reply []string
	for _, dx := range[]int{-2,-1,1,2} {
		c := col + dx
		if c< 'a' || c > 'h' {
			continue
		}
		dy := 3+dx
		if dy >2 {
			dy = 3-dx
		}
		for _, r := range []int{row-dy, row+dy} {
			if r <'1' || r > '8' {
				continue
			}
			reply = append(reply, string(c) + string(r))
		}
	}
	return strings.Join(reply, " ")}

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
		fmt.Println(knight(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
