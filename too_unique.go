package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type rect struct{ x, y, xr, yr int }

func replaceUnique(data string) {
	lines := strings.Split(data, "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	max := 0
	var found []rect
	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			for xr := x + 1; xr <= len(lines[0]); xr++ {
				for yr := y + 1; yr <= len(lines); yr++ {
					if !unique(lines, x, y, xr, yr) {
						break
					}
					area := (xr - x) * (yr - y)
					r := rect{x, y, xr, yr}
					switch {
					case area > max:
						found = []rect{r}
						max = area
					case area == max:
						found = append(found, r)
					}
				}
			}
		}
	}
	for l, line := range lines {
		fmt.Println(replace(line, xranges(found, l)))
	}
}

func replace(line string, rs []xrange) string {
	b := []byte(line)
	for _, r := range rs {
		for j := r.l; j < r.r; j++ {
			b[j] = '*'
		}
	}
	return string(b)
}

type xrange struct{ l, r int }

func xranges(found []rect, y int) []xrange {
	var reply []xrange
	for _, r := range found {
		if y >= r.y && y < r.yr {
			reply = append(reply, xrange{r.x, r.xr})
		}
	}
	return reply
}

func unique(lines []string, x, y0, xr, yr int) bool {
	var seen [26]bool
	for ; x < xr; x++ {
		for y := y0; y < yr; y++ {
			c := lines[y][x] - 'a'
			if seen[c] {
				return false
			}
			seen[c] = true
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expected filename")
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	replaceUnique(string(data))
}
