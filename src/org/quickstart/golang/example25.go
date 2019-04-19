package main

import "fmt"

func main() {

	//如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。
	//当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址：

	var a int
	var ptr *int
	var pptr **int /*指向指针的指针变量为整型*/

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr) /*访问指向指针的指针变量值需要使用两个 * 号*/

	var a2 int = 1
	var ptr1 *int = &a2
	var ptr2 **int = &ptr1
	var ptr3 **(*int) = &ptr2 // 也可以写作：var ptr3 ***int = &ptr2
	// 依次类推
	fmt.Println("a2:", a2)
	fmt.Println("ptr1:", ptr1)
	fmt.Println("ptr2:", ptr2)
	fmt.Println("ptr3:", ptr3)
	fmt.Println("*ptr1:", *ptr1)
	fmt.Println("**ptr2:", **ptr2)
	fmt.Println("**(*ptr3):", **(*ptr3)) // 也可以写作：***ptr3

}
