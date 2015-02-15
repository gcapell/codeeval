package main

import (
	"github.com/gcapell/codeeval/wrap"
)

func readMore(line string) string {
	return line
}

func main() {
	wrap.LineToString(readMore)
}
