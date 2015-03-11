package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

type (
	trie struct {
		next map[byte]*trie
		word string
	}

	queueT struct {
		head, tail int
		q          []string
	}
)


var (
	// length -> trie of words of that length
	words = map[int]*trie {}
	queue queueT
)

func (q *queueT) reset()        {}
func (q *queueT) push(s string) {}
func (q *queueT) empty() bool   { return true }
func (q *queueT) pop() string   { return "" }

// size of social network of word.
func network(word string) int {
	queue.reset()
	queue.push(word)
	seen := map[string]bool{word: true}
	for !queue.empty() {
		word = queue.pop()
		for _, n := range neighbours(word) {
			if !seen[n] {
				seen[n] = true
				queue.push(n)
			}
		}
	}
	return len(seen)
}

// neightbours returns all the words of edit-distance one from word.
func neighbours(word string) []string {
	return concat(
		words[len(word)-1].remove(word),
		words[len(word)+1].insert(word),
		words[len(word)].replace(word))
}

func concat(a ...[]string) []string {
	return a[0] // FIXME
}

// words found by removing one letter
func (t *trie) remove(word string) []string {
	if t == nil {
		return nil
	}
	if len(word) == 1 {
		return []string{t.word}
	}

	return concat(
		t.next[word[0]].remove(word[1:]),
		t.find(word[1:]))
}

// words found by inserting one letter
func (t *trie) insert(word string) []string {
	if len(word) == 0 {
		return []string{t.word}
	}
	var reply []string
	for letter, next := range t.next {
		if letter == word[0] {
			continue
		}
		reply = append(reply, next.find(word)...)
	}
	if next, ok := t.next[word[0]]; ok {
		reply = append(reply, next.insert(word[1:])...)
	}
	return reply
}

// words found by replacing one letter
func (t *trie) replace(word string) []string {
	return nil
}

func (t *trie) contains(word string) bool {
	return false

}
func (t *trie) find(word string) []string {
	return nil
}
func (t *trie)add(word string, letters []byte) *trie {
	if t == nil {
		t = &trie {next: make(map[byte]*trie)}
	}
	if len(letters) == 0 {
		t.word = word
	} else {
		t.next[letters[0]] = t.next[letters[0]].add(word, letters[1:])
	}
	return t
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
	readingDict := false
	var searchTerms []string
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case line == "END OF INPUT":
			readingDict = true
		case readingDict:
			words[len(line)] = words[len(line)].add(line, []byte(line))
		default:
			searchTerms = append(searchTerms, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(words)
	for _, term := range searchTerms {
		fmt.Println(term, network(term))
	}
}
