package practice

import "fmt"

func main() {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)

	//switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}

	//使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
	var falseFlag bool = false
	var trueFlag bool = true
	switch {
	case falseFlag:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case trueFlag:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case falseFlag:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case trueFlag:
		fmt.Println("4、case 条件语句为 true")
	case falseFlag:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}

	/*1. 支持多条件匹配
	switch{
	case 1,2,3,4:
	default:
	}
	2. 不同的 case 之间不使用 break 分隔，默认只会执行一个 case。
	3. 如果想要执行多个 case，需要使用 fallthrough 关键字，也可用 break 终止。
	switch{
	case 1:
	...
	if(...){
	break
	}
	fallthrough // 此时switch(1)会执行case1和case2，但是如果满足if条件，则只执行case1
	case 2:
	...
	case 3:
	}*/

}
