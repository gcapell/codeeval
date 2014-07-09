package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	for line := range linesFromFilename() {
		n, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			log.Fatal(err)
		}
		sum := 0
		for n > 0 {
			var digit int
			n, digit = n/10, n%10
			sum += digit
		}
		fmt.Println(sum)
	}
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
			c <- line
		}
		close(c)
	}()
	return c
}
