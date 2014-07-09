package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	for line := range linesFromFilename() {
		fields := strings.Split(line, ";")
		fmt.Println(lcs(fields[0], fields[1]))
	}
}

func lcs(a, b string) string {
	dp := make([][]int, len(a)+1)
	for j := 0; j <= len(a); j++ {
		dp[j] = make([]int, len(b)+1)
	}

	for j := 0; j < len(a); j++ {
		for k := 0; k < len(b); k++ {
			n := 0
			if a[j] == b[k] {
				n = dp[j][k] + 1
			} else {
				n = max(dp[j][k+1], dp[j+1][k])
			}
			dp[j+1][k+1] = n
		}
	}

	lcsLength := dp[len(a)][len(b)]
	reply := make([]byte, lcsLength)

	r, c := len(a), len(b)
	pos := lcsLength
	for pos > 0 {
		switch dp[r][c] {
		case dp[r-1][c]:
			r--
		case dp[r][c-1]:
			c--
		default:
			r--
			c--
			pos--
			reply[pos] = a[r]
		}
	}

	return string(reply)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
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
