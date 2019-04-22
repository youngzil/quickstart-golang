package practice

import "fmt"

func main() {

	//多个panic只会捕捉最后一个：

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer func() {
		panic("three")
	}()
	defer func() {
		panic("two")
	}()
	panic("one")

	//以下捕获失败
	defer recover()
	//defer fmt.Prinntln(recover)
	defer func() {
		func() {
			recover() //无效，嵌套两层
		}()
	}()

	//以下捕获有效
	defer func() {
		recover()
	}()

}

func except() {
	recover()
}
func test() {
	defer except()
	panic("runtime error")
}

/*panic 与 recover 是 Go 的两个内置函数，这两个内置函数用于处理 Go 运行时的错误，panic 用于主动抛出错误，recover 用来捕获 panic 抛出的错误。

使用场景
一般情况下有两种情况用到：
 程序遇到无法执行下去的错误时，抛出错误，主动结束运行。
 在调试程序时，通过 panic 来打印堆栈，方便定位错误。

引发panic有两种情况，一是程序主动调用，二是程序产生运行时错误，由运行时检测并退出。
发生panic后，程序会从调用panic的函数位置或发生panic的地方立即返回，逐层向上执行函数的defer语句，然后逐层打印函数调用堆栈，直到被recover捕获或运行到最外层函数。
panic不但可以在函数正常流程中抛出，在defer逻辑里也可以再次调用panic或抛出panic。defer里面的panic能够被后续执行的defer捕获。
recover用来捕获panic，阻止panic继续向上传递。recover()和defer一起使用，但是defer只有在后面的函数体内直接被掉用才能捕获panic来终止异常，否则返回nil，异常继续向外传递。*/
