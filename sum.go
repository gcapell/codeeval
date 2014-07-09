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
	data := fileFromFilename()
	targets := strings.Fields(string(data))
	targetn := atoi(targets)
	sum := 0
	for _, n := range targetn {
		sum += n
	}
	fmt.Println(sum)
}

func atoi(s []string) []int {
	reply := make([]int, len(s))
	for j := range s {
		if n, err := strconv.Atoi(s[j]); err != nil {
			log.Fatal(err)
		} else {
			reply[j] = n
		}
	}
	return reply
}

func fileFromFilename() []byte {
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
