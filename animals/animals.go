package main

import (
	"fmt"
	"time"
)

type animal struct {
	name  string
	count int
}

func count(name string, duration int, c chan animal) {
	for i := 0; i < 5; i++ {
		count := i + 1
		time.Sleep(time.Duration(duration) * time.Millisecond)
		fmt.Printf("%s が %d 匹\n", name, count)
		if count == 5 {
			c <- animal{name, count}
		}
	}
}

func main() {
	c := make(chan animal)

	go count("カエル", 200, c)
	go count("アヒル", 100, c)

	x, y := <-c, <-c

	fmt.Printf("%s が 合計 %d 匹\n", x.name, x.count)
	fmt.Printf("%s が 合計 %d 匹\n", y.name, y.count)
}
