package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	primeList = []uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
	
	nums, primes, lastNums uint64
)

func init() {
	for _, p := range primeList {
		primes |= 1 << p
	}
	// The last number we choose must be one less than a prime.
	// Any time we don't have any of these left, we can give up.
	lastNums = primes >> 1  
}

func necklaces(size uint) int {
	nums = (1 << (size + 1)) - 2
	return solutions(1)
}

func remaining(n uint64) []int {
	var reply []int
	for pos := 0; n != 0; pos, n = pos+1, n>>1 {
		if n&1 != 0 {
			reply = append(reply, pos)
		}
	}

	return reply
}

func solutions(n uint) int {
	nums &= ^(1 << n)
	defer func() {
		nums |= 1 << n
	}()

	if nums == 0 {
		return 1
	}
	if nums&lastNums == 0 {
		return 0
	}
	possible := (primes >> n) & nums
	total := 0
	for possible != 0 {
		var next uint
		next, possible = lsb(possible)
		total += solutions(next)
	}
	return total
}

var bit2pos = make(map[uint64]uint)

func init() {
	for j := uint(1); j <= 18; j++ {
		bit2pos[1<<j] = j
	}
}

// lsb returns the position of the least significant bit of n (and n with lsb cleared).
// For example lsb(0b0110) -> (1, 0b0100)
func lsb(n uint64) (uint, uint64) {
	cleared := n & (n - 1)
	return bit2pos[n&^cleared], cleared
}

func atoui(s string) uint {
	n, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	return uint(n)
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
		fmt.Println(necklaces(atoui(scanner.Text())))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
