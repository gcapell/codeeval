package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	perm   byte
	access struct{ user, file int }
	env    map[access]perm
)

var (
	perms = [][]perm{
		{7, 3, 0},
		{6, 2, 4},
		{5, 1, 5},
		{3, 7, 1},
		{6, 0, 2},
		{4, 2, 6},
	}
	bit = map[string]perm{
		"grant": 1,
		"write": 2,
		"read":  4,
	}
)

func checkPerms(line string) bool {
	env := newEnv()
	for _, c := range strings.Fields(line) {
		if !checkCase(c, env) {
			return false
		}
	}
	return true
}

func checkCase(line string, env env) bool {
	parts := strings.Split(line, "=>")
	u, f, act := id(parts[0], "user_"), id(parts[1], "file_"), parts[2]

	if env.get(u, f)&bit[act] == 0 {
		return false
	}
	if act == "grant" {
		env.add(id(parts[4], "user_"), f, bit[parts[3]])
	}
	return true
}

func (e env) get(u, f int) perm {
	if p, ok := e[access{u, f}]; ok {
		return p
	}
	return perms[u][f]
}

func (e env) add(u, f int, b perm) {
	e[access{u, f}] = e.get(u, f) | b
}

func newEnv() env {
	return make(map[access]perm)
}

func id(s, prefix string) int {
	if !strings.HasPrefix(s, prefix) {
		log.Fatal(prefix, s)
	}
	n, err := strconv.Atoi(s[len(prefix):])
	if err != nil {
		log.Fatal(err, s)
	}
	return n - 1
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
	text := map[bool]string{
		false: "False",
		true:  "True",
	}
	for scanner.Scan() {
		fmt.Println(text[checkPerms(scanner.Text())])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
