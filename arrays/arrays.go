package main

import "fmt"

func main() {
	intArray := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("1番目の要素は %d です\n", intArray[0])

	intValues := []int{1, 2, 3, 4, 5, 6}
	intValues[0] -= 100

	fmt.Printf("intValuesの要素数は %d です\n", len(intValues))
	fmt.Printf("intValuesの最初の要素は %d です\n", intValues[0])
	fmt.Printf("intValuesの最後の要素は %d です\n", intValues[len(intValues)-1])

	ints := []int{1, 2, 3, 4, 5, 6}
	index0 := ints[:1]
	index1to4 := ints[1:4]
	index4toLast := ints[4:]
	fmt.Println(index0)
	fmt.Println(index1to4)
	fmt.Println(index4toLast)
}
