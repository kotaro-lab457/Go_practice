package main

import (
	"fmt"
	"strconv"
)

// 定数は大文字から、また型宣言しなくてよし
const Pi = 3.14

const big = 10000000000 + 1 // 変数もできる

const (
	Username = "test_user"
	Password = "test_pass"
)

func text() {
	fmt.Println(Pi, Username, Password)
	fmt.Println(big + 2)
}

// 型変換
func text2() {
	var x int = 1
	xx := float64(x) // 少数へ型変換
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y) //整数へ型変換
	fmt.Printf("%T %v %d\n", yy, yy, yy)
}

func text3() {
	var s string = "14"
	i, _ := strconv.Atoi(s) // 文字列を数字へ変換「 strconv.Atoi() 」
	fmt.Printf("%T %v", i, i)
}

func main() {
	text()
	text2()
	text3()
}
