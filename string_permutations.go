package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"
)

func main() {
	for line := range linesFromFilename() {
		first := true
		letters := []byte(line)
		sort.Sort(ByByte(letters))

		for word := range permute(letters) {
			fmt.Printf("%s%s", comma(&first), word)
		}
		fmt.Println()
		runtime.GC()
	}
}

type ByByte []byte

func (a ByByte) Len() int           { return len(a) }
func (a ByByte) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByByte) Less(i, j int) bool { return a[i] < a[j] }

func permute(letters []byte) chan string {
	ch := make(chan string)
	go func() {
		if len(letters) == 1 {
			ch <- string(letters)
		} else {
			for pos, c := range letters {
				if pos != 0 && c == letters[pos-1] {
					continue // skip dups
				}
				remaining := make([]byte, 0, len(letters))
				remaining = append(remaining, letters[:pos]...)
				remaining = append(remaining, letters[pos+1:]...)

				for s := range permute(remaining) {
					ch <- string(c) + s
				}
			}
		}
		close(ch)
	}()
	return ch
}

func comma(first *bool) string {
	if !*first {
		return ","
	}
	*first = false
	return ""
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
