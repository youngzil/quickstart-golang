package practice

import "fmt"

const MAX int = 3

func main() {

	//变量是一种使用方便的占位符，用于引用计算机内存地址。
	//Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
	/*当一个指针被定义后没有分配到任何变量时，它的值为 nil。
	nil 指针也称为空指针。
	nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
	一个指针变量通常缩写为 ptr。*/
	var a int = 20  /* 声明实际变量 */
	var ip *int     /* 声明指针变量,指向整型 */
	var fp *float32 /* 指向浮点型 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
	fmt.Printf("fp 变量储存的指针地址: %x\n", fp)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)

	//if(ptr != nil)     /* ptr 不是空指针 */
	//if(ptr == nil)    /* ptr 是空指针 */

	//Go 语言指针数组
	a2 := []int{10, 100, 200}
	var i int
	var ptr2 [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr2[i] = &a2[i] /* 整数地址赋值给指针数组 */
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr2[i])
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] address = %d\n", i, ptr2[i])
	}

}
