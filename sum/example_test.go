package sum

import "fmt"

func Example() {
	fmt.Println("A")
	// Output:
	// A
}

func ExampleSum() {
	fmt.Println(Sum([]int{3, 5}...))
	// Output:
	// 8
}
