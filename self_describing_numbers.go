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
	for _, s := range stringsFromFile() {
		if isSelfDescribing(s) {
			fmt.Println("1")
		} else {
			fmt.Println("0")
		}
	}
}

func isSelfDescribing(s string) bool {
	counts := make(map[int]int)   // how many times have we seen this digit
	expected := make(map[int]int) // how many times do we expect to see this digit?

	for pos, c := range s {
		n := atoi(string(c))
		expected[pos] = n
		counts[n] += 1
	}
	for pos, n := range expected {
		if counts[pos] != n {
			return false
		}
	}

	return true
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func stringsFromFile() []string {
	return strings.Fields(string(dataFromFile()))
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
