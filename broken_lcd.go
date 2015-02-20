package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var digits = map[rune]uint8{
	'0': 0xfc, // 1,2,3,4,5,6,
	'1': 0x60, // 2,3,
	'2': 0xda, // 1,2,4,5,7
	'3': 0xf2, // 1,2,3,4,7,
	'4': 0x66, // 2,3,6,7
	'5': 0xb6, // 1,3,4,6,7,
	'6': 0xbe, // 1,3,4,5,6,7
	'7': 0xe0, // 1,2,3,
	'8': 0xfe, // 1,2,3,4,5,6,7
	'9': 0xf6, // 1,2,3,4,6,7
}

type lcd [12]uint8

func canPrint(line string) bool {
	fields := strings.Split(line, ";")
	l, num := parseLCD(fields[0]), parseNum(fields[1])
posLoop:
	for pos := 0; pos+len(num) < len(l); pos++ {
		for i, d := range num {
			if d&l[pos+i] != d {
				continue posLoop
			}
		}
		return true
	}
	return false
}

func parseLCD(s string) lcd {
	var reply lcd
	for pos, bits := range strings.Fields(s) {
		n, err := strconv.ParseUint(bits, 2, 8)
		if err != nil {
			log.Fatal(err, s)
		}
		reply[pos] = uint8(n)
	}
	return reply
}

func parseNum(s string) []uint8 {
	var reply []uint8
	addedDecimal := false
	for _, r := range s {
		if r == '.' {
			reply[len(reply)-1] |= 1
			addedDecimal = true
			continue
		}
		reply = append(reply, digits[r])
	}
	if !addedDecimal {
		reply[len(reply)-1] |= 1
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
		r := 0
		if canPrint(scanner.Text()) {
			r = 1
		}
		fmt.Println(r)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
