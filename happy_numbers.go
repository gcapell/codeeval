package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	for _, n := range numsFromFile() {
		if isHappy(n) {
			fmt.Println("1")
		} else {
			fmt.Println("0")
		}
	}
}

func isHappy(n int) bool {
	seen := make(map[int]bool)

	for n != 1 && !seen[n] {
		seen[n] = true
		n = sumSquaredDigits(n)
	}
	return n == 1
}

func sumSquaredDigits(n int) int {
	sum := 0
	for n > 0 {
		d := n % 10
		n = n / 10
		sum += d * d
	}
	return sum
}

func numsFromFile() []int {
	fields := strings.Fields(string(dataFromFile()))

	reply := make([]int, len(fields))

	for j := range fields {
		if n, err := strconv.Atoi(fields[j]); err != nil {
			log.Fatal(err)
		} else {
			reply[j] = n
		}
	}
	return reply
}

func dataFromFile() []byte {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)

	}
	return data
}
