package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	for line := range linesFromFilename() {
		fmt.Println(removeChars(line))
	}
}

func removeChars(line string) string {

	f := strings.Split(line, ",")
	if len(f) != 2 {
		log.Fatalf("expected src, pattern, got %q", f)
	}
	orig := strings.TrimSpace(f[0])
	pattern := strings.TrimSpace(f[1])

	re := regexp.MustCompile("[" + pattern + "]")
	return re.ReplaceAllString(orig, "")
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
