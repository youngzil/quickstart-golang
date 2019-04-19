package main

import "fmt"

/* 声明全局变量 */
var a2 int = 20

func main() {
	/* main 函数中声明局部变量 */
	var a2 int = 10
	var b int = 20
	var c int = 0

	fmt.Printf("main()函数中 a2 = %d\n", a2)
	c = sum2(a2, b)
	fmt.Printf("main()函数中 a2 = %d\n", a2)
	fmt.Printf("main()函数中 c = %d\n", c)

	//在 for 循环的 initialize（a:=0） 中，此时 initialize 中的 a 与外层的 a 不是同一个变量，initialize 中的 a 为 for 循环中的局部变量，因此在执行完 for 循环后，输出 a 的值仍然为 0。
	var a int = 0
	fmt.Println("for start")
	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}
	fmt.Println("for end")

	fmt.Println(a)

	//此时 initialize 中的 d 便于外层的 d 为同一个变量，因此在执行完 for 循环后，输出 d 的值为 10。
	var d int = 0
	fmt.Println("for start")
	for d = 0; d < 10; d++ {
		fmt.Println(d)
	}
	fmt.Println("for end")
	fmt.Println(d)

}

/* 函数定义-两数相加 */
func sum2(a, b int) int {
	a = a + 1
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)
	return a + b
}
