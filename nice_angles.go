package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)



func angle(line string) string {
	a, err := strconv.Atof(line)
	if err != nil {panic(err)}
	deg := math.Floor(a)
	frac := a - deg
	minutes := frac * 60
	fullMinutes := math.Floor(minutes)
	fracMinutes := minutes - fullMinutes
	seconds := fracMinutes * 60
	
	
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
		fmt.Println(angle(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
