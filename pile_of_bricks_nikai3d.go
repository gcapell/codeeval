package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	reader := bufio.NewReader(data)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		t := strings.Split(s, "|")
		u := strings.Split(t[1], ";")
		h, h1, h2 := make([]int, 2), make([]int, 2), make([]int, 2)
		fmt.Sscanf(t[0], "[%d,%d] [%d,%d]", &h1[0], &h1[1], &h2[0], &h2[1])
		h[0], h[1] = abs(h1[0]-h2[0]), abs(h1[1]-h2[1])
		sort.Ints(h)
		b, b1, b2 := make([]int, 3), make([]int, 3), make([]int, 3)
		var r []int
		for _, i := range u {
			var j int
			fmt.Sscanf(i, "(%d [%d,%d,%d] [%d,%d,%d])", &j, &b1[0], &b1[1], &b1[2], &b2[0], &b2[1], &b2[2])
			b[0], b[1], b[2] = abs(b1[0]-b2[0]), abs(b1[1]-b2[1]), abs(b1[2]-b2[2])
			sort.Ints(b)
			if b[0] <= h[0] && b[1] <= h[1] {
				r = append(r, j)
			}
		}
		if len(r) == 0 {
			fmt.Println("-")
		} else {
			sort.Ints(r)
			var r2 []string
			for _, i := range r {
				r2 = append(r2, fmt.Sprint(i))
			}
			fmt.Println(strings.Join(r2, ","))
		}
	}
}
