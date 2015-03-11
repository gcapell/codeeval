package main

type (
	trie struct {
		next dict[byte]*trie
		word string
	}
	
	queue struct {
		head, tail int
		q []string
	}
)

// length -> trie of words of that length
var (
	words dict[int]trie
	queue queue
)

// size of social network of word.
func network(word string) {
	queue.reset()
	q.push(word)
	seen := dict[string]bool {word:true}
	while !q.empty() {
		word = q.pop()
		for _, n := neighbours(word) {
			if !seen[n] {
				seen[n] = true
				q.push(n)
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

// words found by removing one letter
func (t *trie) remove (word string) []string {
	if t == nil {return nil}
	if len(word)==1 { 	return []string{t.word}}
	
	return concat(
		t.next[word[0]].remove(word[1:]),
		t.find(word[1:]))
}

// words found by inserting one letter
func (t trie) insert (word string) []string {
	if len(word) == 0 {
		return []string{t.word}
	}
	var reply []string
	for letter, next := range t.next {
		if letter == word[0] {
			continue
		}
		reply = append(reply, next.find(word))
	}
	if next, ok := t.next[word[0]], ok {
		reply = append(reply, next.insert(word[1:]))
	}
	return reply
}

// words found by replacing one letter
func (t trie) replace (word string) []string {
}

func (t trie) contains (word string) bool {
	
}

func main() {
	
}