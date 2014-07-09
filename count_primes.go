package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type pair struct{ a, b uint32 }

func main() {

	var ranges []pair
	var maxN uint32

	for line := range linesFromFilename() {
		f := strings.Split(line, ",")
		p := pair{atoi(f[0]), atoi(f[1])}
		ranges = append(ranges, p)
		if p.b > maxN {
			maxN = p.b
		}
	}

	b := Sieve(uint32(maxN))
	for _, p := range ranges {
		fmt.Println(b.CountBetween(p.a, p.b))
	}
}

type Bit6_32 []uint32

func NewBit6_32(n uint32) Bit6_32 {
	return Bit6_32(make([]uint32, n/96+1))
}

func (b Bit6_32) CountBetween(x, y uint32) int {
	// Could do funky stuff counting bits, but let's start with really dumb

	primes := 0
	for j := x; j <= y; j++ {
		if b.IsPrime(j) {
			primes++
		}
	}
	return primes
}

func (b Bit6_32) IsPrime(n uint32) bool {
	if n == 2 || n == 3 {
		return true
	}
	switch n % 6 {
	case 1:
		fallthrough
	case 5:
		return !b.Test(n)
	}
	return false
}

func (b Bit6_32) List() []int {
	var out []int

	n := 1
	for _, word := range b {
		for bit := uint32(0); bit < 32; bit++ {
			if word&(1<<bit) != 0 {
				out = append(out, n)
			}
			if bit%2 == 0 {
				n += 4
			} else {
				n += 2
			}
		}
	}
	return out
}

func (b Bit6_32) String() string {
	return fmt.Sprintf("%v->%v", ([]uint32)(b), b.List())
}

func mod6_32(n uint32) (uint32, uint32, bool) {
	div, rem := n/6, n%6
	bit := div * 2
	switch rem {
	case 1:
	case 5:
		bit++
	default:
		return 0, 0, false
	}

	word := bit / 32
	bit32 := bit % 32

	return word, bit32, true
}

func (b Bit6_32) Set(n uint32) {
	word, bit, fits := mod6_32(n)
	if fits {
		b[word] |= 1 << bit
	}
}

func (b Bit6_32) Test(n uint32) bool {
	word, bit, fits := mod6_32(n)
	if !fits {
		log.Fatalf("expected %d %% 6 to be 1 or 5", n)
	}
	return b[word]&(1<<bit) != 0
}

func Sieve(n uint32) Bit6_32 {
	b := NewBit6_32(n)
	for j := range mod6_stream(uint32(math.Sqrt(float64(n)))) {
		if !b.Test(j) {
			for k := j * j; k < n; k += j {
				b.Set(k)
			}
		}
	}

	return b
}

func mod6_stream(n uint32) chan uint32 {
	c := make(chan uint32)
	go func() {
		for j := uint32(5); j <= n; {
			c <- j
			j += 2
			c <- j
			j += 4
		}
		close(c)
	}()
	return c
}

func atoi(s string) uint32 {
	n, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	return uint32(n)
}

func linesFromFilename() chan string {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan string)

	go func() {
		reader := bufio.NewReader(f)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					log.Fatal(err)
				}
				break
			}
			line = strings.TrimSpace(line)
			if len(line) > 0 {
				c <- line
			}
		}
		close(c)
	}()
	return c
}
