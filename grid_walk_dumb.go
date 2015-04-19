package main

import "fmt"

var total int

type point struct {x,y int}
const boardSize=600
var board[boardSize*boardSize/64]uint64
func (p point)pos() (int, uint64) {
	x := p.x+ boardSize/2
	y := p.y + boardSize/2
	pos := x * boardSize + y
	word, bitPos := pos/64, pos%64
	//fmt.Printf("%v.point->%d,%d\n", p, word, bitPos)
	return word, 1<<uint64(bitPos)
}

func (p point)visited() bool {
	word, mask := p.pos()
	return board[word]&mask != 0
}

var visited int

func (p point)visit() {
	word, mask := p.pos()
	board[word] |= mask
	visited++
}

func main() {
	fillFrom(point{0,0})
	fmt.Println(visited)
}

func fillFrom(p point) {
	if p.visited() {
		return
	}
	if sumDigits(p.x) + sumDigits(p.y) >19 {
		return
	}
	p.visit()
	fillFrom(point{p.x+1, p.y})
	fillFrom(point{p.x-1, p.y})
	fillFrom(point{p.x, p.y+1})
	fillFrom(point{p.x, p.y-1})
}

var sumCache [300] int
func sumDigits(n int) int {
	if n<0 {
		n=-n
	}
	if sumCache[n] > 0 {
		return sumCache[n]
	}
	sum := 0
	for j := n; j > 0; j /= 10 {
		sum += j % 10
	}
	sumCache[n] = sum
	return sum
}

