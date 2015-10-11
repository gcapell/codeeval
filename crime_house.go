package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func minCrim(line string) (int, bool) {
	return sim(parse(line))
}

type event struct {
	enter, masked bool
	who           int
}

func (e event) String() string {
	if e.enter {
		if e.masked {
			return fmt.Sprintf("E")
		} else {
			return fmt.Sprintf("e%d", e.who)
		}
	} else {
		if e.masked {
			return fmt.Sprintf("L")
		} else {
			return fmt.Sprintf("l%d", e.who)
		}
	}
}

func dprint(a ...interface{}) {
	//fmt.Println(a...)
}

func sim(es []event) (int, bool) {
	dprint(es)

	in := make(map[int]bool)
	out := make(map[int]bool)
	inRandom := 0

	for pos, e := range es {
		dprint(e)
		if e.enter {
			if e.masked {
				if n, ok := nextIn(es[pos+1:], false, out); ok {
					// next leaving who is already out
					move(in, out, n)
				} else if n, ok := nextNotIn(es[pos+1:], false, in); ok {
					// next leaving who is not already in
					move(in, out, n)
				} else {
					inRandom++
				}
			} else {
				if in[e.who] {
					return 0, false
				}
				move(in, out, e.who)
			}
		} else {
			if e.masked {
				if n, ok := nextIn(es[pos+1:], true, in); ok {
					// next entering who is already in
					move(out, in, n)
				} else if n, ok := nextNotIn(es[pos+1:], true, out); ok {
					// next entering not already out
					move(out, in, n)
				} else if inRandom > 0 {
					inRandom--
				}
			} else {
				if out[e.who] {
					return 0, false
				}
				move(out, in, e.who)
			}
		}
		dprint("in", in, "out:", out, inRandom)
	}
	return len(in) + inRandom, true
}

func move(to, from map[int]bool, n int) {
	if to[n] {
		log.Fatal("move", to, n)
	}
	to[n] = true
	delete(from, n)
}

func nextIn(es []event, enter bool, in map[int]bool) (int, bool) {
	reversed := make(map[int]bool)
	for _, e := range es {
		if e.masked || !in[e.who] || reversed[e.who] {
			continue
		}

		if e.enter != enter {
			reversed[e.who] = true
			if len(reversed) == len(in) {
				break
			}
			continue
		}
		return e.who, true
	}
	return 0, false
}

// nextNotIn returns the first person in es whose next event is 'enter'
// and who is not in 'in'.
func nextNotIn(es []event, enter bool, in map[int]bool) (int, bool) {
	reversed := make(map[int]bool)
	for _, e := range es {
		if e.masked || in[e.who] || reversed[e.who] {
			continue
		}
		if e.enter != enter {
			reversed[e.who] = true
			continue
		}
		return e.who, true
	}
	return 0, false
}

func parse(line string) []event {
	var events []event
	parts := strings.Split(line, ";")

	for _, es := range strings.Split(parts[1], "|") {
		var what string
		var who int
		fmt.Sscanf(es, "%s %d", &what, &who)
		events = append(events, parseEvent(what, who))
	}
	return events
}

func parseEvent(what string, who int) event {
	return event {
		enter:what=="E",
		masked:who==0,
		who:who,
	}
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
		if n, ok := minCrim(scanner.Text()); ok {
			fmt.Println(n)
		} else {
			fmt.Println("CRIME TIME")
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
