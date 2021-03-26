package main

import "fmt"

func add(x int, y int) {
	fmt.Println("add function")
	fmt.Println(x + y)
	// add function
	// 30
}

func add2(x int, y int) int {
	return x + y
}

// returnを2つ
func add3(x int, y int) (int, int) {
	return x + y, x - y
}

// 関数名を定義 result
func cal(price, item int) (result int) {
	result = price * item
	return
}

func main() {
	add(10, 20)

	r := add2(10, 20)
	fmt.Println(r)
	// 30

	r1, r2 := add3(30, 40)
	fmt.Println(r1, r2)
	// 70 -10

	r3 := cal(100, 2)
	fmt.Println(r3)
	// 200

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)
	// inner func 1
}
