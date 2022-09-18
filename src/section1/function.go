package main

import (
	"fmt"
	"time"
)

func add(x int, y int) {
	fmt.Println(x + y)
	var data = time.Now().Day()
	fmt.Println(data)

}

// returnを2つ
func add3(x int, y int) (int, int) {
	return x + y, x - y
}

// 関数名を定義 result
func cal(price int, item int) (result int) {
	result = price * item
	return
}

const Pi = 3.14

func main() {
	add(10, 20)

	// 型指定なし（暗黙的）
	r1, _ := add3(30, 40)
	fmt.Println(r1)
	// 70 -10

	fmt.Printf("%T\n", r1)

	r3 := cal(10, 2)
	fmt.Println(r3)
	// 20

}
