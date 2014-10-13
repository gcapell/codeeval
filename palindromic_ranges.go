package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func ranges(line string) int {
	var left, right int
	fmt.Sscanf(line, "%d %d", &left, &right)
	x := palRange(left, right)
	// fmt.Println(left, right, x)

	total := 0
	for l := left; l <= right; l++ {
		for r := l; r <= right; r++ {
			if interesting(l, r, x) {
				total++
			}
		}
	}
	return total
}

func interesting(l, r int, pals []int) bool {
	count := 0
	for _, p := range pals {
		if p < l {
			continue
		}
		if p > r {
			break

		}
		count++
	}
	return count%2 == 0
}

func palRange(l, r int) []int {
	g := ascendingPals(l)
	var reply []int
	for {
		p := g()
		if p > r {
			break
		}
		reply = append(reply, p)
	}
	return reply
}

// palLess returns function to generate palindromes >=  n in ascending order.
func ascendingPals(n int) func() int {
	digits := int(math.Floor(math.Log10(float64(n)))) + 1
	rightDigits := digits / 2
	shift := int(math.Pow10(rightDigits))
	left := n / shift
	decade := int(math.Pow10(digits - rightDigits))

	gen := func() int {
		reply := left*shift + reverse(left, digits)
		left++
		if left == decade {
			digits++
			rightDigits = digits / 2
			shift = int(math.Pow10(rightDigits))
			max := int(math.Pow10(digits - rightDigits - 1))
			left = max
			decade = max * 10
		}
		return reply
	}

	// skip past first one if too small
	if left*shift+reverse(left, digits) < n {
		gen()
	}

	return gen
}

func reverse(n, digits int) int {
	if digits%2 == 1 {
		n /= 10
	}
	var reply int
	for ; n > 0; n /= 10 {
		reply = reply*10 + n%10
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
		fmt.Println(ranges(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
