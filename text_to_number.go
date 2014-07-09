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
		fmt.Println(textToNum(line))
	}
}

var values = map[string]int{
	"zero":      0,
	"one":       1,
	"two":       2,
	"three":     3,
	"four":      4,
	"five":      5,
	"six":       6,
	"seven":     7,
	"eight":     8,
	"nine":      9,
	"ten":       10,
	"eleven":    11,
	"twelve":    12,
	"thirteen":  13,
	"fourteen":  14,
	"fifteen":   15,
	"sixteen":   16,
	"seventeen": 17,
	"eighteen":  18,
	"nineteen":  19,
	"twenty":    20,
	"thirty":    30,
	"forty":     40,
	"fifty":     50,
	"sixty":     60,
	"seventy":   70,
	"eighty":    80,
	"ninety":    90,
}

var multipliers = map[string]int{
	"million":  1000000,
	"thousand": 1000,
	"hundred":  100,
}

func textToNum(line string) int {
	negative := false
	s := make([]int, 0)
	for _, f := range strings.Fields(line) {
		if f == "negative" {
			negative = true
		} else if mul, ok := multipliers[f]; ok {
			for j := len(s) - 1; j >= 0; j-- {
				if s[j] > mul {
					break
				}
				s[j] *= mul
			}
		} else if val, ok := values[f]; ok {
			s = append(s, val)
		} else {
			log.Fatal(f)
		}
	}
	total := 0
	for _, n := range s {
		total += n
	}
	if negative {
		total *= -1
	}
	return total
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
