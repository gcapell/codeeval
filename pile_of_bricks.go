package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	for line := range linesFromFilename() {
		hole, bricks := parse(line)

		passing := fitList(hole, bricks)
		if len(passing) > 0 {
			fmt.Println(csv(passing))
		} else {
			fmt.Println("-")
		}
	}
}

func csv(a []int) string {
	s := make([]string,len(a))
	for pos, n := range a {
		s[pos] = strconv.Itoa(n)
	}
	return strings.Join(s, ",")
}

func parse(line string) (hole, []brick) {
	fields := strings.Split(line, "|")
	return parseHole(fields[0]), parseBricks(fields[1])
}

type hole []int

var holeRE = regexp.MustCompile(`\[(-?[0-9]+),(-?[0-9]+)\]`)

func parseHole(s string) hole {
	m := holeRE.FindAllStringSubmatch(s, -1)
	a, b := m[0], m[1]
	return hole{delta(a[1], b[1]), delta(a[2], b[2])}
}

type brick struct {
	idx int
	d   []int
}

var brickRE = regexp.MustCompile(`\(([0-9]+) \[(-?[0-9]+),(-?[0-9]+),(-?[0-9]+)\] \[(-?[0-9]+),(-?[0-9]+),(-?[0-9]+)\]\)`)

func parseBricks(s string) []brick {
	fields := strings.Split(s, ";")
	bricks := make([]brick, len(fields))
	for pos, f := range fields {
		m := brickRE.FindStringSubmatch(f)
		bricks[pos] = brick{atoi(strings.TrimSpace(m[1])),
			[]int{delta(m[2], m[5]), delta(m[3], m[6]), delta(m[4], m[7])}}
	}
	return bricks
}

func delta(a, b string) int {
	d := atoi(a) - atoi(b)
	if d < 0 {
		return -d
	}
	return d
}

func fitList(h hole, bricks []brick) []int {
	var reply []int
	for _, b := range bricks {
		if fit(h, b.d) {
			reply = append(reply, b.idx)
		}
	}
	sort.Ints(reply)
	
	return reply
}

func fit(hole []int, brick []int) bool {
	sort.Ints(hole)
	sort.Ints(brick)

	if hole[0] < brick[0] {
		return false
	}
	if hole[1] >= brick[1] {
		return true
	}

	// Diagonal? see http://www.jstor.org/stable/2691523
	a, b, p, q := float64(hole[1]), float64(hole[0]), float64(brick[1]), float64(brick[0])
	return b >= (2*p*q*a+(p*p-q*q)*math.Sqrt(p*p+q*q-a*a))/(p*p+q*q)
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
			c <- strings.TrimSpace(line)
		}
		close(c)
	}()
	return c
}
