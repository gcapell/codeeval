package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const DEBUG = false

func main() {

	for line := range linesFromFilename() {
		output(pairs(line))
	}
}

type pair struct{ a, b int }

func pairs(line string) []pair {

	f := strings.Split(line, ";")
	target := atoi(f[1])
	nums := atoiSlice(strings.Split(f[0], ","))

	low, high := 0, len(nums)-1

	for nums[high] > target {
		high--
	}

	var reply []pair
	for {
		delta := target - nums[high]
		if delta < 0 {
			log.Fatal("oops", low, high, nums)
		}

		// FIXME - use binary search
		for low+1 < len(nums) && nums[low+1] <= delta {
			low++
		}

		if low >= high {
			break
		}
		if nums[low] == delta {
			reply = append(reply, pair{nums[low], nums[high]})
		}
		high--
		if low > high {
			break
		}
	}
	return reply
}

func output(pairs []pair) {
	if len(pairs) == 0 {
		fmt.Println("NULL")
		return
	}
	for pos, p := range pairs {
		if pos == 0 {
			fmt.Printf("%d,%d", p.a, p.b)
		} else {
			fmt.Printf(";%d,%d", p.a, p.b)
		}
	}
	fmt.Println()
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func atoiSlice(ss []string) []int {
	reply := make([]int, len(ss))
	for pos, s := range ss {
		reply[pos] = atoi(s)
	}
	return reply
}

func linesFromFilename() chan string {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan string)

	go func() {
		reader := bufio.NewReader(f)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					log.Fatal(err)
				}
				break
			}
			line = strings.TrimSpace(line)
			if len(line) > 0 {
				c <- line
			}
		}
		close(c)
	}()
	return c
}
