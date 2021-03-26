package main

import "fmt"

func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params { // intを繰り返し出力
		fmt.Println(param)
	}
}

func main() {
	foo(10, 20)
	// 2 [10 20]
	foo(10, 20, 30)
	// 3 [10 20 30]

	s := []int{1, 2, 3}
	fmt.Println(s)

	foo(s...)
	// [1 2 3]
	// 3 [1 2 3]
	// 1
	// 2
	// 3
}
