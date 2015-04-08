package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type point struct {
	label     int
	flag      uint32
	lat, long float64
	minDist   float64
}

// distance calculates great circle distance using haversine formula
// from http://www.movable-type.co.uk/scripts/latlong.html
func (p *point) distance(q *point) float64 {
	φ1 := radians(p.lat)
	φ2 := radians(q.lat)
	Δφ := φ1 - φ2
	Δλ := radians(p.long - q.long)
	a := sqr(math.Sin(Δφ/2)) + math.Cos(φ1)*math.Cos(φ2)* sqr(math.Sin(Δλ/2))

	return 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

func sqr(a float64) float64 {return a * a}
func radians(degrees float64) float64 { return degrees / 180 * math.Pi}
func swap(ps []*point, j int) {ps[0], ps[j] = ps[j], ps[0]}


var (
	bestDistance float64
	bestPath     []int
	points       []*point
)

func travellingSalesman() {
	distances = make([]float64, 2<<uint32(len(points)))
	bestPath = make([]int, len(points))

	minDistances(points)

	// initialize bestDistance to path 1,2,3,...
	p := points[0]
	for _, q := range points[1:] {
		bestDistance += distances[p.flag|q.flag]
		p = q
	}
	storePath()

	findPaths(points[0], float64(0), points[1:])
}

func findPaths(p *point, soFar float64, other []*point) {
	if len(other) == 0 {
		if soFar < bestDistance {
			bestDistance = soFar
			storePath()
		}
		return
	}
	if soFar+totalMinDist(other) > bestDistance {
		return
	}
	for j := range other {
		swap(other, j)
		q := other[0]
		findPaths(q, soFar+distances[p.flag|q.flag], other[1:])
		swap(other, j)
	}
}

func totalMinDist(ps []*point) float64 {
	var reply float64
	for _, p := range ps {
		reply += p.minDist
	}
	return reply
}

func storePath() {
	for k := range points {
		bestPath[k] = points[k].label
	}
}

var distances []float64

func minDistances(points []*point) {
	for j, p := range points {
		for _, q := range points[j+1:] {
			d := p.distance(q)
			p.updateMinDist(d)
			q.updateMinDist(d)
			distances[p.flag|q.flag] = d
		}
	}
}

func (p *point) updateMinDist(d float64) {
	if p.minDist > d {
		p.minDist = d
	}
}

func atof(s string) float64 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

var pointRE = regexp.MustCompile(`([0-9]+) \| .*\(([-0-9.]+), ([-0-9.]+)\)`)

func parsePoint(s string) *point {
	m := pointRE.FindStringSubmatch(s)
	label := atoi(m[1])

	return &point{label, 1 << uint32(label), atof(m[2]), atof(m[3]), 9999999}
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
		points = append(points, parsePoint(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	travellingSalesman()
	for _, o := range bestPath {
		fmt.Println(o)
	}
}
