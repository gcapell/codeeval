package main

import (
	"bufio"
	"log"
	"math/big"
	"os"
	"strconv"
)

func estimate(terms int64) string {
	var top, total big.Int
	for j := int64(0); j < terms; j++ {
		term(j, terms, &top)
		total.Add(&total, &top)
	}
	var exp big.Int
	exp.Exp(c2, big.NewInt(12*terms+4), nil)
	// fmt.Println(exp.Text(10), total.Text(10))
	var r big.Rat
	r.SetFrac(&exp, &total)
	return r.FloatString(int(float64(terms) * 1.5))
}

var (
	c2 = big.NewInt(2)
	c3 = big.NewInt(3)
)

func term(j, n int64, top *big.Int) {
	top.Binomial(2*j, j)
	top.Exp(top, c3, nil) // choose^3
	top.Mul(top, big.NewInt(42*j+5))

	var exp big.Int
	exp.Exp(c2, big.NewInt(12*(n-j)), nil)
	top.Mul(top, &exp)
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
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

	digits := estimate(3333)

	for scanner.Scan() {
		n := atoi(scanner.Text())
		if n == 1 {
			n = 0
		}

		println(string(digits[n]))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
