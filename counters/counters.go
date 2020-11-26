package main

import "fmt"

func counter() func() {
	count := 0
	return func() {
		count++
		fmt.Println(count)
	}
}

func main() {
	countFunc := counter()
	countFunc()
	countFunc()
	countFunc()
}
