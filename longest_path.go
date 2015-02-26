package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

var (
	size, maxDistance int
	board             string
	lettersUsed       [26]bool
	squaresUsed       [36]bool
)

func extend(p, distance int) {
	if squaresUsed[p] || lettersUsed[board[p]-'a'] {
		return
	}
	squaresUsed[p] = true
	lettersUsed[board[p]-'a'] = true
	distance++
	if distance > maxDistance {
		maxDistance = distance
	}
	var buf [4]int
	for _, p2 := range neighbours(p, buf[:0]) {
		extend(p2, distance)
	}
	squaresUsed[p] = false
	lettersUsed[board[p]-'a'] = false
}

func longest(line string) int {
	board = line
	size = int(math.Sqrt(float64(len(line))))
	maxDistance = 0

	for p, _ := range board {
		extend(p, 0)
	}
	return maxDistance
}

func neighbours(p int, reply []int) []int {
	switch p % size {
	case 0:
		reply = append(reply, p+1)
	case size - 1:
		reply = append(reply, p-1)
	default:
		reply = append(reply, p-1)
		reply = append(reply, p+1)
	}
	switch p / size {
	case 0:
		reply = append(reply, p+size)
	case size - 1:
		reply = append(reply, p-size)
	default:
		reply = append(reply, p-size)
		reply = append(reply, p+size)
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

		fmt.Println(longest(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
