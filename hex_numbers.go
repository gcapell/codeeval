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
	for _, n := range hexFromFile() {
		fmt.Println(n)
	}
}

func hexFromFile() []int {
	fields := strings.Fields(string(dataFromFile()))

	reply := make([]int, len(fields))

	for j := range fields {
		if n, err := strconv.ParseInt(fields[j], 16, 32); err != nil {
			log.Fatal(err)
		} else {
			reply[j] = int(n)
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
