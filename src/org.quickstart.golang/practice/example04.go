package practice

import "fmt"

const (
	j = 1 << iota
	k = 3 << iota
	l
	m
)

func main() {
	//iota，特殊常量，可以认为是一个可以被编译器修改的常量。
	//在定义常量组时，如果不提供初始值，则表示将使用上行的表达式。
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
	fmt.Println("m=", m)
}
