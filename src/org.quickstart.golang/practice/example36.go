package practice

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	//Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。
	//goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
	//Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。 同一个程序中的所有 goroutine 共享同一个地址空间。

	//你会看到输出的 hello 和 world 是没有固定先后顺序。因为它们是两个 goroutine 在执行：
	go say("world") //开启 goroutine
	say("hello")    //开启一个新的 goroutine:
}

/*goroutine 是 golang 中在语言级别实现的轻量级线程，仅仅利用 go 就能立刻起一个新线程。多线程会引入线程之间的同步问题，在 golang 中可以使用 channel 作为同步的工具。
通过 channel 可以实现两个 goroutine 之间的通信。
创建一个 channel， make(chan TYPE {, NUM}) TYPE 指的是 channel 中传输的数据类型，第二个参数是可选的，指的是 channel 的容量大小。
向 channel 传入数据， CHAN <- DATA ， CHAN 指的是目的 channel 即收集数据的一方， DATA 则是要传的数据。
从 channel 读取数据， DATA := <-CHAN ，和向 channel 传入数据相反，在数据输送箭头的右侧的是 channel，形象地展现了数据从隧道流出到变量里。*/
