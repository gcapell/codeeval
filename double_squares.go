package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const DEBUG = false

func main() {

	first := true
	for line := range linesFromFilename() {
		if first {
			first = false
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(pairs(int64(n)))
	}
}

func pairs(n int64) int {

	low := int64(0)
	high := int64(math.Sqrt(float64(n)))

	count := 0
	for {
		delta := n - high*high
		if delta < 0 {
			log.Fatal("oops", low, high, n, high*high)
		}

		// FIXME - use binary search
		for (low+1)*(low+1) <= delta {
			low++
		}

		if low > high {
			break
		}
		if low*low == delta {
			if DEBUG {
				fmt.Printf("\t%d %d\n", high, low)
			}
			count++
		}
		high--
		if low > high {
			break
		}
	}
	return count
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
