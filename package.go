package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type parcel struct {
	name   string
	label  uint64
	weight float64
	value  int
}

type ByDensity []parcel

func (a ByDensity) Len() int      { return len(a) }
func (a ByDensity) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByDensity) Less(i, j int) bool {
	return float64(a[i].value)/a[i].weight > float64(a[j].value)/a[j].weight
}

type solution struct {
	value  int
	labels uint64
}

func (s solution) labelString() string {
	var reply []string
	for j, labels := 0, s.labels; labels > 0; j, labels = j+1, labels>>1 {
		if labels&1 != 0 {
			reply = append(reply, strconv.Itoa(j))
		}
	}
	return strings.Join(reply, ",")
}

func (s solution) String() string {
	return fmt.Sprintf("%d: %s", s.value, s.labelString())
}

var sol     solution

func knapsack(line string) string {
	chunks := strings.Split(line, ":")
	capacity := atof(chunks[0])
	parcels := parsePackages(chunks[1], capacity)
	sort.Sort(ByDensity(parcels))
	sol = greedy(capacity, parcels)

	if sol.labels == 0 {
		return "-"
	}
	tryAll(0, parcels, capacity, 0)
	return sol.labelString()
}

func tryAll(value int, parcels []parcel, capacity float64, labels uint64) {
	if len(parcels) == 0 {
		if value > sol.value {
			sol.value = value
			sol.labels = labels
		}
		return
	}
	if value+maxRemainingValue(parcels, capacity) < sol.value {
		return
	}
	p := parcels[0]
	if p.weight <= capacity {
		tryAll(value+p.value, parcels[1:], capacity-p.weight, labels|p.label)
	}
	tryAll(value, parcels[1:], capacity, labels)
}

func maxRemainingValue(ps []parcel, capacity float64) int {
	var value int
	for _, p := range ps {
		if p.weight <= capacity {
			value += p.value
		}
	}
	return value
}

func greedy(capacity float64, parcels []parcel) solution {
	var reply solution
	for _, p := range parcels {
		if p.weight <= capacity {
			reply.value += p.value
			reply.labels |= p.label
			capacity -= p.weight
		}
	}
	return reply
}

var packageRE = regexp.MustCompile(`\(([0-9]+),([0-9.]+),\$([0-9]+)\)`)

func parsePackages(s string, capacity float64) []parcel {
	var parcels []parcel
	for _, m := range packageRE.FindAllStringSubmatch(s, -1) {
		weight := atof(m[2])
		if weight > capacity {
			continue
		}
		parcels = append(parcels, parcel{
			name:   m[1],
			label:  1 << uint64(atoi(m[1])),
			weight: weight,
			value:  atoi(m[3]),
		})
	}
	return parcels
}

func atof(s string) float64 {
	f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func atoi(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
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
		fmt.Println(knapsack(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
