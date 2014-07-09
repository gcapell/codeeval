package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var morse = map[string]string{
	".-":   "A",
	"-...": "B",
	"-.-.": "C",
	"-..":  "D",
	".":    "E",
	"..-.": "F",
	"--.":  "G",
	"....": "H",
	"..":   "I",
	".---": "J",
	"-.-":  "K",
	".-..": "L",
	"--":   "M",
	"-.":   "N",
	"---":  "O",
	".--.": "P",
	"--.-": "Q",
	".-.":  "R",
	"...":  "S",
	"-":    "T",
	"..-":  "U",
	"...-": "V",
	".--":  "W",
	"-..-": "X",
	"-.--": "Y",
	"--..": "Z",

	".----": "1",
	"..---": "2",
	"...--": "3",
	"....-": "4",
	".....": "5",
	"-....": "6",
	"--...": "7",
	"---..": "8",
	"----.": "9",
	"-----": "0",
}

func main() {
	for line := range linesFromFilename() {
		fmt.Println(convertMorse(line))
	}
}

func convertMorse(line string) string {
	convertedWords := make([]string,0)

	words := strings.Split(line, "  ")
	for _, word := range words {
		var convertedWord string

		letters := strings.Split(word, " ")
		for _, letter := range letters {
			text, ok := morse[letter]
			if !ok {
				log.Fatalf("no entry for %q", letter)
			}
			// fmt.Print(letter, text)
			convertedWord += text
		}
		convertedWords = append(convertedWords, convertedWord)
	}
	return strings.Join(convertedWords, " ")
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
