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

var letters = map[int]string{
	1:     "I",
	5:     "V",
	10:    "X",
	50:    "L",
	100:   "C",
	500:   "D",
	1000:  "M",
	5000:  "_",
	10000: "_",
}

func main() {

	for line := range linesFromFilename() {
		n, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			log.Fatal(err)
		}
		toRoman(n)
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

func toRoman(n int) {
	i, v, x := 1, 5, 10
	reply := make([]string, 0)
	for n > 0 {
		var digit int
		digit, n = n%10, n/10
		if digit != 0 {
			letter := convertDigit(digit, letters[i], letters[v], letters[x])
			reply = append(reply, letter)
		}
		i, v, x = i*10, v*10, x*10
	}

	for j := len(reply) - 1; j >= 0; j-- {
		fmt.Print(reply[j])
	}
	fmt.Println()
}

func convertDigit(n int, i, v, x string) string {
	switch n {
	case 1:
		return i
	case 2:
		return i + i
	case 3:
		return i + i + i
	case 4:
		return i + v
	case 5:
		return v
	case 6:
		return v + i
	case 7:
		return v + i + i
	case 8:
		return v + i + i + i
	case 9:
		return i + x
	}
	return ""
}
