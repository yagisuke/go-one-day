package main

import "fmt"

func main() {
	scores := [][3]int{
		{10000, 20000, 30000},
		{1000, 2000, 3000},
		{100, 200, 300},
		{10, 20, 30},
	}

	for i := 0; i < len(scores); i++ {
		sum := 0
		for k := 0; k < 3; k++ {
			sum += scores[i][k]
		}
		fmt.Printf("受験者%d: 平均値%d点\n", i+1, sum/3)
	}
}
