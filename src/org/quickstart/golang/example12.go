package main

import "fmt"

func main() {
	/* 局部变量定义 */
	var a int = 100

	/* 判断布尔表达式 */
	/* 使用 if 语句判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	} else {
		/* 如果条件为 false 则执行以下语句 */
		fmt.Printf("a 不小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)

	//语言 if 语句嵌套
	/* 定义局部变量 */
	var b int = 200
	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n")
		}
	}
	fmt.Printf("a 值为 : %d\n", a)
	fmt.Printf("b 值为 : %d\n", b)

	//if ... else if ... else... 实例：
	var age int = 23
	if age == 25 {
		fmt.Println("true")
	} else if age < 25 {
		fmt.Println("too small")
	} else {
		fmt.Println("too big")
	}

	//寻找到 100 以内的所有的素数
	// var count,c int   //定义变量不使用也会报错
	var count int
	var flag bool
	count = 1
	//while(count<100) {    //go没有while
	for count < 100 {
		count++
		flag = true
		//注意tmp变量  :=
		for tmp := 2; tmp < count; tmp++ {
			if count%tmp == 0 {
				flag = false
			}
		}

		// 每一个 if else 都需要加入括号 同时 else 位置不能在新一行
		if flag == true {
			fmt.Println(count, "素数")
		} else {
			continue
		}
	}

	//判断输入的数字是奇偶性
	var s int // 声明变量 s 是需要判断的数
	fmt.Println("输入一个数字：")
	fmt.Scan(&s)

	if s%2 == 0 { //     取 s 处以 2 的余数是否等于 0
		fmt.Println("%d是偶数\n", s) //如果成立
	} else {
		fmt.Println("%d不是偶数\n", s) //否则
	}
	fmt.Println("%d的值是：", s)

	//条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了，如下所示:
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	/*（1） 不需使用括号将条件包含起来
	（2） 大括号{}必须存在，即使只有一行语句
	（3） 左括号必须在if或else的同一行
	（4） 在if之后，条件语句之前，可以添加变量初始化语句，使用；进行分隔
	（5） 在有返回值的函数中，最终的return不能在条件语句中*/

	//判断用户密码输入
	var apass int
	var bpass int
	fmt.Printf("请输入密码：   \n")
	fmt.Scan(&apass)
	if apass == 5211314 {
		fmt.Printf("请再次输入密码：")
		fmt.Scan(&bpass)
		if bpass == 5211314 {
			fmt.Printf("密码正确，门锁已打开")
		} else {
			fmt.Printf("非法入侵，已自动报警")
		}
	} else {
		fmt.Printf("非法入侵，已自动报警")
	}

}
