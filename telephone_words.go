package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	for line := range linesFromFilename() {
		expand2(line)
		fmt.Println()
	}
}

func expand2(line string) {
	idx := make([]int, len(line))

	first := true
	reply := initialWord(line)
	for {
		fmt.Printf("%s%s", comma(&first), reply)
		if clocked := inc(idx, line, reply); clocked {
			break
		}
	}
}

func initialWord(line string) []byte {
	reply := make([]byte, len(line))
	for pos := range line {
		reply[pos] = letters[line[pos]][0]
	}
	return reply
}

func inc(idx []int, line string, word []byte) bool {
	for bit := len(idx) - 1; bit >= 0; bit-- {
		thisLetters := letters[line[bit]]
		n := (idx[bit] + 1) % len(thisLetters)
		idx[bit] = n
		word[bit] = thisLetters[n]
		if n != 0 {
			return false
		}
	}
	return true
}

func expand(line string) chan string {
	ch := make(chan string)
	go func() {

		if len(line) == 1 {
			for _, c := range letters[line[0]] {
				ch <- string(c)
			}
		} else {
			for _, c := range letters[line[0]] {
				for s := range expand(line[1:]) {
					ch <- string(c) + s
				}
			}
		}

		close(ch)
	}()
	return ch
}

func comma(first *bool) string {
	if !*first {
		return ","
	}
	*first = false
	return ""
}

var letters = map[byte]string{
	'0': "0",
	'1': "1",
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func linesFromFilename() chan string {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan string)

	go func() {
		reader := bufio.NewReader(f)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					log.Fatal(err)
				}
				break
			}
			c <- strings.TrimSpace(line)
		}
		close(c)
	}()
	return c
}
