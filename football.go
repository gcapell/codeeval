package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func transpose(line string) string {
	teamToCountry := make(map[int][]int)
	countries := strings.Split(line, "|")
	for c, line := range countries {
		for _, ts := range strings.Fields(line) {
			t, err := strconv.Atoi(ts)
			if err != nil {
				log.Fatal(err)
			}
			teamToCountry[t] = append(teamToCountry[t], c)
		}
	}
	var teams []int
	for t := range teamToCountry {
		teams = append(teams, t)
	}
	sort.Ints(teams)
	var b bytes.Buffer
	for _, t := range teams {
		fmt.Fprintf(&b, "%d:", t)
		cs := teamToCountry[t]
		sort.Ints(cs)
		for pos, c := range cs {
			if pos == 0 {
				fmt.Fprintf(&b, "%d", c+1)
			} else {
				fmt.Fprintf(&b, ",%d", c+1)
			}
		}
		fmt.Fprintf(&b, "; ")
	}
	return b.String()
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
		fmt.Println(transpose(scanner.Text()))

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
