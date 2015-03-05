package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

type ip uint32

func (k ip) String() string {
	return fmt.Sprintf("%x", uint32(k))
}

var (
	peers      [][]int
	nodeParent []int
)

func path(src, dst int) string {
	// BFS
	fmt.Println("path", src, dst)
	q := newQueue(len(peers)*2)
	seen := map[int]bool{src:true}
	genMarker := -1
	q.push(src)
	q.push(genMarker)
	var penultimates []int
	for !q.empty() {
		n := q.pop()
		fmt.Println(q)
		if n == genMarker {
			q.push(genMarker)
			if len(penultimates) != 0 {
				break
			}
			continue
		}
		for _, peer := range peers[n] {
			if seen[peer] {
				continue
			}
			seen[peer] = true
			nodeParent[peer] = n
			if peer == dst {
				penultimates = append(penultimates, n)
				break
			}
			q.push(peer)
		}
	}
	if len(penultimates) == 0 {
		return "No connection"
	}
	fmt.Println(penultimates)
	return ""
	
	var reply []int
	for n := dst; ; n = nodeParent[n] {
		reply = append(reply, n)
		if n == src {
			break
		}
	}
	// reverse
	for i, j := 0, len(reply)-1; i < j; i, j = i+1, j-1 {
		reply[i], reply[j] = reply[j], reply[i]
	}
	fmt.Println(src, dst, reply)
	return ""
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

	for _, p := range parts {
		subParts := strings.Split(p, ":")
		if len(subParts) < 2 {
			continue
		}
		id := atoi(strings.Trim(subParts[0], ", "))
		nets := strings.Split(strings.Trim(subParts[1], " ["), ",")
		for _, n := range nets {
			key := parseNetString(strings.Trim(n, " '"))
			netToNode[key] = append(netToNode[key], id)
			nodeToNet[id] = append(nodeToNet[id], key)
		}
	}
	peers = make([][]int, len(nodeToNet))
	nodeParent = make([]int, len(nodeToNet))
	for id, nets := range nodeToNet {
		seen := map[int]bool{id:true}
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
	fmt.Println("peers", peers)
}

func parseNetString(s string) ip {
	parts := strings.Split(s, "/")
	addr := net.ParseIP(parts[0])
	if addr == nil {
		log.Fatalf("parse(%s)->nil", parts[0])
	}
	ipBytes := addr.Mask(net.CIDRMask(atoi(parts[1]), 32))
	return ip(uint32(ipBytes[0])<<24 | uint32(ipBytes[1])<<16 | uint32(ipBytes[2])<<8 | uint32(ipBytes[3]))
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
	first := true
	for scanner.Scan() {
		if first {
			parseNet(scanner.Text())
			first = false
			continue
		}
		chunks := strings.Fields(scanner.Text())
		fmt.Println(path(atoi(chunks[0]), atoi(chunks[1])))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
