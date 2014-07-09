package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"log"
)

func main() {
	data := fileFromFilename()
	lines := strings.Split(string(data), "\n")

	n, err := strconv.Atoi(strings.TrimSpace(lines[0]))
	if err != nil {
		log.Fatal(err)
	}

	sort.Sort(ByLen(lines[1:]))
	for j := 1; j <= n; j++ {
		fmt.Println(lines[j])
	}
}

type ByLen []string

func (a ByLen) Len() int           { return len(a) }
func (a ByLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLen) Less(i, j int) bool { return len(a[i]) > len(a[j]) }

func fileFromFilename() []byte {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)

	}
	return data
}
