package main

import "fmt"

// func counter() func() {
// 	count := 0

// 	fmt.Println("カウントを初期化しました")

// 	return func() {
// 		count++
// 		fmt.Println(count)
// 	}
// }

func freeCounter(start int) func(int) int {
	count := start

	fmt.Printf("フリーカウンターを%dからはじめます\n", count)

	return func(add int) int {
		fmt.Printf("%dをたして", add)
		count += add
		return count
	}
}

func main() {
	countFunc := freeCounter(10)
	fmt.Println(countFunc(2))
	fmt.Println(countFunc(5))
	fmt.Println(countFunc(7))
}
