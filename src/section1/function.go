package main

import (
	"fmt"
	"time"
)

func add(x int, y int) {
	fmt.Println(x + y)
	data := time.Now()
	fmt.Println(data)
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

	r1, r2 := add3(30, 40)
	fmt.Println(r1, r2)
	// 70 -10

	r3 := cal(10, 2)
	fmt.Println(r3)
	// 20
}
