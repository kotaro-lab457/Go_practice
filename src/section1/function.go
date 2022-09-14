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
func cal(price, item int) (result int) {
	result = price * item
	return
}

const Pi = 3.14

func main() {
	add(10, 20)

	// 型指定なし（暗黙的）
	r1, r2 := add3(30, 40)
	fmt.Println(r1, r2)
	// 70 -10

	r3 := cal(10, 2)
	fmt.Println(r3)
	// 20

	var a2 int64 = 500
	var a3 int = 1
	fmt.Printf("%T\n", int32(a3))
	fmt.Println((a2 + int64(a3)))

	var s string = "Hello"

	fmt.Println(string(s[0]))

	// var arr2 [15]string = [15]string{"A", "B"}
	// fmt.Println(arr2)
	var arr2 [2]string = [...]string{"A", "B"}
	fmt.Println(arr2)

	var arrB [2]int = [2]int{100, 900}

	arrB[1] = 300

	fmt.Println(arrB)
}
