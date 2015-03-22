package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const size = 2 * 3 * 5 * 7

func ugly(s string) int {
	total := 0
	for m, cnt := range mods(s) {
		if m%2 == 0 || m%3 == 0 || m%5 == 0 || m%7 == 0 {
			total += cnt
		}
	}
	return total
}

func mods(s string) map[int]int {
	modsByDepth := make(map[int]map[int]int)
	for depth := 1; depth <= len(s); depth++ {
		n := atoi(s[:depth])
		d := map[int]int{n: 1}
		for subDepth := 1; subDepth < depth; subDepth++ {
			n := atoi(s[subDepth:depth])
			for m, cnt := range modsByDepth[subDepth] {
				dst1, dst2 := dst(m, n)
				d[dst1] += cnt
				d[dst2] += cnt
			}
		}
		modsByDepth[depth] = d
	}
	return modsByDepth[len(s)]
}

func dst(orig, delta int) (int, int) {
	dst1 := (orig + delta) % size
	dst2 := (orig - delta) % size
	if dst2 < 0 {
		dst2 += size
	}
	return dst1, dst2
}

func atoi(s string) int {
	n := 0
	for _, c := range s {
		n = (n*10 + int(c) - '0') % size
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
	for scanner.Scan() {
		fmt.Println(ugly(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
