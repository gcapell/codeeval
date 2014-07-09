package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const DEBUG = false

func main() {
	for line := range linesFromFilename() {
		result := fmt.Sprintf("%.5f", calculator(line))

		result = strings.TrimRight(result, "0")	// BUG! 10 -> 1 :-(
		result = strings.TrimRight(result, ".")
		fmt.Println(result)
	}
}

func calculator(line string) float64 {
	if DEBUG {
		fmt.Printf("\n%v\n", line)
	}

	var n floatStack
	var tok tokenStack

	for t := range tokens(line) {
		if DEBUG {
			fmt.Println("\n\ttoken", t)
		}

		// http://wcipeg.com/wiki/Shunting_yard_algorithm
		switch t.t {
		case numberT:
			n.push(t.n)
		case openBracket:
			tok.push(t)
		case closeBracket:
			for {
				popped := tok.pop()
				if popped.t == openBracket {
					break
				}
				popped.op.apply(&n)
			}
		case operatorT:
			for !tok.empty() {
				peek := tok.peek()
				if peek.t == openBracket {
					break
				}
				if t.op.left {
					if peek.op.precedence < t.op.precedence {
						break
					}
				} else {
					if peek.op.precedence <= t.op.precedence {
						break
					}
				}
				tok.pop().op.apply(&n)
			}
			tok.push(t)
		}
		if DEBUG {
			fmt.Printf("\t%v %v\n", n, tok)
		}
	}
	// Empty operator stack
	for !tok.empty() {
		tok.pop().op.apply(&n)
		if DEBUG {
			fmt.Printf("\t%v %v\n", n, tok)
		}
	}
	return n.pop()
}

type floatStack []float64
type tokenStack []token

func (f *floatStack) push(n float64) {
	*f = append(*f, n)
}

func (f *floatStack) pop() float64 {
	fs := []float64(*f)
	n := fs[len(fs)-1]
	*f = fs[:len(fs)-1]
	return n
}

func (t token) String() string {
	switch t.t {
	case openBracket:
		return "("
	case closeBracket:
		return ")"
	case numberT:
		return fmt.Sprintf("%v", t.n)
	case operatorT:
		return t.op.name
	}
	return "BOGUS"
}

func (t *tokenStack) push(a token) {
	*t = append(*t, a)
}

func (t *tokenStack) pop() token {
	ts := []token(*t)
	val := ts[len(ts)-1]
	*t = ts[:len(ts)-1]
	return val
}

func (t *tokenStack) empty() bool { return len(*t) == 0 }

func (t *tokenStack) peek() token { return []token(*t)[len(*t)-1] }

type tokenType int

const (
	numberT tokenType = iota
	operatorT
	openBracket
	closeBracket
)

type operator struct {
	precedence int
	left       bool
	apply      func(*floatStack)
	name       string
}

type token struct {
	t  tokenType
	n  float64
	op operator
}

func exp(f *floatStack) {
	b, a := f.pop(), f.pop()
	f.push(math.Pow(a, b))
}
func mul(f *floatStack) {
	f.push(f.pop() * f.pop())
}
func div(f *floatStack) {
	b, a := f.pop(), f.pop()
	f.push(a / b)
}
func add(f *floatStack) {
	f.push(f.pop() + f.pop())
}
func sub(f *floatStack) {
	b, a := f.pop(), f.pop()
	f.push(a - b)
}
func uminus(f *floatStack) {
	f.push(-f.pop())
}

var opTable = map[byte]operator{
	'^': operator{4, false, exp, "^"},
	'*': operator{3, true, mul, "*"},
	'/': operator{3, true, div, "/"},
	'+': operator{2, true, add, "+"},
	'-': operator{2, true, sub, "-"},
}

var uminusOp = operator{5, false, uminus, "u-"}

func tokens(line string) chan token {
	ch := make(chan token)
	unary := true
	go func() {

		for pos := 0; pos < len(line); {
			// skip whitespace
			for line[pos] == ' ' {
				pos++
			}
			c := line[pos]
			switch {
			case unary && c == '-':
				pos++
				unary = false
				ch <- token{t: operatorT, op: uminusOp}

			case ('0' <= c && c <= '9'):
				delta, n := parseN(line[pos:])
				ch <- token{t: numberT, n: n}
				unary = false
				pos += delta
			case c == '(':
				ch <- token{t: openBracket}
				unary = true
				pos++
			case c == ')':
				ch <- token{t: closeBracket}
				unary = false
				pos++
			default:
				op, ok := opTable[c]
				if !ok {
					log.Fatal("%q not in %q", string(c), op)
				}
				pos++
				unary = true
				ch <- token{t: operatorT, op: op}
			}
		}
		close(ch)
	}()
	return ch
}

func parseN(s string) (int, float64) {
	pos := 0
	for ; pos < len(s); pos++ {
		c := s[pos]
		if c == '.' || c >= '0' && c <= '9' {
			continue
		}
		break
	}

	f, err := strconv.ParseFloat(s[:pos], 64)
	if err != nil {
		log.Fatal(err)
	}
	return pos, f
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
			if len(line)>0 {
				c <- line
			}
		}
		close(c)
	}()
	return c
}
