package main

import "fmt"

func main() {
	for j := 1; j <= 12; j++ {
		for k := 1; k <= 12; k++ {
			switch k {
			case 1:
				fmt.Printf("%2d", j*k)
			default:
				fmt.Printf("%4d", j*k)
			}
		}
		fmt.Println()
	}
}
