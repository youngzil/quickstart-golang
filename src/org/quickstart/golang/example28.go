//Go 语言切片是对数组的抽象。
//Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

/*你可以声明一个未指定大小的数组来定义切片：
var identifier []type
切片不需要说明长度。

或使用make()函数来创建切片:
var slice1 []type = make([]type, len)
也可以简写为
slice1 := make([]type, len)

也可以指定容量，其中capacity为可选参数。
make([]T, length, capacity)
这里 len 是数组的长度并且也是切片的初始长度。

切片初始化

s :=[] int {1,2,3 }
直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3

s := arr[:]
初始化切片s,是数组arr的引用

s := arr[startIndex:endIndex]
将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片

s := arr[startIndex:]
缺省endIndex时将表示一直到arr的最后一个元素

s := arr[:endIndex]
缺省startIndex时将表示从arr的第一个元素开始

s1 := s[startIndex:endIndex]
通过切片s初始化切片s1

s :=make([]int,len,cap)
通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片*/

package main

import "fmt"

func main() {
	var numbers5 = make([]int, 3, 5)
	printSlice(numbers5)

	//一个切片在未初始化之前默认为 nil，长度为 0，实例如下：
	var numbers4 []int
	printSlice(numbers4)
	if numbers4 == nil {
		fmt.Println("切片是空的")
	}

	//切片截取:可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]，

	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

	//append() 和 copy() 函数
	//如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。

	fmt.Println("--------------------------------")
	var numbers6 []int
	printSlice(numbers6)

	/* 允许追加空切片 */
	numbers6 = append(numbers6, 0)
	printSlice(numbers6)

	/* 向切片添加一个元素 */
	numbers6 = append(numbers6, 1)
	printSlice(numbers6)

	/* 同时添加多个元素 */
	//当 numbers = [0, 1] 时，append(numbers, 2, 3, 4) 为什么 cap 从 2 变成 6 ？
	//经过实践得知，append(list, [params])，先判断 list 的 cap 长度是否大于等于 len(list) + len([params])，如果大于那么 cap 不变，否则 cap 等于 max{cap(list), cap[params]}，所以当 append(numbers, 2, 3, 4) cap 从 2 变成 6。
	numbers6 = append(numbers6, 2, 3, 4)
	printSlice(numbers6)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers7 := make([]int, len(numbers6), (cap(numbers6))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers7, numbers6)
	printSlice(numbers7)

	//我们可以看出切片，实际的是获取数组的某一部分，len切片<=cap切片<=len数组，切片由三部分组成：指向底层数组的指针、len、cap。

	/*sliceTest 函数是切片测试代码，简单的两种初始化方式。
	twoDimensionsArray 函数是二维数组测试函数。
	测试结果:
	1.二维数组中的元素(一位数组)个数 > 限制的列数,异常;
	2.二维数组中的元素(一位数组)个数 <= 限制的列数,正常;
	假设列数为 3, 某个一位数组 {1}, 则后两个元素,默认为 0。*/

	sliceTest()
	twoDimensionArray()

	//	我们基于原数组或者切片创建一个新的切片后，那么新的切片的大小和容量是多少呢？
	//	这里有个公式，对于底层数组容量是 k 的切片 slice[i:j] 来说：
	//长度: j-i
	//容量: k-i
	testCapacity()

	//合并多个数组
	testmerge()

	//在做函数调用时，slice 按引用传递，array 按值传递：
	changeSliceTest()

	//使用 copy 函数要注意对于 copy(dst, src)，要初始化 dst 的 size，否则无法复制。
	testCopy()

}

func printSlice(x []int) {

	//len() 和 cap() 函数
	//切片是可索引的，并且可以由 len() 方法获取长度。
	//切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少

	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func sliceTest() {
	arr := []int{1, 2, 3, 4, 5}
	s := arr[:]
	for e := range s {
		fmt.Println(s[e])
	}
	s1 := make([]int, 3)
	for e := range s1 {
		fmt.Println(s1[e])
	}
}

func twoDimensionArray() {
	/* 数组 - 5 行 2 列*/
	var a = [][]int{{0, 0}, {1, 2}, {2}, {3, 6}, {4, 8}}
	var i, j int

	/* 输出数组元素 */
	for i = 0; i < len(a); i++ {
		for j = 0; j < len(a[i]); j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

func testCapacity() {

	//	我们基于原数组或者切片创建一个新的切片后，那么新的切片的大小和容量是多少呢？
	//	这里有个公式，对于底层数组容量是 k 的切片 slice[i:j] 来说：
	//长度: j-i
	//容量: k-i

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSlice(numbers)
	fmt.Printf("%d\n", numbers[1:3])
	fmt.Printf("%d\n", numbers[2:7])
	fmt.Printf("%d\n", numbers[:3])
	fmt.Printf("%d\n", numbers[4:])
	number1 := make([]int, 0, 5)
	number2 := numbers[:3]
	printSlice(number1)
	printSlice(number2)
	number3 := numbers[2:5]
	printSlice(number3)
	number4 := numbers[3:8]
	printSlice(number4)
}

func testmerge() {

	var arr1 = []int{1, 2, 3}
	var arr2 = []int{4, 5, 6}
	var arr3 = []int{7, 8, 9}
	var s1 = append(append(arr1, arr2...), arr3...)
	fmt.Printf("s1: %v\n", s1)
}

func changeSliceTest() {
	arr1 := []int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 3}

	fmt.Println("before change arr1, ", arr1)
	changeSlice(arr1) // slice 按引用传递
	fmt.Println("after change arr1, ", arr1)

	fmt.Println("before change arr2, ", arr2)
	changeArray(arr2) // array 按值传递
	fmt.Println("after change arr2, ", arr2)

	fmt.Println("before change arr3, ", arr3)
	changeArrayByPointer(&arr3) // 可以显式取array的 指针
	fmt.Println("after change arr3, ", arr3)
}

func changeSlice(arr []int) {
	arr[0] = 9999
}

func changeArray(arr [3]int) {
	arr[0] = 6666
}

func changeArrayByPointer(arr *[3]int) {
	arr[0] = 6666
}
func testCopy() {

	//错误示例:
	dst := make([]int, 0)
	src := []int{1, 2, 3}
	copy(dst, src)
	printSlice(src)
	printSlice(dst)

	//正确示例:
	dst2 := make([]int, 3) // 令size=3
	src2 := []int{1, 2, 3}
	copy(dst2, src2)
	printSlice(src2)
	printSlice(dst2)
}
