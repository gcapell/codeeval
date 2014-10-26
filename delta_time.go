package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func delta(line string) string {
	f := strings.Fields(line)
	t1, t2 := parse(f[0]), parse(f[1])
	return t1.diff(t2).String()
}

type time uint

func (t time) String() string {
	m, s := t/60, t%60
	h, m := m/60, m%60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func (t time) diff(u time) time {
	if t > u {
		return t - u
	}
	return u - t
}

func parse(line string) time {
	var h, m, s time
	fmt.Sscanf(line, "%d:%d:%d", &h, &m, &s)
	return h*60*60 + m*60 + s
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expected filename")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(delta(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
