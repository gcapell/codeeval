package main

import (
	"bufio"

	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type pos struct{ x, y int }
type edge struct {
	d2 int
	p  pos
}
type edgeHeap []edge

func (eh edgeHeap) Len() int            { return len(eh) }
func (eh edgeHeap) Less(i, j int) bool  { return eh[i].d2 < eh[j].d2 }
func (eh edgeHeap) Swap(i, j int)       { eh[i], eh[j] = eh[j], eh[i] }
func (eh *edgeHeap) Push(x interface{}) { *eh = append(*eh, x.(edge)) }

func (eh *edgeHeap) Pop() interface{} {
	old := *eh
	n := len(old)
	e := old[n-1]
	*eh = old[0 : n-1]
	return e
}

func update(distances *edgeHeap, points map[pos]bool, src pos) {
	delete(points, src)
	for p := range points {
		heap.Push(distances, distance(src, p))
	}
}

func minSpanLength(line string) int {
	src, points := mapPos(strings.Fields(line))

	var distances edgeHeap
	update(&distances, points, src)

	var total float64
	for len(points) > 0 {
		var next edge
		for {
			next = heap.Pop(&distances).(edge)
			if points[next.p] {
				break
			}
		}
		total += math.Sqrt(float64(next.d2))
		update(&distances, points, next.p)
	}
	return int(math.Ceil(total))
}

func distance(src, dst pos) edge {
	dx, dy := src.x-dst.x, src.y-dst.y
	return edge{dx*dx + dy*dy, dst}
}

func atoi(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return n
}

func mapPos(chunks []string) (pos, map[pos]bool) {
	reply := make(map[pos]bool)
	var p pos
	for _, chunk := range chunks {
		ns := strings.Split(chunk, ",")
		p = pos{atoi(ns[0]), atoi(ns[1])}
		if !reply[p] {
			reply[p] = true
		}
	}
	return p, reply
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
		fmt.Println(minSpanLength(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
