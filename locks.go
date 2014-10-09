package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func smart(doors, iterations int) int {
	switch iterations {
	case 0:
		return doors
	case 1:
		return doors - 1
	}
	iterations--
	cycles, extra := doors/6, doors%6
	unlocked := cycles * 3
	even := iterations%2 == 0
	if even {
		unlocked += cycles
	}

	if extra >= 1 {
		unlocked++
	}
	if extra >= 3 && even {
		unlocked++
	}
	if extra == 5 {
		unlocked++
	}

	// last iteration
	switch extra {
	case 0, 1, 5:
		unlocked--
	case 2, 4:
		unlocked++
	case 3:
		if even {
			unlocked--
		} else {
			unlocked++
		}
	}

	return unlocked
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
		var doors, iterations int
		fmt.Sscanf(scanner.Text(), "%d %d", &doors, &iterations)
		fmt.Println(smart(doors, iterations))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
