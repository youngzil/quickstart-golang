package practice

import "fmt"

func main() {

	//类型转换用于将一种数据类型的变量转换为另外一种类型的变量。Go 语言类型转换基本格式如下：
	//type_name(expression)
	//type_name 为类型，expression 为表达式。

	//以下实例中将整型转化为浮点型，并计算结果，将结果赋值给浮点型变量：
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)
}
