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
	for p := range fileToPairs() {
		fmt.Println(csvFmt(intersect(p.a, p.b)))
	}
}

type csvFmt []int

func (x csvFmt) String() string {
	s := make([]string, len(x))
	for k, v := range x {
		s[k] = strconv.Itoa(v)
	}
	return strings.Join(s, ",")
}

func intersect(a, b []int) []int {
	c := make([]int, 0, len(a))
	for ap, bp := 0, 0; ap < len(a) && bp < len(b); {
		switch {
		case a[ap] < b[bp]:
			ap += 1
		case a[ap] > b[bp]:
			bp += 1
		default:
			c = append(c, a[ap])
			ap += 1
			bp += 1
		}
	}
	return c
}

type pair struct{ a, b []int }

func newPair(s string) (pair, error) {
	sides := strings.Split(s, ";")
	var p pair
	if len(sides) != 2 {
		return p, fmt.Errorf("expected two ;-delimited lists, got %q", s)
	}
	a, err := csv(sides[0])
	if err != nil {
		return p, err
	}
	b, err := csv(sides[1])
	if err != nil {
		return p, err
	}
	p.a, p.b = a, b
	return p, nil
}

func csv(s string) ([]int, error) {
	ss := strings.Split(strings.TrimSpace(s), ",")
	n := make([]int, len(ss))
	for j := range ss {
		if i, err := strconv.Atoi(ss[j]); err != nil {
			return nil, err
		} else {
			n[j] = i
		}
	}
	return n, nil
}

func fileToPairs() chan pair {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan pair)

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
			p, err := newPair(strings.TrimSpace(line))
			if err != nil {
				log.Fatal(err)
			}
			c <- p
		}
		close(c)
	}()
	return c
}
