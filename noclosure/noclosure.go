package main

import "fmt"

func addCount(count *int) {
	*count++
}

func freeAddCount(count *int, add int) {
	*count += add
}

func main() {
	counter := 0
	freeCounter := 10

	for i := 0; i < 5; i++ {
		fmt.Printf("Before addCount: %d\n", counter)
		addCount(&counter)
		fmt.Printf("After addCount: %d\n", counter)

		fmt.Printf("Before freeAddCount: %d\n", counter)
		freeAddCount(&counter, freeCounter)
		fmt.Printf("After freeAddCount: %d\n", counter)
	}
}
