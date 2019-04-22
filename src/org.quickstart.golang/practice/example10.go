package practice

import "fmt"

func main() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	//&	返回变量存储地址	&a; 将给出变量的实际地址。
	//*	指针变量。	*a; 是一个指针变量

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)

	//
	fmt.Println("赋值前")
	fmt.Println("ptr=", ptr)
	fmt.Println("&a=", &a)

	/*  & 和 * 运算符实例 */
	ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("*ptr 为 %d\n", *ptr)

	//指针变量 * 和地址值 & 的区别：指针变量保存的是一个地址值，会分配独立的内存来存储一个整型数字。当变量前面有 * 标识时，才等同于 & 的用法，否则会直接输出一个整型数字。
	fmt.Println("赋值后")
	fmt.Println("ptr=", ptr)
	fmt.Println("&a=", &a)
}
