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
		fmt.Println(absurd(line))
	}
}

func absurd(line string) int {
	f := strings.Split(line, ";")
	count, _ := strconv.Atoi(f[0])
	expected := (count - 1) * (count - 2) / 2

	nums := strings.Split(f[1], ",")
	sum := 0
	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		sum += n
	}

	return sum - expected
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
