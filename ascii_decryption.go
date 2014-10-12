package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func decrypt(line string) string {
	chunks := strings.Split(line, "|")
	length, lastLetter, cipherText := atoi(chunks[0]),
		strings.TrimSpace(chunks[1])[0], mapInt(strings.Fields(chunks[2]))
	if false {
		fmt.Println(length, lastLetter)
	}

	space := min(cipherText)
	delta := space - ' '

	var reply bytes.Buffer
	for _, c := range cipherText {
		reply.WriteByte(byte(c - int(delta)))
	}
	return reply.String()
}

func notCheating(cipherText []int, length, space int, lastLetter uint8) {
	words := wordsOfLength(cipherText, length, space)
	fmt.Println(words)
	found, pos := identical(cipherText, words, length, 0)
	if !found {
		panic("no duplicate")
	}
	delta := uint8(cipherText[pos+length-1]) - lastLetter
	fmt.Println(delta)
}

func identical(cipherText []int, words []int, length, off int) (bool, int) {
	splits := make(map[int][]int)

	for _, w := range words {
		c := cipherText[w+off]
		splits[c] = append(splits[c], w)
	}
	for _, group := range splits {
		if len(group) < 2 {
			continue
		}
		if length == off+1 {
			return true, group[0]
		}
		if found, pos := identical(cipherText, group, length, off+1); found {
			return found, pos
		}
	}
	return false, 0
}

func wordsOfLength(cipherText []int, length, space int) []int {
	var reply []int

	prev := 0
	for pos, c := range cipherText {
		if c != space {
			continue
		}
		if pos == prev+length {
			reply = append(reply, prev)
		}
		prev = pos + 1
	}
	if len(cipherText)-prev == length {
		reply = append(reply, prev)
	}
	return reply
}

func min(as []int) int {
	reply := as[0]
	for _, a := range as {
		if a < reply {
			reply = a
		}
	}
	return reply
}

func atoi(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return n
}

func mapInt(chunks []string) []int {
	var reply []int
	for _, s := range chunks {
		reply = append(reply, atoi(s))
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
		fmt.Println(decrypt(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
