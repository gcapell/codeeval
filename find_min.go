package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	for line := range linesFromFilename() {
		n, k, a, b, c, r := parse(line)
		fmt.Println(find_min(n, k, prng(a, b, c, r, k)))
	}
}

func find_min(n, k int, r chan int) int {
	counts := make([]int, k+1)
	ring := make([]int, k+1)

	// Generate random numbers (and count the useful ones)
	for j := 0; j <= k; j++ {
		next := <-r
		ring[j] = next
		if next <= k {
			counts[next]++
		}
	}

	// Store initially missing numbers
	missing := &IntHeap{}
	heap.Init(missing)
	for n, count := range counts {
		if count == 0 {
			heap.Push(missing, n)
		}
	}

	n -= k

	// first cycle
	for j := 0; j <= k; j++ {
		next := heap.Pop(missing)
		remove := ring[j]
		if remove < k {
			counts[remove]--
			if counts[remove] == 0 {
				heap.Push(missing, remove)
			}
		}
		ring[j] = next.(int)
		n--
		if n == 0 {
			return ring[j]
		}
	}

	return ring[(n-1)%(k+1)]
}

func parse(line string) (n, k, a, b, c, r int) {
	fields := strings.Split(line, ",")
	f := func(j int) int {
		n, err := strconv.Atoi(fields[j])
		if err != nil {
			log.Fatal(err)
		}
		return n
	}
	return f(0), f(1), f(2), f(3), f(4), f(5)
}

func prng(a, b, c, r, k int) chan int {
	ch := make(chan int)
	go func() {
		ch <- a
		val := a
		for j := 0; j<k; j++{
			val = (b*val + c) % r
			ch <- val
		}
		close(ch)
	}()
	return ch
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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
			c <- strings.TrimSpace(line)
		}
		close(c)
	}()
	return c
}
