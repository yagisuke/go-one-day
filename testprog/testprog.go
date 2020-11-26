package main

import "fmt"

func assert(fn func(int, int) int, a int, b int, result int) bool {
	return fn(a, b) == result
}

func main() {
	fmt.Println(assert(func(a int, b int) int {
		return a + b
	}, 5, 13, 18))

	fmt.Println(assert(func(a int, b int) int {
		return a * b * b
	}, 3, 5, 15))
}
