package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var alpha = ` !"#$%&'()*+,-./0123456789:<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`

func gronsfeld(line string) string {
	fields := strings.Split(line, ";")
	key, cipher := fields[0], fields[1]
	var plain string
	for pos, r := range cipher {
		p := strings.IndexRune(alpha, r)
		p -= int(key[pos % len(key)]) - '0'
		if p<0 {
			p += len(alpha)
		}
		plain += string(alpha[p])
	}
	return plain
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
		fmt.Println(gronsfeld(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
