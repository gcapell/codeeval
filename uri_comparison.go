package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"os"
	"strings"
)

func main() {
	for line := range linesFromFilename() {
		if sameURL(splitURLs(line)) {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}

func splitURLs(line string) (string, string) {
	fields := strings.Split(line, ";")
	return fields[0], fields[1]
}

func canonicalURL(s string) *url.URL {
	u, err := url.Parse(s)

	if err != nil {
		log.Println(err)
	}
	host, port, err := net.SplitHostPort(u.Host)
	if err != nil {
		port = ""
	}
	if port == "" {
		port = "80"
	}
	u.Host = net.JoinHostPort(host, port)

	return u
}

func sameURL(a, b string) bool {
	au := canonicalURL(a)
	bu := canonicalURL(a)
	// log.Printf("%q->%#v, %q->%#v\n", a, au, b, bu)
	return au.Scheme == bu.Scheme && au.Host == bu.Host && au.Path == bu.Path

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
