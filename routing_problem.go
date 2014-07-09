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
	ch := linesFromFilename()
	netLine := <-ch
	p := parser{netLine, 0}
	net := p.net()
	for line := range ch {
		f := strings.Fields(line)
		fmt.Println(paths(net, f[0], f[1]))
	}
}

type Net map[int][]port

func (p *parser) net() Net {
	var net Net = make(map[int][]port)
	p.skip("{")
	nodes := make(map[int][]port)
	for node, ports, ok := p.node(); ok; {
		net[node] = ports
	}
	p.skip("}")
	return net
}

func (p *parser) node() (int, []port, bool) {
	n, ok := p.number()
	if !ok {
		return 0, nil, false
	}
	p.skip(":")
	nets := p.netList()
	p.skipMaybe(",")
	return n, nets, true
}

func (p *parser) netList() []port {
	p.skip("[")
	ports := make([]port)
	for p.peek() == "'" {
		ports = append(ports, p.port())
		p.skipMaybe(",")
	}
	p.skip("]")
	return ports
}

type port struct {
	net  uint32
	mask uint8
}

type parser struct {
	s   string
	pos int
}

func (p *parser) skip(string) {
}

func (p *parser) number() (int, bool) {
}

func (p *parser) skipMaybe(s string) {
}

func (p *parser) port() string {
}

func paths(n *net, src, dst string) string {
	fmt.Println("paths", n, src, dst)
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
