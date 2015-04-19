package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"container/heap"
)

const (
	nextStop     = 7
	changeRoute  = 12
	infiniteCost = 9999999
)

type node struct {
	name string
	next   map[*node]int
	cost   int
	index  int
	onHeap bool
}

func newNode(name string) *node {
	return &node{name:name,
		next:make(map[*node]int), 
		cost:infiniteCost,
		index: -1, 
		onHeap:false}
}

func network(line string) string {
	chunks := strings.Split(line, ";")
	src, dst := parseSrcDst(chunks[0])
	idToNode = make(map[int][]*node)
	for _, c := range chunks[1:] {
		addNodes(parseRoute(c))
	}

	start, finish := newNode("start"), newNode("finish")
	start.cost = 0

	// 0-cost edges from start
	for _, n := range idToNode[src] {
		start.next[n] = 0
	}
	// 0-cost edges to finish
	for _, n := range idToNode[dst] {
		n.next[finish] = 0
	}
	if cost, ok := dijkstraCost(start, finish); ok {
		return strconv.Itoa(cost)
	} else {
		return "None"
	}
}

type PriorityQueue []*node

func (pq PriorityQueue) Len() int          { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	size := len(*pq)
	n := x.(*node)
	n.index = size
	n.onHeap = true
	*pq = append(*pq, n)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	item.onHeap = false // for safetyß
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(n *node, cost int) {
	if cost < n.cost {
		n.cost = cost
		if n.onHeap {
			heap.Fix(pq, n.index)
		} else {
			heap.Push(pq, n)
		}
	}
}

func dijkstraCost(start, finish *node) (int, bool) {
	var pq PriorityQueue

	heap.Init(&pq)
	heap.Push(&pq, start)
	for len(pq) > 0 {
		if pq[0].index != 0 {
			log.Println("problem", pq)
		}
		n := heap.Pop(&pq).(*node)
		if n == finish {
			return n.cost, true
		}
		for o, delta := range n.next {
			pq.update(o, n.cost+delta)
		}
	}
	return 0, false
}

var idToNode map[int][]*node

func addNodes(ns []int) {
	var prev, first *node
	for _, id := range ns {
		np := newNode(strconv.Itoa(id))
		for _, o := range idToNode[id] {
			o.next[np] = changeRoute
			np.next[o] = changeRoute
		}
		idToNode[id] = append(idToNode[id], np)
		if prev != nil {
			prev.next[np] = nextStop
		} else {
			first = np
		}
		prev = np
	}
	prev.next[first] = nextStop
}

var routeExp = regexp.MustCompile(`R[0-9]+=\[(.*)\]`)

func parseRoute(s string) []int {
	s = strings.TrimSpace(s)
	m := routeExp.FindStringSubmatch(s)
	if len(m) != 2 {
		log.Fatalf("addRoute %q %v", s, m)
	}
	chunks := strings.Split(m[1], ",")
	var nodes []int
	for _, c := range chunks {
		nodes = append(nodes, atoi(c))
	}
	return nodes
}

var srcDstRegexp = regexp.MustCompile(`\(([0-9]+),([0-9]+)\)`)

func parseSrcDst(s string) (int, int) {

	m := srcDstRegexp.FindStringSubmatch(s)
	if len(m) != 3 {
		log.Fatalf("parseSrcDst %q", s)
	}
	return atoi(m[1]), atoi(m[2])
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
	for scanner.Scan() {
		fmt.Println(network(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
