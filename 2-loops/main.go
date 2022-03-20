package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 5; i++ {
		sum += 1
	}

	fmt.Println(sum)

	// Simple while sllop
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)

	// Infinte loop
	// sum := 0
	for {
		sum++
		break // only way to escape the loop
	}
	// fmt.Println(sum) // never reached
}
