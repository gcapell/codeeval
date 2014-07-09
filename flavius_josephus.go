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

func main() {
	for line := range linesFromFilename() {
		s := fmt.Sprintf("%v", flavius(parse(line)))
		fmt.Println(s[1 : len(s)-1])
	}
}

func parse(line string) (int, int) {
	n := atoi(strings.Split(line, ","))
	return n[0], n[1]
}

func flavius(people, jump int) []int {
	dead := make([]bool, people)
	deaths := make([]int, people)

	pos := 0
	for j := 0; j < people; j++ {

		k := 0
		// Skip until pos points to jump'th live
		for {
			if !dead[pos] {
				k++
				if k == jump {
					break
				}
			}
			pos = (pos + 1) % people
		}
		dead[pos] = true
		deaths[j] = pos
	}
	return deaths
}

func atoi(ss []string) []int {
	ns := make([]int, len(ss))

	for pos, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		ns[pos] = n
	}
	return ns
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
