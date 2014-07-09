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
		if n%2 == 0 {
			fmt.Println("1")
		} else {
			fmt.Println("0")
		}
	}
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
