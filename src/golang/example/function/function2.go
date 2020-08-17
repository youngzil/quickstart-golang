package function

import "fmt"

func Test2() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
}
