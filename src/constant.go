package main

import "fmt"

// 定数は大文字から、また型宣言しなくてよし
const Pi = 3.14

const big = 10000000000 + 1 // 変数もできる

const (
	Username = "test_user"
	Password = "test_pass"
)

func main() {
	fmt.Println(Pi, Username, Password)
	fmt.Println(big + 2)
}
