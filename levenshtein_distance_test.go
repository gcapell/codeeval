package main

import (
	"testing"
	"strings"
	"sort"
	
)

const dictionary = `
cat
car
rat
rate
scat
cats
`

func init() {
	words := strings.Fields(dictionary)
	for _, w := range words {
		addWord(w)
	}
}

func TestAdd(t *testing.T) {
	src := "cat"
	want := []string{"scat", "cats"}
	
	got := words[4].add(src)
	
	if !sameWords(got, want) {
		t.Errorf("add(%s)->%q, want %q", src, got, want)
	}
}

func sameWords(a,b []string) bool{
	sort.Strings(a)
	sort.Strings(b)
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] !=  b[i] {
			return false
		}
	}
	return true
}

func TestRemove(t *testing.T) {
	
}

func TestReplace(t *testing.T) {
	
}

