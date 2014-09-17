package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)



func age(line string) string {
	n, err := strconv.Atoi(line)
	if err != nil {panic(err)}
	switch  {
		case n<0 || n > 100: return "This program is for humans"
		case n<=2: return "Still in Mama's arms"
		case n<=4: return "Preschool Maniac"
		case n<=11: return "Elementary school"
		case n<=14: return "Middle school"
		case n <= 18: return "High school"
		case n<=22: return "College"
		case n<=65: return "Working for the man"
		default: return "The Golden Years"
	}
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
		fmt.Println(age(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
