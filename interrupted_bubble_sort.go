package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func interrupted(line string) string {
	chunks := strings.Split(line, "|")
	nums, iterations := mapInt(chunks[0]), mapInt(chunks[1])[0]
	for j := 0; j < iterations; j++ {
		if !bubble(nums) {
			break
		}
	}
	reply := fmt.Sprint(nums)
	return reply[1 : len(reply)-1]
}

func bubble(n []int) bool {
	changed := false
	for j := 0; j < len(n)-1; j++ {
		if n[j] > n[j+1] {
			n[j], n[j+1] = n[j+1], n[j]
			changed = true
		}
	}
	return changed
}

func mapInt(s string) (reply []int) {
	for _, f := range strings.Fields(s) {
		n, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}
		reply = append(reply, n)
	}
	return
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
		fmt.Println(interrupted(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
