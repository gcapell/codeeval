package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	message = "012222 1114142503 0313012513 03141418192102 0113 2419182119021713 06131715070119"
	key     = "BHISOECRTMGWYVALUZDNFJKPQX"
	lookup  = map[int]byte{}
)

func main() {
	fmt.Println(strings.Join(mapf(decrypt, strings.Fields(message)), " "))
}

func init() {
	for pos, val := range key {
		lookup[int(val-'A')] = byte('A' + pos)
	}
}

func mapf(f func(string) string, input []string) (reply []string) {
	for _, s := range input {
		reply = append(reply, f(s))
	}
	return reply
}

func decrypt(cipher string) string {
	reply := ""
	for len(cipher) > 0 {
		n, _ := strconv.Atoi(cipher[:2])
		cipher = cipher[2:]
		reply += string(lookup[n])
	}
	return reply

}
