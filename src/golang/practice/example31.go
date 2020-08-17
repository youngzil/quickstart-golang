package main

import "fmt"

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func main() {

	//Go 语言支持递归。但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。
	//递归函数对于解决数学上的问题是非常有用的，就像计算阶乘，生成斐波那契数列等。

	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))

	//以下实例通过 Go 语言的递归函数实现斐波那契数列：
	//斐波纳契数列以如下被以递归的方法定义：F0=0，F1=1，Fn=F(n-1)+F(n-2)（n>=2，n∈N*）。
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	fmt.Println()

	//
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci3(i))
	}
	fmt.Println()

	//求平方根
	//原理: 计算机通常使用循环来计算 x 的平方根。从某个猜测的值 z 开始，我们可以根据 z² 与 x 的近似度来调整 z，产生一个更好的猜测：
	//z -= (z*z - x) / (2*z)
	//重复调整的过程，猜测的结果会越来越精确，得到的答案也会尽可能接近实际的平方根。
	fmt.Println(get_sqrt(2))
	fmt.Println(get_sqrt(3))
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

//更好的一种 fibonacci 实现，用到多返回值特性，降低复杂度：
func fibonacci2(n int) (int, int) {
	if n < 2 {
		return 0, n
	}
	a, b := fibonacci2(n - 1)
	return b, a + b
}

func fibonacci3(n int) int {
	_, b := fibonacci2(n)
	return b
}

func sqrt(x float64, i float64) (float64, float64) {
	remain := (i*i - x) / (2 * i)
	i = i - remain
	if remain > 0 {
		return sqrt(x, i)
	} else {
		return i, remain
	}
}
func get_sqrt(x float64) float64 {
	i, _ := sqrt(x, x)
	return i
}
