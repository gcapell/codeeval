package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func candies(line string) int {
	var vampires, zombies, witches, houses int

	fmt.Sscanf(line, "Vampires: %d, Zombies: %d, Witches: %d, Houses: %d",
		&vampires, &zombies, &witches, &houses)

	candy := (vampires*3 + zombies*4 + witches*5) * houses
	children := vampires + zombies + witches

	return int(math.Floor(float64(candy) / float64(children)))
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
		fmt.Println(candies(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
