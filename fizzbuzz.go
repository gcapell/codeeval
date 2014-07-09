package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"google3/base/go/log"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expected 'fizzbuzz {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		fields := strings.Fields(line)

		if len(fields) != 3 {
			log.Fatal("expected 3 fields, got %#v", fields)
		}
		i, err := mapInt(fields)
		if err != nil {
			log.Fatal(err)
		}
		fizzbuzz(i[0], i[1], i[2])
	}
}

func mapInt(ss []string) ([]int, error) {
	reply := make([]int, len(ss))
	for j, s := range ss {
		var err error
		reply[j], err = strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
	}
	return reply, nil
}

func fizzbuzz(a, b, n int) {
	reply := make([]string, 0)
	for j := 1; j <= n; j++ {
		var s string
		switch {
		case j%a == 0 && j%b == 0:
			s = "FB"
		case j%a == 0:
			s = "F"
		case j%b == 0:
			s = "B"
		default:
			s = strconv.Itoa(j)
		}
		reply = append(reply, s)
	}
	fmt.Println(strings.Join(reply, " "))
}
