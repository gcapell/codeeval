package main

import "fmt"

func main() {
	fmt.Println(visits(0, 1))
}

var neighbours = [][]uint8{
	{1, 4},
	{0, 5, 2},
	{1, 3, 6},
	{2, 7},
	{0, 5, 8},
	{1, 4, 9, 6},
	{2, 5, 7, 10},
	{3, 6, 11},
	{4, 9, 12},
	{8, 5, 10, 13},
	{9, 6, 11, 14},
	{10, 7, 15},
	{8, 13},
	{9, 12, 14},
	{10, 13, 15},
	{14, 11},
}

func visits(p uint8, visited uint32) int {
	if p == 15 {
		return 1
	}
	total := 0
	for _, o := range neighbours[p] {
		if visited&(1<<o) != 0 {
			continue
		}
		total += visits(o, visited|(1<<o))
	}
	return total
}
