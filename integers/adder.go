package integers

import "fmt"

func Add(x, y int) int {
	return x + y
}

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
}
