package main

import (
	"fmt"
)

func num() {
	a, b := 10, 20            // コロン＆イコールで変数名が定義できる
	var xi float64 = 1.234567 // 「float」は少数
	fmt.Println(a, b, xi)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", xi)
}

// varで定義する変数名は関数の外でも宣言できる
var (
	ai   int    = 1000   // 「int」は整数
	s    string = "test" // 静的型付け（データ型）
	t, f bool   = true, false
)

func main() {
	num()
	fmt.Println(ai, s, t, f)
}
