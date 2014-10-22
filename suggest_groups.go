package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func suggestGroups(data string) {
	lines := strings.Split(data, "\n")
	group2user := make(map[string]map[string]bool)
	user2friends := make(map[string][]string)

	for _, line := range lines {
		chunks := strings.Split(line, ":")
		if len(chunks) < 3 {
			continue
		}
		user, friends, groups := chunks[0], strings.Split(chunks[1], ","), strings.Split(chunks[2], ",")
		for _, g := range groups {
			if len(g) == 0 {
				continue
			}
			d, ok := group2user[g]
			if !ok {
				d = make(map[string]bool)
				group2user[g] = d
			}
			d[user] = true
		}
		user2friends[user] = append(user2friends[user], friends...)
		for _, f := range friends {
			user2friends[f] = append(user2friends[f], user)
		}
	}

	suggestions := make(map[string][]string)
	for user, friends := range user2friends {
		for g, members := range group2user {
			if !members[user] && majority(friends, members) {
				suggestions[user] = append(suggestions[user], g)
			}
		}
	}
	keys := make([]string, 0, len(suggestions))
	for u := range suggestions {
		keys = append(keys, u)
	}
	sort.Strings(keys)
	for _, k := range keys {
		sort.Strings(suggestions[k])
		fmt.Printf("%s:%s\n", k, strings.Join(suggestions[k], ","))
	}
}

func majority(friends []string, group map[string]bool) bool {
	count := 0
	for _, f := range friends {
		if group[f] {
			count++
		}
	}
	return count*2 >= len(friends)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expected filename")
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	suggestGroups(string(data))
}
