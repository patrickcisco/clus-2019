package main

import "fmt"

// In Go, variables are explicitly declared and used by the compiler
var (
	isReptile bool
	isMammal  = true
)

// GOAT is a constant
// Go supports constants of character, string, boolean, and numeric values.
const GOAT = "goat"

func main() {
	var dog string
	cat := "cat"
	dog = "dog"
	sentence := fmt.Sprintf("%vs and %vs have how many legs? %v", cat, dog, 4)
	fmt.Println(sentence)
}
