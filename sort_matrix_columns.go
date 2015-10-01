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

func mapInts(rows []string) [][]int {
	var reply [][]int
	for _, row := range rows {
		reply = append(reply, mapInt(row))
	}
	return reply
}

func mapInt(s string) []int {
	var reply []int
	for _, f := range strings.Fields(s) {
		n, err := strconv.Atoi(f)
		if err != nil {
			log.Fatal(err)
		}
		reply = append(reply, n)
	}
	return reply
}

func sortCols(line string) string {
	rows := mapInts(strings.Split(line, "|"))
	sort.Sort(ByCol(rows))

	var outRows []string
	for _, row := range rows {
		var outN []string
		for _, n := range row {
			outN = append(outN, strconv.Itoa(n))
		}
		outRows = append(outRows, strings.Join(outN, " "))
	}
	return strings.Join(outRows, " | ")
}

type ByCol [][]int

func (a ByCol) Len() int { return len(a[0]) }
func (a ByCol) Swap(i, j int) {
	for r := 0; r < len(a); r++ {
		a[r][i], a[r][j] = a[r][j], a[r][i]
	}
}
func (a ByCol) Less(i, j int) bool {
	for r := 0; r < len(a); r++ {
		iv, jv := a[r][i], a[r][j]
		switch {
		case iv < jv:
			return true
		case jv < iv:
			return false
		}
	}
	return false
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
		fmt.Println(sortCols(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
