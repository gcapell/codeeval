package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type event struct {
	t     time.Time
	start bool
}

type byEvent []event

func (a byEvent) Len() int      { return len(a) }
func (a byEvent) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byEvent) Less(i, j int) bool {
	if a[i].t.Before(a[j].t) {
		return true
	}
	if a[i].t.After(a[j].t) {
		return false
	}
	return a[i].start
}

func merge(periods []string) int {
	events := make([]event, 0, len(periods)*2)
	const layout = "Jan 2006"
	for _, p := range periods {
		half := strings.Split(p, "-")
		start, err0 := time.Parse(layout, strings.TrimSpace(half[0]))
		end, err1 := time.Parse(layout, strings.TrimSpace(half[1]))

		if err0 != nil || err1 != nil {
			log.Fatalf("%q -> %s, %q->%s", half[0], err0, half[1], err1)
		}
		events = append(events, event{start, true}, event{end, false})
	}
	sort.Sort(byEvent(events))

	months := 0
	inPeriod := false
	overlapping := 0
	var start time.Time

	for _, e := range events {
		if e.start {
			if !inPeriod {
				inPeriod = true
				start = e.t
			}
			overlapping++
			continue
		}
		// finish event
		overlapping--
		if overlapping < 0 {
			panic("negative overlap")
		}
		if overlapping == 0 {
			sy, sm, _ := start.Date()
			ey, em, _ := e.t.Date()
			months += 12*(ey-sy) + int(em) - int(sm) + 1
			inPeriod = false
		}
	}
	if overlapping != 0 {
		panic("remaining overlap")
	}
	return months
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
		periods := strings.Split(scanner.Text(), ";")
		fmt.Println(merge(periods) / 12)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
