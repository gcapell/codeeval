package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func stupid(line string) string {
	chunks := strings.Split(line, "|")
	ns := mapAtoi(strings.Fields(chunks[0]))
	iters, err := strconv.Atoi(strings.TrimSpace(chunks[1]))
	if err != nil {
		log.Fatal(err)
	}
	for j := 0; j < iters; j++ {
		sort(ns)
	}
	return separated(ns)
}

func sort(n []int) {
	for j := 0; j+1 < len(n); j++ {
		if n[j] > n[j+1] {
			n[j], n[j+1] = n[j+1], n[j]
			return
		}
	}
}

func mapAtoi(nss []string) []int {
	var reply []int
	for _, ns := range nss {
		n, err := strconv.Atoi(ns)
		if err != nil {
			log.Fatal(err)
		}
		reply = append(reply, n)
	}
	return reply
}

func separated(ns []int) string {
	var b bytes.Buffer
	for pos, n := range ns {
		if pos != 0 {
			fmt.Fprintf(&b, " ")
		}
		fmt.Fprintf(&b, "%d", n)
	}
	return b.String()
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
		fmt.Println(stupid(scanner.Text()))

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
