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

const corpus = `
Mary had a little lamb its fleece was white as snow;
And everywhere that Mary went, the lamb was sure to go.
It followed her to school one day, which was against the rule;
It made the children laugh and play, to see a lamb at school.
And so the teacher turned it out, but still it lingered near,
And waited patiently about till Mary did appear.
"Why does the lamb love Mary so?" the eager children cry; 
"Why, Mary loves the lamb, you know" the teacher did reply."
`

var (
	posList = make(map[string][]int)
	words   []string
)

func init() {
	fields := strings.Fields(corpus)
	for _, f := range fields {
		words = append(words, strings.Trim(f, `";.,?"`))
	}
	for pos, w := range words {
		posList[w] = append(posList[w], pos)
	}
}

type score struct {
	word  string
	score float32
}

func (s score) String() string {
	return fmt.Sprintf("%s,%.3f", s.word, s.score)
}

type ByScore []score

func (a ByScore) Len() int      { return len(a) }
func (a ByScore) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByScore) Less(i, j int) bool {
	if a[i].score != a[j].score {
		return a[i].score > a[j].score
	}
	return a[i].word < a[j].word
}

func predictions(pos []int) string {
	counts := make(map[string]int)
	for _, p := range pos {
		if p+1 == len(words) {
			continue
		}
		counts[words[p+1]]++
	}

	var scores []score
	for w, c := range counts {
		s := float32(c) / float32(len(pos))
		scores = append(scores, score{w, s})
	}
	sort.Sort(ByScore(scores))
	var reply []string
	for _, s := range scores {
		reply = append(reply, s.String())
	}
	return strings.Join(reply, ";")
}

func findPositions(line string) []int {
	chunks := strings.Split(line, ",")
	n, err := strconv.Atoi(chunks[0])
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Fields(chunks[1])
	if len(words) != n-1 {
		log.Fatal(words, n)
	}
	var prev, pos []int
	first := true
	for _, w := range words {
		if first {
			first = false
			pos = posList[w]
		} else {
			pos = consecutives(prev, posList[w])
		}
		prev = pos
	}
	return pos
}

// consecutives returns the elements of 'b' that
// are one greater than elements of 'a'
func consecutives(a, b []int) []int {
	var reply []int
	for j, k := 0, 0; j < len(a) && k < len(b); {
		switch {
		case a[j]+1 == b[k]:
			reply = append(reply, b[k])
			j, k = j+1, k+1
		case a[j] < b[k]:
			j++
		case a[j] >= b[k]:
			k++
		}
	}
	return reply
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
		fmt.Println(predictions(findPositions(scanner.Text())))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
