package practice

import "fmt"

//类型相同多个变量, 非全局变量
var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2 // 和 python 很像,不需要显示声明类型，自动推断
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main() {

	g, h := 123, "hello" // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误
	println(x, y, a, b, c, d, e, f, g, h)

	//是全局变量是允许声明但不使用。 同一类型的多个变量可以声明在同一行，如：

	var a1, b1, c1 int
	//多变量可以在同一行进行赋值，如：

	var d1 string
	a1, b1, c1, d1 = 5, 7, 9, "abc"
	//上面这行假设了变量 a，b 和 c 都已经被声明，否则的话应该这样使用：
	//a, b, c := 5, 7, "abc"

	println(a1, b1, c1, d1)
	/*右边的这些值以相同的顺序赋值给左边的变量，所以 a 的值是 5， b 的值是 7，c 的值是 "abc"。
	这被称为 并行 或 同时 赋值。
	如果你想要交换两个变量的值，则可以简单地使用 a, b = b, a，两个变量的类型必须是相同。
	空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。
	_ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。
	并行赋值也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：val, err = Func1(var1)。*/

	_, numb, strs := numbers() //只获取函数返回值的后两个
	fmt.Println(numb, strs)    //????????为什么在第一行打印，不知道为什么,使用控制台是没有问题的，这个应该是Goland IDE的问题

}

//一个可以返回多个值的函数
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}