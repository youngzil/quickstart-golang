package practice

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {

	//在上面的例子中，我们定义了一个接口Phone，接口里面有一个方法call()。然后我们在main函数里面定义了一个Phone类型变量，并分别为之赋值为NokiaPhone和IPhone。然后调用call()方法，输出结果如下：

	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	//
	var man Man
	man = new(Woman)
	fmt.Println(man.name())
	fmt.Println(man.age())

	//
	man = new(Men)
	fmt.Println(man.name())
	fmt.Println(man.age())

	/*func (name string) imp() string{
		print("这是实现方法的写法")
	}

	func sum(x int,y int) int{
		print("这是正常写法")
	}*/
}

type Man interface {
	name() string
	age() int
}

type Woman struct {
}

func (woman Woman) name() string {
	return "Jin Yawei"
}
func (woman Woman) age() int {
	return 23
}

type Men struct {
}

func (men Men) name() string {
	return "liweibin"
}
func (men Men) age() int {
	return 27
}
