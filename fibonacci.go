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
	targetn, max := atoi(targets)
	values := make(map[int]int)
	for _, n := range targetn {
		values[n] = 0
	}

	for a, b, nth := 1, 1, 2; nth <= max; a, b, nth = b, a+b, nth+1 {
		if _, ok := values[nth]; ok {
			values[nth] = b
		}
	}

	for _, j := range targetn {
		switch j {
		case 0:
			fmt.Println(0)
		case 1:
			fmt.Println(1)
		case 2:
			fmt.Println(1)
		default:
			fmt.Println(values[j])
		}
	}
}

func atoi(s []string) ([]int, int) {
	reply := make([]int, len(s))
	max := 0
	for j := range s {
		if n, err := strconv.Atoi(s[j]); err != nil {
			log.Fatal(err)
		} else {
			reply[j] = n
			if n > max {
				max = n
			}
		}
	}
	return reply, max
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
