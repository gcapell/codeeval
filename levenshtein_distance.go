package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expected filename")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	readingDict := false
	var searchTerms []string
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case line == "END OF INPUT":
			readingDict = true
		case readingDict:
			addWord(line)
		default:
			searchTerms = append(searchTerms, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for _, term := range searchTerms {
		seen = make(map[string]bool)
		network(term)
		fmt.Println(len(seen))
	}
}

func addWord(s string) {
	n := len(s)
	words[n] = words[n].addWord(s, []byte(s))
}

var seen map[string]bool

// traverse social network of word.
func network(word string) {
	seen[word] = true
	for _, n := range neighbours(word) {
		if !seen[n] {
			network(n)
		}
	}
}

// neightbours returns all the words of edit-distance one from word.
func neighbours(word string) []string {
	var reply []string
	if t := words[len(word)-1]; t != nil {
		reply = append(reply, t.remove(word)...)
	}
	if t := words[len(word)]; t != nil {
		reply = append(reply, t.replace(word)...)
	}
	if t := words[len(word)+1]; t != nil {
		reply = append(reply, t.add(word)...)
	}
	return reply
}

type trie struct {
	next map[byte]*trie
	word string
}

// length -> trie of words of that length
var words = map[int]*trie{}

// words found by removing one letter
func (t *trie) remove(word string) []string {
	if len(word) == 1 {
		return []string{t.word}
	}
	reply := t.find(word[1:])
	if next := t.next[word[0]]; next != nil {
		reply = append(reply, next.remove(word[1:])...)
	}
	return reply
}

// words found by adding one letter
func (t *trie) add(word string) []string {
	if len(word) == 0 {
		var reply []string
		for _, next := range t.next {
			reply = append(reply, next.word)
		}
		return reply
	}
	var reply []string
	for letter, next := range t.next {
		var found []string
		if letter == word[0] {
			found = next.add(word[1:])
		} else {
			found = next.find(word)
		}
		reply = append(reply, found...)
	}
	return reply
}

// words found by replacing one letter
func (t *trie) replace(word string) []string {
	if len(word) == 0 {
		return nil
	}
	var reply []string

	for letter, next := range t.next {
		var found []string
		if letter == word[0] {
			found = next.replace(word[1:])
		} else {
			found = next.find(word[1:])

		}
		reply = append(reply, found...)
	}
	return reply
}

func (t *trie) find(word string) []string {
	for _, r := range word {
		if t = t.next[byte(r)]; t == nil {
			return nil
		}
	}
	return []string{t.word}
}

func (t *trie) addWord(word string, letters []byte) *trie {
	if t == nil {
		t = &trie{next: make(map[byte]*trie)}
	}
	if len(letters) == 0 {
		t.word = word
	} else {
		t.next[letters[0]] = t.next[letters[0]].addWord(word, letters[1:])
	}
	return t
}
