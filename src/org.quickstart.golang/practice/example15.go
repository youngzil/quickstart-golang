package practice

import "fmt"

func main() {

	//如果循环中条件语句永远不为 false 则会进行无限循环，我们可以通过 for 循环语句中只设置一个条件表达式来执行无限循环：
	/*for true {
		fmt.Printf("这是无限循环。\n");
	}*/

	//输出 1-100 素数
	var C, c int //声明变量
	C = 1        /*这里不写入FOR循环是因为For语句执行之初会将C的值变为1，当我们goto A时for语句会重新执行（不是重新一轮循环）*/
A:
	for C < 100 {
		C++ //C=1不能写入for这里就不能写入
		for c = 2; c < C; c++ {
			if C%c == 0 {
				goto A //若发现因子则不是素数
			}
		}
		fmt.Println(C, "是素数")
	}

	//Go 语言实现99乘法表
	for i := 1; i <= 9; i++ { // i 控制行，以及计算的最大值
		for j := 1; j <= i; j++ { // j 控制每行的计算个数
			fmt.Printf("%d*%d=%d ", j, i, j*i)
		}
		fmt.Println("")
	}

	//另一个方法输出 1-100 素数:
	var a, b int
	for a = 2; a <= 100; a++ {
		for b = 2; b <= (a / b); b++ {
			if a%b == 0 {
				break
			}
		}
		if b > (a / b) {
			fmt.Printf("%d\t是素数\n", a)
		}
	}

}
