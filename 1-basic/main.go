package main

import "fmt"

func main() {
	// Bariables can be declared with var
	var b string
	b = "hello"
	// "" empty string
	// 0
	// false

	a := "hello" // understand the type from value from : , means string
	a = "world"

	if a == "world" {
		fmt.Println("a is hello")
	} else {
		fmt.Println("a is not hello")
	}

	if b == "" {
		fmt.Println("Primitive zero values are always not nil")
	}
}
