package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	nextStop     = 7
	changeRoute  = 12
	infiniteCost = 9999999
	debug        = false
)

type node struct {
	name   string
	next   map[*node]int
	cost   int
	index  int
	onHeap bool
}

func newNode(name string) *node {
	return &node{name: name,
		next:   make(map[*node]int),
		cost:   infiniteCost,
		index:  -1,
		onHeap: false}
}

func network(line string) string {
	chunks := strings.Split(line, ";")
	var src, dst int
	if _, err := fmt.Sscanf(chunks[0], "(%d,%d)", &src, &dst); err != nil {
		log.Fatal(err)
	}
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
	dot(start)
	if cost, ok := dijkstraCost(start, finish); ok {
		return strconv.Itoa(cost)
	} else {
		return "None"
	}
}

func dprint(args ...interface{}) {
	if debug {
		fmt.Println(args...)
	}
}

func dot(start *node) {
	if !debug {
		return
	}
	fmt.Println("digraph G {")
	for o, cost := range start.next {
		fmt.Printf("%s -> %s [ label = \"%d\" ];\n", start.name, o.name, cost)
	}

	for _, ns := range idToNode {
		for _, n := range ns {
			if len(n.next) == 0 {
				log.Fatal("no links from:", *n)
			}
			for o, cost := range n.next {
				fmt.Printf("%s -> %s [ label = \"%d\" ];\n", n.name, o.name, cost)
			}
		}
	}
	fmt.Println("}")
}

type PriorityQueue []*node

func (pq PriorityQueue) Len() int { return len(pq) }

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
	item.index = -1     // for safety
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
		n.onHeap = false
		dprint("popped", n.name, n.cost)
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

func addNodes(route int, ns []int) {
	var prev *node
	for _, id := range ns {
		np := newNode(fmt.Sprintf("%d.%d", route, id))
		for _, o := range idToNode[id] {
			o.next[np] = changeRoute
			np.next[o] = changeRoute
		}
		idToNode[id] = append(idToNode[id], np)
		if prev != nil {
			prev.next[np] = nextStop
			np.next[prev] = nextStop
		}
		prev = np
	}
}

var routeExp = regexp.MustCompile(`R([0-9]+)=\[(.*)\]`)

// parseRoute('R4=[1,2,3]') -> 4,[1,2,3]
func parseRoute(s string) (int, []int) {
	s = strings.TrimSpace(s)
	m := routeExp.FindStringSubmatch(s)
	if len(m) != 3 {
		log.Fatalf("addRoute %q %v", s, m)
	}
	chunks := strings.Split(m[2], ",")
	var nodes []int
	for _, c := range chunks {
		nodes = append(nodes, atoi(c))
	}
	return atoi(m[1]), nodes
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
