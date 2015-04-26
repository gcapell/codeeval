package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func efficientDelivery(s string) {
	ch := strings.Fields(s)
	tankers, oil := parseTankers(ch[0]), atoi(ch[1])

	deliveries, shortFall := deliver(tankers, oil)
	if len(deliveries) == 0 {
		fmt.Println(shortFall)
		return
	}
	for _, d := range deliveries {
		fmt.Printf("[")
		for pos := 0; pos < len(d); pos++ {
			if pos != 0 {
				fmt.Printf(",")
			}
			fmt.Printf("%d", d[len(d)-pos-1])
		}
		fmt.Printf("]")
	}
	fmt.Println()
}

func deliver(tankers []int, oil int) ([][]int, int) {
	t, tankers := tankers[0], tankers[1:]
	if len(tankers) == 0 {
		extra := oil % t
		if extra == 0 {
			return [][]int{[]int{oil / t}}, 0
		}
		return nil, t - extra
	}
	var reply [][]int
	minRem := t - oil%t
	for j := 0; oil >= 0; j, oil = j+1, oil-t {
		deliveries, rem := deliver(tankers, oil)
		for _, d := range deliveries {
			reply = append(reply, append(d, j))
		}
		if rem < minRem {
			minRem = rem
		}
	}
	return reply, minRem
}

var digits = regexp.MustCompile("[0-9]+")

func parseTankers(s string) []int {
	var reply []int
	for _, m := range digits.FindAllString(s, -1) {
		reply = append(reply, atoi(m))
	}
	return reply
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
	for scanner.Scan() {
		efficientDelivery(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
