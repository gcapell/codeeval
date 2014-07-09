package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	for n := range intsFromFilename() {
		fmt.Println(bitcount(n))
	}
}

func bitcount(n int) int {
	bits := 0
	for n>0 {
		bits += n&1
		n >>= 1
	}
	return bits
}

func intsFromFilename() chan int {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan int)

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
			n, err := strconv.Atoi(strings.TrimSpace(line))
			if err != nil {
				log.Fatal(err)
			}
			c <- n
		}
		close(c)
	}()
	return c
}

