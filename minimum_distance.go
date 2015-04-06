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

func minimumDistance(line string) int {
	chunks := strings.Fields(line)
	nums := make([]int,0, len(chunks))
	for _, c := range chunks[1:] {
		nums = append(nums, atoi(c))
	}
	sort.Ints(nums)
	d := 0
	for l, r := 0, len(nums)-1; l<r; l,r = l+1, r-1 {
		d += nums[r] - nums[l]
	}
	return d
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
		fmt.Println(minimumDistance(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
