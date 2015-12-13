package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var key = map[byte]byte{
	'a': 'u',
	'b': 'v',
	'c': 'w',
	'd': 'x',
	'e': 'y',
	'f': 'z',
	'g': 'n',
	'h': 'o',
	'i': 'p',
	'j': 'q',
	'k': 'r',
	'l': 's',
	'm': 't',
	'n': 'g',
	'o': 'h',
	'p': 'i',
	'q': 'j',
	'r': 'k',
	's': 'l',
	't': 'm',
	'u': 'a',
	'v': 'b',
	'w': 'c',
	'x': 'd',
	'y': 'e',
	'z': 'f',
}

func crypt(line string) string {
	out := make([]byte, len(line))
	for i := 0; i < len(line); i++ {
		var ok bool
		if out[i], ok = key[line[i]]; !ok {
			out[i] = line[i]
		}
	}
	return string(out)
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
		fmt.Println(crypt(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
