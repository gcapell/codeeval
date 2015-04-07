package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type match struct {
	errors int
	s      string
}
type ByMatch []match

func (a ByMatch) Len() int      { return len(a) }
func (a ByMatch) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByMatch) Less(i, j int) bool {
	if a[i].errors < a[j].errors {
		return true
	}
	if a[i].errors > a[j].errors {
		return false
	}
	return a[i].s < a[j].s
}

func matches(line string) string {
	chunks := strings.Fields(line)
	pattern, errors, dna := []byte(chunks[0]), atoi(chunks[1]), []byte(chunks[2])
	f := make([]int, len(pattern)+1)
	var matches []match
	for start, end := 0, len(pattern); end <= len(dna); start, end = start+1, end+1 {
		dist := levenshtein(pattern, dna[start:end], f)
		if dist <= errors {
			matches = append(matches, match{dist, string(dna[start:end])})
		}
	}
	if len(matches) == 0 {
		return "No match"
	}
	sort.Sort(ByMatch(matches))
	var reply []string
	for _, m := range matches {
		reply = append(reply, m.s)
	}
	return strings.Join(reply, " ")
}

// http://en.wikibooks.org/wiki/Algorithm_Implementation/Strings/Levenshtein_distance#Go, basically.
func levenshtein(a, b []byte, f []int) int {
	for j := range f {
		f[j] = j
	}

	for _, ca := range a {
		j := 1
		fj1 := f[0] // fj1 is the value of f[j - 1] in last iteration
		f[0]++
		for _, cb := range b {
			mn := min(f[j]+1, f[j-1]+1) // delete & insert
			if cb != ca {
				mn = min(mn, fj1+1) // change
			} else {
				mn = min(mn, fj1) // matched
			}

			fj1, f[j] = f[j], mn // save f[j] to fj1(j is about to increase), update f[j] to mn
			j++
		}
	}

	return f[len(f)-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
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
		fmt.Println(matches(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
