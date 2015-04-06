package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func numberOperations(line string) bool {
	chunks := strings.Fields(line)
	ns := make([]int, len(chunks))
	for j, chunk := range chunks {
		var err error
		if ns[j], err = strconv.Atoi(chunk); err != nil {
			log.Fatal(err)
		}
	}
	for j := range ns {
		swap(ns, j)
		if extend(ns[0], ns[1:]) {
			return true
		}
		swap(ns, j)
	}
	return false
}

func extend(soFar int, ns []int) bool {
	if len(ns) == 0 {
		return soFar == 42
	}
	for j := range ns {
		swap(ns, j)
		if extend(soFar*ns[0], ns[1:]) ||
			extend(soFar+ns[0], ns[1:]) ||
			extend(soFar-ns[0], ns[1:]) {
			return true
		}
		swap(ns, j)
	}
	return false
}

func swap(ns []int, j int) {
	if j != 0 {
		ns[0], ns[j] = ns[j], ns[0]
	}
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
		reply := "NO"
		if numberOperations(scanner.Text()) {
			reply = "YES"
		}
		fmt.Println(reply)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
