package main

import (
	"fmt"
	"time"
)

func say3(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, (i+1)*100)
	}
}
func say2(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(s, (i+1)*150)
	}
}

func say4(s string, ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(s, (i+1)*150)
	}
	ch <- 0
	close(ch)
}

func main() {
	//我们单独写一个 say2 函数来跑 goroutine，并且 Sleep 时间设置长一点，150 毫秒，看看会发生什么：

	go say2("world")
	say3("hello")

	//问题来了，say2 只执行了 3 次，而不是设想的 5 次，为什么呢？
	//原来，在 goroutine 还没来得及跑完 5 次的时候，主函数已经退出了。
	//我们要想办法阻止主函数的结束，要等待 goroutine 执行完成之后，再退出主函数：
	//我们引入一个信道，默认的，信道的存消息和取消息都是阻塞的，在 goroutine 中执行完成后给信道一个值 0，则主函数会一直等待信道中的值，一旦信道有值，主函数才会结束。
	ch := make(chan int)
	go say4("world", ch)
	say2("hello")
	fmt.Println(<-ch)

}
