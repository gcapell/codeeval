package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"net"
)

type ip uint32
func (k ip) String() string {
	return fmt.Sprintf("%x", uint32(k))
}

func parseNet(line string) int {
	line = strings.Trim(line, " {}")
	parts := strings.Split(line, "]")
	
	netToNode := make (map[ip][]int)
	nodeToNet := make(map[int][]ip)
	for _, p := range parts {
		subParts := strings.Split(p, ":")
		if len(subParts)<2 {
			continue
		}
		id := atoi(strings.Trim(subParts[0], ", "))
		nets := strings.Split(strings.Trim(subParts[1], " ["), ",")
		for _, n := range nets {
			ipBytes := parseNetString(strings.Trim(n, " '"))
			key := ip(uint32(ipBytes[0])<<24 | uint32(ipBytes[1]) <<16 | uint32(ipBytes[2])<<8 | uint32(ipBytes[3]))
			netToNode[key] = append(netToNode[key], id)
			nodeToNet[id] = append(nodeToNet[id], key)
		}
	}
	fmt.Println("netToNode", netToNode)
	for _, nodes := range netToNode {
		for _, n := range nodes {
			fmt.Printf("%d --", n)
		}
		fmt.Println()
	}
	fmt.Println("nodeToNet", nodeToNet)
	return 0
}

func parseNetString(s string) net.IP{
	parts := strings.Split(s, "/")
	ip := net.ParseIP(parts[0])
	if ip == nil {
		log.Fatalf("parse(%s)->nil", parts[0])
	}
	return ip.Mask(net.CIDRMask(atoi(parts[1]), 32))
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
		fmt.Println(parseNet(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
