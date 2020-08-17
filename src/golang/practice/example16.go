package main

import "fmt"

func main() {

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	//使用循环嵌套来输出 2 到 100 间的素数：
	/* 定义局部变量 */
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}

	//Go 语言中 break 语句用于以下两方面：
	//用于循环语句中跳出循环，并开始执行循环之后的语句。
	//break 在 switch（开关语句）中在执行一条case后跳出语句的作用。

	/* 定义局部变量 */
	var a2 int = 10

	/* for 循环 */
	for a2 < 20 {
		fmt.Printf("a2 的值为 : %d\n", a2)
		a2++
		if a2 > 15 {
			/* 使用 break 语句跳出循环 */
			break
		}
	}

	//Go 语言的 continue 语句 有点像 break 语句。但是 continue 不是跳出循环，而是跳过当前循环执行下一次循环语句。
	//for 循环中，执行 continue 语句会触发for增量语句的执行。
	/* 定义局部变量 */
	var a3 int = 10

	/* for 循环 */
	for a3 < 20 {
		if a3 == 15 {
			/* 跳过此次循环 */
			a3 = a3 + 1
			continue
		}
		fmt.Printf("a3 的值为 : %d\n", a3)
		a3++
	}

	//Go 语言的 goto 语句可以无条件地转移到过程中指定的行。
	//goto 语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。
	//但是，在结构化程序设计中一般不主张使用 goto 语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难。
	/* 定义局部变量 */
	var a4 int = 10

	/* 循环 */
LOOP:
	for a4 < 20 {
		if a4 == 15 {
			/* 跳过迭代 */
			a4 = a4 + 1
			goto LOOP
		}
		fmt.Printf("a4的值为 : %d\n", a4)
		a4++
	}

}
