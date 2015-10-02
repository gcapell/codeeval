package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func lucky(line string) int64 {
	n, err := strconv.Atoi(line)
	if err != nil {
		log.Fatal(err)
	}
	return sumSquared(exp(n / 2))
}

type poly []int64

func sumSquared(p poly) int64 {
	var reply int64
	for _, n := range p {
		reply += n * n
	}
	return reply
}

var expCache = map[int]poly{
	1: []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}

func exp(n int) poly {
	if r, ok := expCache[n]; ok {
		return r
	}
	r := expWork(n)
	expCache[n] = r
	return r
}

var powers = []int{32, 16, 8, 4, 2}

func expWork(n int) poly {
	for _, p := range powers {
		if n == p {
			return mul(exp(n/2), exp(n/2))
		}
		if n > p {
			return mul(exp(p), exp(n-p))
		}
	}
	log.Fatal("notreached")
	return nil
}

func mul(a, b poly) poly {
	reply := make([]int64, len(a)+len(b) -1)
	for bPos, bVal := range b {
		for aPos, aVal := range a {
			reply[aPos+bPos] += aVal * bVal
		}
	}
	fmt.Println("mul", a, b, "=", reply)
	return reply
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
		fmt.Println(lucky(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
