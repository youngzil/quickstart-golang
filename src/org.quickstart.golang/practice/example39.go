package main

import (
	"fmt"
	"time"
)

func main() {

	//更好的展示边入边出概念：

	c := make(chan int, 10)

	go fibonacci6(cap(c), c)

	for v := range c {
		fmt.Println("out:", time.Now())
		fmt.Println(v)
	}
}

func fibonacci6(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		fmt.Println("in:", time.Now())
		time.Sleep(100)
		x, y = y, x+y
	}

	close(c)
}
