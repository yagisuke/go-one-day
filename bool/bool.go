package main

import "fmt"

func main() {
	a := true
	fmt.Printf("aの値は%t\n", a)

	b := false
	fmt.Printf("a == bの値は%t\n", a == b)
	fmt.Printf("a != bの値は%t\n", a != b)

	c := 12
	d := 34
	fmt.Printf("c < dの値は%t\n", c < d)
	fmt.Printf("-c < -dの値は%t\n", -c < -d)
}
