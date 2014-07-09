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
		fmt.Println(decode(line))
	}
}

func decode(line string) string {
	z := strings.IndexAny(line, "01")
	header, encoded := line[:z], line[z:]

	var decoded []byte
	for c := range getCodes([]byte(encoded)) {
		decoded = append(decoded, header[c.flat()])
	}
	return string(decoded)
}

type code struct {
	length, message uint8
}

var offset = map[uint8]uint8{
	1: 0,
	2: 1,
	3: 4,
	4: 11,
	5: 26,
	6: 57,
	7: 120,
}

func (c code) flat() uint8 {
	return offset[c.length] + c.message

}
func getCodes(encoded []byte) chan code {
	ch := make(chan code)
	go func() {
		for {
			length := pop(&encoded, 3)
			if length == 0 {
				break
			}
			for {
				message := pop(&encoded, length)
				if message == (1<<length)-1 {
					break
				}
				ch <- code{length, message}
			}
		}
		close(ch)
	}()
	return ch
}

func pop(line *[]byte, length uint8) uint8 {
	chunk := string((*line)[:length])
	*line = (*line)[length:]

	n, err := strconv.ParseInt(chunk, 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	return uint8(n)
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
