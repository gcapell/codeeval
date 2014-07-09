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

const DEBUG = false

var denominations = []struct {
	amt int
	s   string
}{
	{10000, "ONE HUNDRED"},
	{5000, "FIFTY"},
	{2000, "TWENTY"},
	{1000, "TEN"},
	{500, "FIVE"},
	{200, "TWO"},
	{100, "ONE"},
	{50, "HALF DOLLAR"},
	{25, "QUARTER"},
	{10, "DIME"},
	{5, "NICKEL"},
	{1, "PENNY"},
}

func main() {

	for line := range linesFromFilename() {
		change := calcChange(line)
		switch {
		case change == 0:
			fmt.Println("ZERO")
		case change < 0:
			fmt.Println("ERROR")
		default:
			fmt.Println(calc(change))
		}
	}
}

func calc(ch int) string {
	var change []string

	for _, d := range denominations {
		for ch >= d.amt {
			change = append(change, d.s)
			ch -= d.amt
		}
		if ch == 0 {
			break
		}
	}
	return strings.Join(change, ",")
}

func calcChange(line string) int {
	f := strings.Split(line, ";")
	fn := func(s string) int {
		g := strings.Split(s, ".")
		cents := atoi(g[0]) * 100
		if len(g) > 1 {
			cents += atoi(g[1])
		}
		return cents
	}
	return fn(f[1]) - fn(f[0])
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
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
