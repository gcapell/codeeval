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
		board(line)
	}
}

var isMine = map[uint8]bool{
	'.': false,
	'*': true,
}

func board(line string) {
	var rows, cols int
	var desc string
	_, err := fmt.Sscanf(line, "%d,%d;%s", &rows, &cols, &desc)
	if err != nil {
		log.Fatalf("%q, %q\n", err, line)
	}

	// Read
	a := make([][]bool, rows)
	for r := 0; r < rows; r++ {
		row := make([]bool, cols)
		for c := 0; c < cols; c++ {
			row[c] = isMine[desc[r*cols+c]]
		}
		a[r] = row
	}

	mineCount := func(baseR, baseC int) string {
		count := 0
		delta := []int{-1, 0, 1}
		for _, dr := range delta {
			r := baseR + dr
			if r < 0 || r == rows {
				continue
			}
			for _, dc := range delta {
				c := baseC + dc
				if c < 0 || c == cols {
					continue
				}
				if dr == 0 && dc == 0 {
					continue
				}
				if a[r][c] {
					count++
				}
			}
		}
		return fmt.Sprintf("%d", count)
	}

	// Calc/display
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if a[r][c] {
				fmt.Printf("*")
			} else {
				fmt.Printf(mineCount(r, c))
			}
		}
	}
	fmt.Println()
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
