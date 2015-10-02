package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

func lucky(line string) *big.Int {
	n, err := strconv.Atoi(line)
	if err != nil {
		log.Fatal(err)
	}
	return sumSquared(exp(n / 2))
}

type poly []*big.Int

func sumSquared(p poly) *big.Int {
	reply, square := big.NewInt(0), big.NewInt(0)
	for _, n := range p {
		reply.Add(reply, square.Mul(n, n))
	}
	return reply
}

var (
	one      = big.NewInt(1)
	expCache = map[int]poly{
		1: []*big.Int{one, one, one, one, one, one, one, one, one, one},
	}
)

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
	reply := make([]*big.Int, len(a)+len(b)-1)
	for j := range reply {
		reply[j] = big.NewInt(0)
	}
	prod := big.NewInt(0)
	for bPos, bVal := range b {
		for aPos, aVal := range a {
			p := reply[aPos+bPos]
			p.Add(p, prod.Mul(aVal, bVal))
		}
	}
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
