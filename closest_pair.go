package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	for ps := range pointsFromFilename() {
		fmt.Printf("%.4f\n", closestPair(ps))
	}
}

func closestPair(ps []point) float64 {
	sort.Sort(ByX(ps))

	firstPlausible := 0
	minDistance := ps[0].dist(ps[len(ps)-1])

	for eval := 1; eval < len(ps); eval++ {
		for j := firstPlausible; j < eval; j++ {
			if ps[eval].x-ps[j].x > minDistance {
				firstPlausible = j + 1
				//	log.Println("firstPlausible->", firstPlausible)
				continue
			}
			d := ps[j].dist(ps[eval])
			if d < minDistance {
				minDistance = d
				//	log.Println("minDistance->", d, eval, j)
			}
		}
	}

	return minDistance
}

func (p point) dist(o point) float64 {
	return math.Sqrt(math.Pow(p.x-o.x, 2) + math.Pow(p.y-o.y, 2))
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

type point struct{ x, y float64 }

type ByX []point

func (a ByX) Len() int           { return len(a) }
func (a ByX) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByX) Less(i, j int) bool { return a[i].x < a[j].x }

func newPoint(s string) point {
	ss := strings.Fields(s)
	return point{float64(atoi(ss[0])), float64(atoi(ss[1]))}
}

func pointsFromFilename() chan []point {
	lines := linesFromFilename()
	c := make(chan []point)
	go func() {
		for {
			count := atoi(<-lines)
			if count == 0 {
				break
			}
			ps := make([]point, count)
			for j := 0; j < count; j++ {
				ps[j] = newPoint(<-lines)
			}
			c <- ps
		}
		close(c)
	}()
	return c
}

func linesFromFilename() chan string {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan string)

	go func() {
		reader := bufio.NewReader(f)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					log.Fatal(err)
				}
				break
			}
			c <- strings.TrimSpace(line)
		}
		close(c)
	}()
	return c
}
