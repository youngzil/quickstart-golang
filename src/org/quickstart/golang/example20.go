package main

import "fmt"

/* 声明全局变量 */
var g int

/* 声明全局变量 */
var h int = 20

/* 声明全局变量 */
var l2 int = 20

func main() {

	//局部变量、全局变量和形式参数。

	/* 声明局部变量 */
	var a, b, c int

	/* 初始化参数 */
	a = 10
	b = 20
	c = a + b
	g = a + b

	fmt.Printf("结果： a = %d, b = %d and c = %d\n", a, b, c)
	fmt.Printf("结果： a = %d, b = %d and g = %d\n", a, b, g)

	//Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。
	/* 声明局部变量 */
	var h int = 10

	fmt.Printf("结果： h = %d\n", h)

	/*初始化局部和全局变量
	不同类型的局部和全局变量默认值为：
	数据类型	初始化默认值
	int	0
	float32	0
	pointer	nil*/

	//形式参数会作为函数的局部变量来使用
	/* main 函数中声明局部变量 */
	var l2 int = 10
	var m int = 20
	var n int = 0

	fmt.Printf("main()函数中 l = %d\n", l2)
	n = sum(l2, m)
	fmt.Printf("main()函数中 n = %d\n", n)

}

/* 函数定义-两数相加 */
func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)

	return a + b
}
