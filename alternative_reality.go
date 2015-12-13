package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type prob struct{ n, d int }

var cache = make(map[prob]int)

func smaller(d int) int {
	switch d {
	case 50:
		return 25
	case 25:
		return 10
	case 10:
		return 5
	default:
		return 1
	}
}

// How many ways can we change `n` using denoms <= d?
func change(n, d int) int {
	if d == 1 || n == 0 {
		return 1
	}
	key := prob{n, d}
	if v, ok := cache[key]; ok {
		return v
	}

	s := smaller(d)
	total := 0
	for ; n >= 0; n -= d {
		total += change(n, s)
	}
	cache[key] = total
	return total
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
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(change(n, 50))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
