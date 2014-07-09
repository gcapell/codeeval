package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	for _, n := range numsFromFile() {
		if isArmstrong(n) {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}

func isArmstrong(n int) bool {
	dd := digits(n)
	var sum float64
	for _, d := range dd {
		sum += math.Pow(float64(d), float64(len(dd)))
	}
	return int(sum) == n
}

func digits(n int) []int {
	reply := make([]int, 0)
	for n > 0 {
		digit := n % 10
		n = n / 10
		reply = append(reply, digit)
	}
	return reply
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
