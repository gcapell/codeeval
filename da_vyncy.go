package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"bytes"
)

type suffix struct {
	s    string
	id   int
	dead bool
}

type suffixes []*suffix

func (a suffixes) Len() int      { return len(a) }
func (a suffixes) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a suffixes) Less(i, j int) bool {
	switch {
	case a[i].dead:
		return false
	case a[j].dead:
		return true
	}
	return a[i].s < a[j].s
}

func (s suffix) String() string {
	if s.dead {
		return fmt.Sprintf("X(%d)", s.id)
	}
	return fmt.Sprintf("%q(%d-%d)", s.s, s.id, len(s.s))
}

func (ss suffixes)String() string{
	var b bytes.Buffer
	for _, s := range ss {
		if s.dead {
			break
		}
		fmt.Fprintf(&b, "%s\n", s)
	}
	return b.String()
}

var (
	fragmentList suffixes
	suffixList   suffixes
	fragmentMap  = make(map[int]*suffix)
	suffixMap    = make(map[int][]*suffix)
)

func assemble(s string) string {
	for id, c := range strings.Split(s, ";") {
		f := &suffix{s: c, id: id}
		fragmentList = append(fragmentList, f)
		fragmentMap[id] = f
		for off := 0; off < len(c); off++ {
			s := &suffix{s: c[off:], id: id}
			suffixList = append(suffixList, s)
			suffixMap[id] = append(suffixMap[id], s)
		}
	}

	var merged *suffix
	merges := len(fragmentList) - 1
	for j := 0; j < merges; j++ {
		sort.Sort(fragmentList)
		sort.Sort(suffixList)
		log.Println("fragments:\n", fragmentList)
		// log.Println("suffixes", summary(suffixList))
		f, s := bestMatch()
		f2, s2 := dumbMatch()
		if len(s.s) != len(s2.s) {
			log.Fatalf("smart: %s,%s\ndumb:%s,%s", f, s, f2, s2)
		}
		merged = merge(s, f)
	}
	return merged.s
}

func bestMatch() (bestF, bestS *suffix) {
	maxOverlap := 0

	fp, sp := 0, 0
	for fp < len(fragmentList) && sp < len(suffixList) {
		f := fragmentList[fp]
		s := suffixList[sp]
		
		if f.id == s.id {
			sp++
			continue
		}		
		if f.dead || s.dead {
			return
		}
		if len(s.s) > maxOverlap && strings.HasPrefix(f.s, s.s) {
			maxOverlap, bestF, bestS = len(s.s), f, s
		}
		if s.s < f.s { sp++ } else { fp++ }
	}
	return
}

//  suffix: [xyz]COMMON, fragment: COMMONabc
// Afterwards: Left: xyzCOMMONabc, right: dead
func merge(suffix, fragment *suffix) *suffix{
	suffixFragment := fragmentMap[suffix.id]
	overlap := len(suffix.s)
	trailingUncommon := fragment.s[overlap:]
	suffixFragment.s += trailingUncommon
	fragment.dead = true
	
	add := len(suffixFragment.s) - len(suffix.s)
	log.Printf(`merge "%s%s"(%d), "%s%s"(%d)`, 
		suffixFragment.s[:add], strings.ToUpper(suffix.s), suffix.id, 
		strings.ToUpper(fragment.s[:overlap]), fragment.s[overlap:], fragment.id)
		
	if !strings.HasPrefix(fragment.s, suffix.s) {
		log.Fatal("hoo boy")
	}

	for _, s := range suffixMap[suffix.id] {
		s.s += trailingUncommon
	}
	for _, s := range suffixMap[fragment.id] {
		if len(s.s) > len(trailingUncommon) {
			s.dead = true
			continue
		}
		s.id = suffix.id
		suffixMap[suffix.id] = append(suffixMap[suffix.id], s)
	}
	return suffixFragment
}

func dumbMatch() (bestF, bestS *suffix) {
	maxOverlap := 0
	for _, f := range fragmentList {
		if f.dead {
			return
		}
		for _, g := range fragmentList {
			if g.dead {
				break
			}
			if f== g {
				continue
			}
			o := dumbOverlap(f.s, g.s)
			if o > maxOverlap {
				bestF, bestS, maxOverlap = g, &suffix{s:f.s[len(f.s)-o:] , id:f.id}, o
			}
		}
	}
	return
}

func dumbOverlap(a,b string) int {
	var overlap int
	max := len(b)
	if max>len(a){
		max = len(a)
	}
	for o:=1; o<=max; o++ {
		if b[:o] == a[len(a)-o:] {
			overlap = o
		}
	}
	return overlap
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
		fmt.Println(assemble(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
