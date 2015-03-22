package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func freq(s string) int {
	samples := strings.Fields(s)
	fsamples := make([]float64, len(samples))
	for pos, s := range samples {
		fsamples[pos] = atof(s)
	}
	bringToZero(fsamples)

	negative := false
	cross := 0
	for _, s := range fsamples {
		if negative {
			if s > 0 {
				cross++
				negative = false
			}
		} else {
			if s < 0 {
				negative = true
			}
		}
	}
	return cross * 10

}

func bringToZero(samples []float64) {
	size := len(samples)
	chunk := size / 40
	a1 := avg(samples[:chunk])
	a2 := avg(samples[size-chunk:])
	for pos := range samples {
		samples[pos] -= a1 + (a2-a1)*(float64(pos)/float64(size))
	}
}

func avg(ss []float64) float64 {
	var total float64
	for _, s := range ss {
		total += s
	}
	return total / float64(len(ss))
}

func atof(s string) float64 {
	n, err := strconv.ParseFloat(s, 64)
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
	for scanner.Scan() {
		fmt.Println(freq(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
