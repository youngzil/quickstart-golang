package main

import "fmt"

import "unsafe"

//常量还可以用作枚举：
const (
	a1 = "abc"
	b1 = len(a1)
	c1 = unsafe.Sizeof(a1)
)

func main() {

	//常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)

	//常量还可以用作枚举：
	println(a1, b1, c1)

	//字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节。
	var aa = "hello"
	println(unsafe.Sizeof(aa))

}
