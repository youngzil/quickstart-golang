// You can edit this code!
// Click here and start typing.
package practice

import "fmt"

func main() {
	/* 这是我的示例程序 */

	var age int = 20
	var bl bool = true
	fmt.Println("age=", age)
	fmt.Println("error=", bl)
	fmt.Printf("hello, world\n")

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	//以下几种类型为 nil：

	var a1 *int
	var a2 []int
	var a3 map[string]int
	var a4 chan int
	var a5 func(string) int
	var a6 error // error 是接口

	fmt.Println("test= ", a1, a2, a3, a4, a5, a6)

	//第二种，根据值自行判定变量类型。
	//第三种，省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：
	//var intVal int
	//intVal :=1 // 这时候会产生编译错误

	intVal, intVal1 := 1, 2 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
	fmt.Println("%v %v ", intVal, intVal1)

	// 声明一个变量并初始化
	var a = "RUNOOB"
	fmt.Println(a)

	// 没有初始化就为零值
	var b3 int
	fmt.Println(b3)

	// bool 零值为 false
	var c bool
	fmt.Println(c)

}
