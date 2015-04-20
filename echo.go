package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expected 'prog {filename}'")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(os.Stdout, f); err != nil {
		log.Fatal(err, "in Copy")
	}
}
