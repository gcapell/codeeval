package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	for ns := range csvIntsFromFilename() {
		fmt.Println(bitpositions(ns[0], ns[1], ns[2]))
	}
}

func bitpositions(n, p1, p2 uint) string {
	b1 := n & (1 << (p1 - 1))
	b2 := n & (1 << (p2 - 1))
	if (b1 == 0) == (b2 == 0) {
		return "true"
	}
	return "false"
}

func csvInts(s string) ([]uint, error) {
	ss := strings.Split(strings.TrimSpace(s), ",")
	n := make([]uint, len(ss))
	for j := range ss {
		if i, err := strconv.Atoi(ss[j]); err != nil {
			return nil, err
		} else {
			n[j] = uint(i)
		}
	}
	return n, nil
}

func csvIntsFromFilename() chan []uint {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan []uint)

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
			ns, err := csvInts(line)
			if err != nil {
				log.Fatal(err)
			}
			c <- ns
		}
		close(c)
	}()
	return c
}
