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
	label   int
	flag    uint32
	x, y    float64
	minDist float64
}

func (p *point) updateMinDist(d float64) {
	if p.minDist > d {
		p.minDist = d
	}
}

var distances []float64

func (p *point)String()string {
	return fmt.Sprintf("%d: %f", p.label, p.minDist)
}

func (p *point)distance(q)float64 {
	return 
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

var bestDistance float64

func tsp(points []*point) []int {
	distances = make([]float64, 2<<uint32(len(points)))
	minDistances(points)
	fmt.Println(points)
	fmt.Println(distances)
	
	// initialize bestDistance to path 1,2,3,...
	p:= points[0]
	for _, q := range points[1:]{
		bestDistance += distances[p.flag|q.flag]
		p = q
	}
	
	return nil
}

func minDistances(points []*point) {
	for j, p := range points {
		for _, q := range points[j+1:] {
			dx, dy := p.x-q.x, p.y-q.y
			d := math.Sqrt(dx*dx+dy*dy)
			p.updateMinDist(d)
			q.updateMinDist(d)
			n := p.flag | q.flag
			distances[p.flag|q.flag] = d
		}
	}
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
	var points []*point
	for scanner.Scan() {
		points = append(points, parsePoint(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for _, o := range tsp(points) {
		fmt.Println(o + 1)
	}
}
