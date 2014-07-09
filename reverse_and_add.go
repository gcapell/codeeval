package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	for line := range linesFromFilename() {
		iters, pal := reverseAndAdd(int64(atoi(line)))
		fmt.Println(iters, pal)
	}
}

func reverseAndAdd(n int64) (int, int64) {
	orig := n
	for j :=0; j<100; j++ {
		r := reverse(n)
		// fmt.Println(r, n)
		if r == n {
			return j, n
		}
		n += r
	}
	log.Fatal("not found", n, orig)
	return 0,0
}

func reverse(n int64) int64 {
	var r,d int64
	for n>0 {
		d, n = n%10, n/10
		r = r*10 +d
	}
	return r
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
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
