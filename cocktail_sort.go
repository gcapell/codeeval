package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func simCocktail(line string) []int {
	chunks := strings.Split(line, " | ")
	iters := atoi(chunks[1])
	nums := mapAtoi(strings.Fields(chunks[0]))
	for j := 0; j < iters; j++ {
		cocktail(nums[j : len(nums)-j])
	}
	return nums
}

func cocktail(nums []int) {
	for j := 0; j < len(nums)-1; j++ {
		sort(nums, j)
	}
	for j := len(nums) - 2; j >= 0; j-- {
		sort(nums, j)
	}
}

func sort(nums []int, j int) {
	if nums[j] > nums[j+1] {
		nums[j], nums[j+1] = nums[j+1], nums[j]
	}
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func mapAtoi(ss []string) []int {
	reply := make([]int, len(ss))
	for n, s := range ss {
		reply[n] = atoi(s)
	}
	return reply
}

func printNums(ns []int) {
	for pos, n := range ns {
		if pos != 0 {
			fmt.Print(" ")
		}
		fmt.Print(n)
	}
	fmt.Println()
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
		printNums(simCocktail(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
