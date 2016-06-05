package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func countPivots(nums []int) int {
	switch len(nums) {
	case 0, 1:
		return 0
	case 2:
		return 1
	}
	pos := pivot(nums)
	return 1 + countPivots(nums[:pos]) + countPivots(nums[pos+1:])
}

func pivot(nums []int) int {
	p := nums[0]
	ppos := 0
	left := 0          // nums[x] <= p for all x<left
	right := len(nums) // nums[x] > p for all x >= right
	for {
		var j int
		for j = right - 1; j > ppos && nums[j] >= p; j-- {
		}
		if j == ppos {
			return ppos
		}
		nums[j], nums[ppos] = nums[ppos], nums[j]
		ppos, right = j, j+1

		for j = left; j < ppos && nums[j] <= p; j++ {
		}
		if j == ppos {
			return ppos
		}
		nums[j], nums[ppos] = nums[ppos], nums[j]
		ppos, left = j, j
	}
}

func mapNums(line string) []int {
	fields := strings.Fields(line)
	reply := make([]int, len(fields))
	for pos, f := range fields {
		var err error
		reply[pos], err = strconv.Atoi(f)
		if err != nil {
			log.Fatal(err)
		}
	}
	return reply
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
		fmt.Println(countPivots(mapNums(scanner.Text())))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
