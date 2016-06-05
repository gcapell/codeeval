package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct{ x, y int }

func contained(line string) bool {
	chunks := strings.Split(line, "|")
	var prison []point
	for _, ch := range strings.Split(chunks[0], ", ") {
		prison = append(prison, parseCoord(ch))
	}
	return contains(prison, parseCoord(chunks[1]))
}

func contains(prison []point, person point) bool {
	c := 0
	for j := 0; j < len(prison); j++ {
		on, cut := intersect(prison[j], prison[(j+1)%len(prison)], person)
		if on {
			return true
		}
		if cut {
			c++
		}
	}
	return c%2 == 1
}

// intersect says if vertical ray from -Inf to c intersects segment (a,b)
// We return (on, cut) where `on` indicates if c itself intersects (a,b),
// `cut` indicates if vertical ray up to c intersects.
func intersect(a, b, c point) (bool, bool) {
	// a and b both above c -> nope
	if a.y > c.y && b.y > c.y {
		return false, false
	}
	// ensure a is to left of b
	if a.x > b.x {
		a, b = b, a
	}
	// a and b both on left or right of c -> nope
	if !(a.x <= b.x && b.x >= c.x) {
		return false, false
	}
	if a.y <= c.y && b.y <= c.y {
		return a.y == c.y && b.y == c.y, true
	}
	// a is left-ish, b is right-ish,
	// a/b are on opposite vertical sides

}

func parseCoord(s string) point {
	fs := strings.Fields(s)
	return point{atoi(fs[0]), atoi(fs[1])}
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
	status := map[bool]string{
		true:  "Prisoner",
		false: "Citizen",
	}
	for scanner.Scan() {
		fmt.Println(status[contained(scanner.Text())])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
