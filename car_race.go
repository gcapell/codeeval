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

type section struct {
	distance float64
	curve    int
}

var track []section

func loadTrack(line string) {
	fields := strings.Fields(line)
	var s section
	for len(fields) > 0 {
		s.distance, s.curve, fields = atof(fields[0]), atoi(fields[1]), fields[2:]
		track = append(track, s)
	}
}

type car struct {
	id     int
	max    float64
	accelT float64
	decelT float64
	lap    float64
	speed  float64
}

type byLap []*car

func (a byLap) Len() int           { return len(a) }
func (a byLap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLap) Less(i, j int) bool { return a[i].lap < a[j].lap }

var cars []*car

func loadCar(line string) {
	var c car
	fmt.Sscanf(line, "%d %f %f %f", &c.id, &c.max, &c.accelT, &c.decelT)
	c.max /= 3600
	cars = append(cars, &c)
}

func race() {
	for _, c := range cars {
		for _, s := range track {
			c.race(s)
		}
	}
}

func (c *car) race(s section) {
	accelT := c.accelT * (c.max - c.speed) / c.max
	accelD := accelT * (c.speed + c.max) / 2
	endSpeed := float64(180-s.curve) / 180 * c.max
	decelT := c.decelT * (c.max - endSpeed) / float64(c.max)
	decelD := decelT * (endSpeed + c.max) / 2
	cruiseD := s.distance - accelD - decelD
	if cruiseD < 0 {
		panic("no cruise")
	}
	cruiseT := cruiseD / c.max

	c.lap += accelT + cruiseT + decelT
	c.speed = endSpeed
}

func atof(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
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
	scanner.Scan()
	loadTrack(scanner.Text())
	for scanner.Scan() {
		loadCar(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	race()

	sort.Sort(byLap(cars))
	for _, c := range cars {
		fmt.Printf("%d %.02f\n", c.id, c.lap)
	}
}
