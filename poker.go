package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	for line := range linesFromFilename() {
		left, right := parse(line[:15]), parse(line[15:])
		fmt.Println(left.winner(right))
	}
}

// hand types
type handType int

const (
	highCard handType = iota
	pair
	twoPair
	trips
	straight
	flush
	fullHouse
	quads
	straightFlush
)

type hand struct {
	t    handType
	kick []int
}

func (h hand) winner(o hand) string {
	if h.t > o.t {
		return "left"
	}
	if h.t < o.t {
		return "right"
	}
	for pos := range h.kick {
		if h.kick[pos] > o.kick[pos] {
			return "left"
		}
		if h.kick[pos] < o.kick[pos] {
			return "right"
		}
	}
	return "none"
}

func parse(line string) hand {
	cards := strings.Fields(line)
	isFlush := true
	var prevSuit rune
	counts := make(map[int]int)
	for pos, c := range cards {
		val, suit := parseCard(c)
		if pos != 0 && suit != prevSuit {
			isFlush = false
		}
		prevSuit = suit
		counts[val]++
	}

	// occurences -> list of cards, e.g. pairs -> [3,4]
	reverseCounts := make(map[int][]int)
	for k, v := range counts {
		reverseCounts[v] = append(reverseCounts[v], k)
	}

	isStraight := detectStraight(reverseCounts[1])

	if isStraight && isFlush {
		return hand{straightFlush, reverseCounts[1]}
	}

	if len(reverseCounts[4]) != 0 {
		return hand{quads, reverseCounts[4]}
	}
	if len(reverseCounts[3]) > 0 && len(reverseCounts[2]) > 0 {
		return hand{fullHouse, []int{reverseCounts[3][0], reverseCounts[2][0]}}
	}
	if isFlush {
		return hand{flush, reverseCounts[1]}
	}
	if isStraight {
		return hand{straight, reverseCounts[1]}
	}
	if len(reverseCounts[3]) > 0 {
		return hand{trips, append(reverseCounts[3], reverseCounts[1]...)}
	}
	if len(reverseCounts[2]) == 2 {
		sort.Sort(Descending(reverseCounts[2]))
		return hand{twoPair, append(reverseCounts[2], reverseCounts[1][0])}
	}
	if len(reverseCounts[2]) == 1 {
		return hand{pair, append(reverseCounts[2], reverseCounts[1]...)}
	}
	return hand{highCard, reverseCounts[1]}
}

var courtVals = map[byte]int{
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func parseCard(s string) (int, rune) {
	v, ok := courtVals[s[0]]
	if !ok {
		var err error
		v, err = strconv.Atoi(s[:1])
		if err != nil {
			log.Fatalf("%s from %q", err, s)
		}
	}
	return v, rune(s[1])
}

func detectStraight(c []int) bool {
	sort.Sort(Descending(c))
	if len(c)<5 {
		return false
	}
	for j := 1; j < 5; j++ {
		if c[j] != c[j-1]-1 {
			return false
		}
	}
	return true
}

type Descending []int

func (a Descending) Len() int           { return len(a) }
func (a Descending) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Descending) Less(i, j int) bool { return a[i] > a[j] }

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
