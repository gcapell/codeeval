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

type ip uint32

func (k ip) String() string {
	return fmt.Sprintf("%x", uint32(k))
}

var (
	peers      [][]int
	nodeParent [][]int
	foundPaths [][]int
)

type ByDict [][]int

func (a ByDict) Len() int      { return len(a) }
func (a ByDict) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDict) Less(i, j int) bool {
	for c := range a[i] {
		if a[i][c] < a[j][c] {
			return true
		}
		if a[i][c] > a[j][c] {
			return false
		}
	}
	return false
}

const failMsg = "No connection"

func path(src, dst int) {
	if dst >= len(peers) {
		fmt.Println(failMsg)
		return
	}
	// BFS
	q := newQueue(len(peers) * 2)
	seen := map[int]bool{src: true}
	thisGen := map[int]bool{src: true}
	oldGen := map[int]bool{src: true}
	genMarker := -1
	q.push(src)
	q.push(genMarker)
	nodeParent = make([][]int, len(peers))

	for {
		n := q.pop()
		if n == genMarker {
			if q.empty() || len(nodeParent[dst]) != 0 {
				break
			}
			q.push(genMarker)
			for k := range thisGen {
				oldGen[k] = true
			}
			thisGen = make(map[int]bool)
			continue
		}
		for _, peer := range peers[n] {
			if !oldGen[peer] {
				nodeParent[peer] = append(nodeParent[peer], n)
			}
			if peer == dst {
				break
			}
			if !seen[peer] {
				thisGen[peer] = true
				seen[peer] = true
				q.push(peer)
			}
		}
	}

	if len(nodeParent[dst]) == 0 {
		fmt.Println(failMsg)
		return
	}
	
	buf := make([]int, len(peers))
	foundPaths = make([][]int, 0)
	backpaths(dst, src, buf, 0)
	sort.Sort(ByDict(foundPaths))

	for n, path := range foundPaths {
		if n != 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("[")
		for k, node := range path {
			if k != 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%d", node)
		}
		fmt.Printf("]")
	}
	fmt.Println()
}

func reversedCopy(a []int) []int {
	b := make([]int, len(a))
	for i, n := range a {
		b[len(a)-i-1] = n
	}
	return b
}

func backpaths(src, dst int, path []int, depth int) {
	path[depth] = src
	depth++
	if src == dst {
		foundPaths = append(foundPaths, reversedCopy(path[:depth]))
		return
	}
	for _, n := range nodeParent[src] {
		backpaths(n, dst, path, depth)
	}
}

type queue struct {
	q          []int
	head, tail int
}

func (q queue) String() string {
	return fmt.Sprintf("%v###%v", q.q[:q.head], q.q[q.head:q.tail])
}

func newQueue(size int) queue {
	return queue{make([]int, size), 0, 0}
}

func (q *queue) push(c int) {
	q.q[q.tail] = c
	q.tail++
	if q.tail == len(q.q) {
		log.Fatal("queue full", len(q.q), *q)
	}
}
func (q *queue) pop() int {
	if q.empty() {
		log.Fatal("empty pop")
	}
	c := q.q[q.head]
	q.head++
	return c
}

func (q *queue) empty() bool {
	return q.tail == q.head
}

func parseNet(line string) {
	line = strings.Trim(line, " {}")
	parts := strings.Split(line, "]")
	netToNode := make(map[ip][]int)
	nodeToNet := make(map[int][]ip)

	nodeCount := 0
	for _, p := range parts {
		subParts := strings.Split(p, ":")
		if len(subParts) < 2 {
			continue
		}
		id := atoi(strings.Trim(subParts[0], ", "))
		nodeCount++
		nets := strings.Split(strings.Trim(subParts[1], " ["), ",")
		for _, n := range nets {
			if len(n) == 0 {
				continue
			}
			key := parseNetString(strings.Trim(n, " '"))
			netToNode[key] = append(netToNode[key], id)
			nodeToNet[id] = append(nodeToNet[id], key)
		}
	}
	peers = make([][]int, nodeCount)
	for id, nets := range nodeToNet {
		seen := map[int]bool{id: true}
		for _, net := range nets {
			for _, other := range netToNode[net] {
				if seen[other] {
					continue
				}
				seen[other] = true
				peers[id] = append(peers[id], other)
			}
		}
	}
}

func parseNetString(s string) ip {
	parts := strings.Split(s, "/")

	var addr uint32
	for _, addrPart := range strings.Split(strings.TrimSpace(parts[0]), ".") {
		addr = (addr << 8) + uint32(atoi(addrPart))
	}

	mask := uint32(0xffffffff) << uint32(32-atoi(parts[1]))

	return ip(addr & mask)
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
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
	first := true
	for scanner.Scan() {
		if first {
			parseNet(scanner.Text())
			first = false
			continue
		}
		chunks := strings.Fields(scanner.Text())
		path(atoi(chunks[0]), atoi(chunks[1]))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
