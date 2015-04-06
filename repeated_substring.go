package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
)

type ByLex [][]byte

func (a ByLex) Len() int           { return len(a) }
func (a ByLex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLex) Less(i, j int) bool { return bytes.Compare(a[i], a[j]) == -1 }

func longestRepeatedSubstring(line string) string {
	sa := suffixArray([]byte(line))
	sort.Sort(ByLex(sa))

	var found []byte
	var earliest int

	for j := 0; j < len(sa)-1; j++ {
		a, b := sa[j], sa[j+1]
		longestPossible, maxLen := longestPossibleMatch(len(a), len(b))

		if longestPossible < len(found) {
			continue
		}

		// match is less than best so far
		if !match(a, b, len(found)) {
			continue
		}

		// match is equal to best length found.  improve?
		improved := len(found)
		for ; improved < longestPossible; improved++ {
			if a[improved] != b[improved] {
				break
			}
		}
		proposed := a[:improved]
		if allSpace(proposed) {
			continue
		}
		if len(proposed) > len(found) || maxLen > earliest {
			found = proposed
			earliest = maxLen
		}

	}
	if len(found) == 0 {
		return "NONE"
	}
	return string(found)
}

func match(a, b []byte, n int) bool {
	for k := n - 1; k >= 0; k-- {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

func allSpace(s []byte) bool {
	for _, c := range s {
		if c != ' ' {
			return false
		}
	}
	return true
}

// longestPossibleMatch is length of shortest suffix,
// or difference between lengths (which is the longest non-overlapping section),
// whichever is smaller.
func longestPossibleMatch(a, b int) (int, int) {
	if a > b {
		return min(b, a-b), a
	}
	return min(a, b-a), b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func suffixArray(word []byte) [][]byte {
	a := make([][]byte, len(word))
	for j := 0; j < len(word); j++ {
		a[j] = word[j:]
	}
	return a
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
		fmt.Println(longestRepeatedSubstring(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
