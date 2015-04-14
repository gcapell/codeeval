package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type addrPair struct{ a, b string } // a<b

var links = make(map[addrPair]uint)
var linkCount int

func addLine(s string) {
	chunks := strings.Split(s, "\t")
	src, dst := chunks[1], chunks[2]
	var flag uint = 1
	if src > dst {
		src, dst = dst, src
		flag = 2
	}
	links[addrPair{src, dst}] |= flag
	linkCount++
}

type node struct {
	neighbours map[*node]struct{}
	id         uint64
}

func (n *node) String() string {
	return fmt.Sprintf("<%d>", len(n.neighbours))
}

var nodes = make(map[string]*node)

func getNode(name string) *node {
	if n, ok := nodes[name]; ok {
		return n
	}
	n := &node{make(map[*node]struct{}), 0}
	nodes[name] = n
	return n
}

func makeNodes() {
	var empty struct{}
	for k, v := range links {
		if v != 3 {
			continue
		}
		n1, n2 := getNode(k.a), getNode(k.b)
		n1.neighbours[n2] = empty
		n2.neighbours[n1] = empty
	}
	links = nil
}

// remove nodes of degree < 2 (they are never part of a clique of size 3)
func trimNodes() {
	for name, n := range nodes {
		if removeLowDegree(n) {
			delete(nodes, name)
		}
	}
}

func removeLowDegree(n *node) bool {
	if len(n.neighbours) >= 2 {
		return false
	}
	for o := range n.neighbours {
		delete(n.neighbours, o)
		delete(o.neighbours, n)
		removeLowDegree(o)
	}
	return true
}

type bitset uint64
type idNode struct {
	name       string
	id         uint64
	neighbours bitset
}

func (n *idNode) String() string {
	return fmt.Sprintf("%x", n.neighbours)
}

var nodesById []*idNode

func idNodes() {
	var id uint64
	nodesById = make([]*idNode, len(nodes))
	for name, old := range nodes {
		nodesById[int(id)] = &idNode{name: name, id: id}
		old.id = id
		id++
	}
	for _, old := range nodes {
		n := nodesById[int(old.id)]
		for o := range old.neighbours {
			n.neighbours |= 1 << o.id
		}
	}
}

// sample [40040040 82100400 81080 1001802010 1001802008 6420010000 40040001 81004 908400800 8200004000  82100002 908400100 80084 1001800018 8200000200 414220000 6420000020 414208000 40000041 1084 82000402 414028000 908000900 1001002018 1000802018 80100402 410228000 900400900 404228000 6400010020 40041 2100402 80840090033 8000004200 6034238020 108400900 1802018 4420010020 2420010020 200004200]

var cliques []bitset

func bronKerbosch(r, p, x bitset) {
	if p == 0 && x == 0 {
		if len(elems(r)) >= 3 {
			cliques = append(cliques, r)
		}
	}
	for _, v := range elems(p) {
		neighbours := nodesById[v].neighbours
		bit := bitset(1 << uint(v))
		bronKerbosch(r|bit, p&neighbours, x&neighbours)
		p &= ^bit
		x |= bit
	}
}

func elems(b bitset) []int {
	var reply []int
	for n := 0; b > 0; n, b = n+1, b>>1 {
		if b&1 != 0 {
			reply = append(reply, n)
		}
	}
	return reply
}

func report() {
	makeNodes()
	trimNodes()
	idNodes()
	nodes = nil

	bronKerbosch(0, all(len(nodesById)), 0)
	for _, s := range cliqueStrings() {
		fmt.Println(s)
	}
}

func cliqueStrings() []string {
	reply := make([]string, 0, len(cliques))
	for _, c := range cliques {
		ids := elems(c)
		addrs := make([]string, 0, len(ids))
		for _, id := range ids {
			addrs = append(addrs, nodesById[id].name)
		}
		sort.Strings(addrs)
		reply = append(reply, strings.Join(addrs, ", "))
	}
	sort.Strings(reply)
	return reply
}

func all(n int) bitset {
	return (1 << uint(n)) - 1
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
		addLine(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	report()
}
