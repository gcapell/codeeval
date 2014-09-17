package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func angle(line string) string {
	a, err := strconv.ParseFloat(line, 64)
	if err != nil {
		panic(err)
	}
	deg := math.Floor(a)
	frac := a - deg
	minutes := frac * 60
	fullMinutes := math.Floor(minutes)
	fracMinutes := minutes - fullMinutes
	seconds := fracMinutes * 60
	fullSeconds := math.Floor(seconds)
	// FIXME - rounding to 60 seconds??
	return fmt.Sprintf(`%d.%02d'%02d"`, int(deg), int(fullMinutes), int(fullSeconds))
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
