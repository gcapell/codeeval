package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type permute struct {
	p int
	r rune
}

func (p permute) String() string {
	return fmt.Sprintf("(%q, %d)", p.r, p.p)
}

type ByR []permute

func (a ByR) Len() int      { return len(a) }
func (a ByR) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByR) Less(i, j int) bool {
	if a[i].r < a[j].r {
		return true
	}
	if a[i].r > a[j].r {
		return false
	}
	return a[i].p < a[j].p
}

func ibwt(line string) string {
	fmt.Printf("%q\n", line)
	permutes := make([]permute, 0, len(line))
	for p, r := range line {
		permutes = append(permutes, permute{p, r})
	}
	// fmt.Println(permutes)
	sort.Sort(ByR(permutes))
	fmt.Println(permutes)
	k := strings.Index(line, "$")
	if k == -1 {
		log.Fatal("no end", line)
	}
	var b bytes.Buffer
	for _ = range permutes {
		p := permutes[k]
		b.WriteRune(p.r)
		k = p.p
	}
	return b.String()
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
		line := strings.TrimRight(scanner.Text(), "|")
		fmt.Println(ibwt(line))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
