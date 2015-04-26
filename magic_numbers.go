package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func magic(s string) {
	chunks := strings.Fields(s)
	first := true
	for a, b := atoui(chunks[0]), atoui(chunks[1]); a <= b; a++ {
		if isMagic(digits(a)) {
			if !first {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", a)
			first = false
		}
	}
	if first {
		fmt.Println("-1")
	} else {
		fmt.Println()
	}
}

func atoui(s string) uint64 {
	n, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

var digitBuf [10]uint16

func digits(n uint64) []uint16 {
	reply := digitBuf[:0]
	for ; n > 0; n /= 10 {
		reply = append(reply, uint16(n%10))
	}
	for l, r := 0, len(reply)-1; l < r; l, r = l+1, r-1 {
		reply[l], reply[r] = reply[r], reply[l]
	}
	return reply
}

func isMagic(digit []uint16) (reply bool) {
	var pos uint16
	var used uint16 = 1 << digit[pos]
	var visited uint16 = 1 << pos
	steps := 1

	for {
		pos = (pos + digit[pos]) % uint16(len(digit))
		if (visited&(1<<pos))|(used&(1<<digit[pos])) != 0 {
			break
		}
		used |= 1 << digit[pos]
		visited |= 1 << pos
		steps++
	}
	return pos == 0 && steps == len(digit)
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
		magic(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
