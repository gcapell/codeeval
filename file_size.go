package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	fi, err := os.Stat(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fi.Size())
}
