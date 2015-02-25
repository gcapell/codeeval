package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	metaRE = regexp.MustCompile(`[*?[\]]+`)
	tx     = map[rune]string{
		'*': ".*",
		'?': ".",
		']': "]",
		'[': "[",
	}
)

func translate(in string) *regexp.Regexp {
	locs := metaRE.FindAllStringIndex(in, -1)

	var translated bytes.Buffer
	translated.WriteRune('^')
	p := 0

	for _, loc := range locs {
		before := in[p:loc[0]]
		qBefore := regexp.QuoteMeta(before)

		translated.WriteString(qBefore)
		for _, r := range in[loc[0]:loc[1]] {
			translated.WriteString(tx[r])
		}
		p = loc[1]
	}
	after := in[p:len(in)]
	qAfter := regexp.QuoteMeta(after)
	translated.WriteString(qAfter)
	translated.WriteRune('$')

	s := translated.String()
	// fmt.Printf("%q -> %q\n", in, s)
	re, err := regexp.Compile(s)
	if err != nil {
		log.Fatal(err)
	}
	return re
}

func reFilter(line string) string {
	f := strings.Fields(line)
	re, names := translate(f[0]), f[1:]

	var reply []string
	for _, name := range names {
		if re.MatchString(name) {
			reply = append(reply, name)
		}
	}
	if len(reply) == 0 {
		return "-"
	}
	return strings.Join(reply, " ")
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
		fmt.Println(reFilter(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
