package main

import (
	"fmt"
)

// Here's a function that will take an arbitrary number
// of `int`s as arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func swap(x, y int) (int, int) {
	return y, x
}
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	x := 1
	y := 2

	x, y = swap(x, y)
	fmt.Println(x, y)

	sum(1, 2, 3)
	// If you already have multiple args in a slice,
	// apply them to a variadic function using
	// `func(slice...)` like this.
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(nextInt())

}
