package main

import "fmt"

func main() {

	//初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
	//如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：

	var balance [10]float32
	var balance2 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var balance3 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance, balance2, balance3)

	var n1 [10]int /* n1 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n1[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n1[j])
	}

	//Go 语言支持多维数组

	var a4 = [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}
	//注意：以上代码中倒数第二行的 } 必须要有逗号，因为最后一行的 } 不能单独一行，也可以写成这样：

	var a5 = [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}} /* 第三行索引为 2 */

	fmt.Println(a4, a5)

	/* 数组 - 5 行 2 列*/
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var m, n int

	/* 输出数组元素 */
	for m = 0; m < 5; m++ {
		for n = 0; n < 2; n++ {
			fmt.Printf("a[%d][%d] = %d\n", m, n, a[m][n])
		}
	}

	//向函数传递数组参数，你需要在函数定义时，声明形参为数组
	/* 数组长度为 5 */
	var balance4 = []int{1000, 2, 3, 17, 50}
	var avg float32

	/* 数组作为参数传递给函数 */
	avg = getAverage(balance4, 5)

	/* 输出返回的平均值 */
	fmt.Printf("平均值为: %f ", avg)
	a6 := 1.69
	b := 1.7
	c := a6 * b    // 结果应该是2.873
	fmt.Println(c) // 输出的是2.8729999999999998

	a7 := 1690                         // 表示1.69
	b7 := 1700                         // 表示1.70
	c7 := a7 * b7                      // 结果应该是2873000表示 2.873
	fmt.Println(c7)                    // 内部编码
	fmt.Println(float64(c7) / 1000000) // 显示

	//初始化数组长度后，元素可以不进行初始化，或者不进行全部初始化，但未进行数组大小初始化的数组初始化结果元素大小就为多少。

	var array = []int{1, 2, 3, 4, 5}
	/* 未定义长度的数组只能传给不限制数组长度的函数 */
	setArray(array)
	/* 定义了长度的数组只能传给限制了相同数组长度的函数 */
	var array2 = [5]int{1, 2, 3, 4, 5}
	setArray2(array2)

	var arr = [][]float32{{-1, -2}, {-3, -4}, {-5}}
	prt(arr)

}

func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}

func setArray(params []int) {
	fmt.Println("params array length of setArray is : ", len(params))
}

func setArray2(params [5]int) {
	fmt.Println("params array length of setArray2 is : ", len(params))
}

func prt(arr [][]float32) {
	for i := 0; i < 3; i++ {
		println(arr[i][0])
	}
}
