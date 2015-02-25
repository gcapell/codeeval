package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func luhn(line string) bool {
	double := false
	sum := 0
	for p := len(line) - 1; p >= 0; p-- {
		r := line[p]
		if r == ' ' {
			continue
		}
		n := r - '0'
		if double {
			n *= 2
			if n > 9 {
				n = n%10 + 1
			}
		}
		sum += int(n)
		double = !double
	}
	return sum%10 == 0
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
		if luhn(scanner.Text()) {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
