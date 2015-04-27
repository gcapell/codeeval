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
	off  int
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
	return a[i].s[a[i].off:] < a[j].s[a[j].off:]
}

func (s suffix) String() string {
	if s.dead {
		return fmt.Sprintf("X(%d)", s.id)
	}
	return fmt.Sprintf("%q(%d)", s.s[s.off:], s.id)
}

// cmp compares two suffixes at a position, returns -1,0,1
func cmp(a, b *suffix, pos int) int {
	if pos+a.off >= len(a.s) {
		return -1
	}
	if pos+b.off >= len(b.s) {
		return 1
	}
	ac := a.s[pos+a.off]
	bc := b.s[pos+b.off]
	switch {
	case ac < bc:
		return -1
	case ac > bc:
		return 1
	default:
		return 0
	}
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

func summary(ss suffixes) map[int]int {
	counts := make(map[int]int)
	for _, s := range ss {
		if s.dead {
			break
		}
		counts[s.id]++
	}
	return counts
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
			s := &suffix{s: c, id: id, off: off}
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
		fid, sid, overlap := bestMatch()
		fid2, sid2, overlap2 := dumbMatch()
		if !(overlap == overlap2) {
			log.Fatalf("smart: %s,%s{%d}, dumb:%s,%s{%d}",
				fragmentMap[sid], fragmentMap[fid], overlap,
				fragmentMap[sid2], fragmentMap[fid2], overlap2)
		}
		merged = merge(bestMatch())
	}
	return merged.s
}

func dumbMatch() (fid, sid, overlap int) {
	for _, f := range fragmentList {
		if f.dead {
			continue
		}
		for _, g := range fragmentList {
			if g.dead || f== g {
				continue
			}
			o := dumbOverlap(f.s, g.s)
			if o > overlap {
				fid, sid, overlap = g.id, f.id, o
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

func merge(fid, sid, overlap int) *suffix{
	prefix := fragmentMap[fid]
	suffix := fragmentMap[sid]

	add := len(suffix.s) - overlap
	log.Printf(`merge "%s%s"(%d), "%s%s"(%d)`, 
		suffix.s[:add], strings.ToUpper(suffix.s[add:]), sid, 
		strings.ToUpper(prefix.s[:overlap]), prefix.s[overlap:],fid)
	if suffix.s[add:] != prefix.s[:overlap] {
		for _, s := range suffixMap[sid]{
			log.Println("s:",s)
		}
		log.Fatal("hoo boy")
	}
	suffix.s += prefix.s[overlap:]
	prefix.dead = true

	for _, s := range suffixMap[sid] {
		s.s = suffix.s
	}
	for _, s := range suffixMap[fid] {
		if s.off < overlap {
			s.dead = true
		}
		s.s = suffix.s
		s.id = sid
		s.off += add
	}
	suffixMap[sid] = append(suffixMap[sid], suffixMap[fid]...)
	return suffix
}

func bestMatch() (fid, sid, best int) {
	fp, sp := 0, 0
	defer func() {
		if fid == sid {
			log.Fatal("WTF???", fid, sid)
		}
	}()
outer:
	for fp < len(fragmentList) && sp < len(suffixList) {
		f := fragmentList[fp]
		s := suffixList[sp]
		
		if f.id == s.id {
			sp++
			continue
		}
		if f.id == 11 || s.id ==11 || f.id == 21 || s.id == 21 {
			log.Printf("f:%s, s:%s", f, s)
		}
		
		if f.dead || s.dead {
			return
		}
		p := 0
		for {
			switch cmp(f, s, p) {
			case -1:
				fp++
				continue outer
			case 1:
				sp++
				continue outer
			case 0:
				p++
				c := ""
				if s.off +p == len(s.s){
					c="!"
				}
				if p>20{
					log.Printf("%q{%d}%s\ts:%s\tf:%s", f.s[:p], p, c, s, f)
					log.Println(s.off, p, len(s.s))
				}
				if p > best && s.off + p == len(s.s){
					best, fid, sid = p, f.id, s.id
				}
			}
		}
	}
	return
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
