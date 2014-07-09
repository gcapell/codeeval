package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	// We've basically been asked for a bunch of Fibonacci numbers.
	// Generate up until the biggest number required, storing
	// the ones we want, then print.

	nums := numsFromFilename()

	max := 0
	lookup := make(map[int]string)
	for _, n := range nums {
		if n > max {
			max = n
		}
		lookup[n] = ""
	}

	a, b := big.NewInt(1), big.NewInt(1)

	for j := 1; j <= max; j++ {
		c := b.Add(a, b)
		a, b = c, a
		if _, ok := lookup[j]; ok {
			lookup[j] = b.String()
		}
		// fmt.Println(a,b)
	}

	for _, n := range nums {
		fmt.Println(lookup[n])
	}
}

func numsFromFilename() []int {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fields := strings.Fields(string(bytes))
	nums := make([]int, len(fields))
	for pos, f := range fields {
		var err error
		nums[pos], err = strconv.Atoi(f)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nums
}
