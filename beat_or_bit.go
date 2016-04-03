package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var gray = make(map[string]int)

func init() {
	code := []string{"0", "1"}
	for {
		for pos, val := range code {
			gray[val] = pos
		}
		if len(code) > 100 {
			break
		}
		code = reflect(code)
	}
}

func reflect(s []string) []string {
	reply := make([]string, len(s)*2)
	for j := 0; j < len(s); j++ {
		reply[j] = "0" + s[j]
	}
	for j := 0; j < len(s); j++ {
		reply[len(s)+j] = "1" + s[len(s)-j-1]
	}
	return reply
}

func translate(line string) string {
	var reply []string
	for _, chunk := range strings.Split(line, " | ") {
		reply = append(reply, fmt.Sprintf("%d", gray[chunk]))
	}
	return strings.Join(reply, " | ")
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
		fmt.Println(translate(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
