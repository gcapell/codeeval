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
		sortAndTrim(&suffixList)
		sortAndTrim(&fragmentList)
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

func sortAndTrim(ss *suffixes){
	sort.Sort(*ss)
	k := len(*ss) -1
	for (*ss)[k].dead {
		k--
	}
	*ss = ([]*suffix)(*ss)[:k+1]
}

func bestMatch() (bestF, bestS *suffix) {
	maxOverlap := 0

	for _, f := range fragmentList {
		if f.dead {
			break
		}
		if len(f.s) <= maxOverlap {
			continue
		}
		if s := bestWordMatch(f, maxOverlap); s != nil {
			bestF, bestS, maxOverlap = f, s, len(s.s)
		}
	}
	return bestF, bestS
}

// Find largest word in  (sorted) suffixList (>=n chars)
// which is a prefix of 's'.  
func bestWordMatch(f *suffix, n int) *suffix{
	
	pos := sort.Search(len(suffixList), func(p int)bool{
		return suffixList[p].s>=f.s[:n]
	})
	var found *suffix
	for ;pos < len(suffixList); pos++ {
		p := suffixList[pos]
		// do the first 'n' characters match?
		if ! strings.HasPrefix(p.s, f.s[:n]) {
			break
		}
		if p.id == f.id {
			continue
		}
		// is it a full match?
		if strings.HasPrefix(f.s[n:], p.s[n:]) {
			found = p
			n = len(p.s)
		}
	}
	if found != nil {
		log.Printf("%q (id:%d) has prefix %q (from id:%d)", f.s, f.id, found.s, found.id )
		if ! strings.HasPrefix(f.s, found.s) {
			log.Fatal(f.s, found.s)
		}
	}
	return found
}

//  suffix: [xyz]COMMON, fragment: COMMONabc
// Afterwards: Left: xyzCOMMONabc, right: dead
func merge(suffix, fragment *suffix) *suffix{
	suffixFragment := fragmentMap[suffix.id]
	log.Printf("merge(%s(%s), %s)", suffix, suffixFragment.s, fragment)
	if ! strings.HasSuffix(suffixFragment.s, suffix.s) {
		log.Fatal("expected %q to have suffix %q", suffixFragment.s, suffix.s)
	}
	overlap := len(suffix.s)
	trailingUncommon := fragment.s[overlap:]
	
	add := len(suffixFragment.s) - len(suffix.s)
	log.Printf("suffixFragment: %q(%d), suffix:%q(%d), add:%d",
		suffixFragment.s, len(suffixFragment.s),
		suffix.s, len(suffix.s), 
		add)
	log.Printf(`merge "%s%s"(%d), "%s%s"(%d)`,
		suffixFragment.s[:add], strings.ToUpper(suffix.s), suffix.id, 
		strings.ToUpper(fragment.s[:overlap]), fragment.s[overlap:], fragment.id)
		
	if !strings.HasPrefix(fragment.s, suffix.s) {
		log.Fatal("hoo boy")
	}
	suffixFragment.s += trailingUncommon
	fragment.dead = true

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
