package main

import (
	"fmt"
)

func do(i interface{}) {
	// interfaceで型が決まっていない（int型） タイプアサーション
	// I := i.(int)
	// // i = I * 2
	// I *= 2
	// fmt.Println(I)
	// S := i.(string)
	// fmt.Println(S + "!")
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
		// 20
	case string:
		fmt.Println(v + "!")
		// Kei!
	default:
		fmt.Printf("%T\n", v)
		// bool
	}

}

// インターフェースできたものをタイプアサーションすることで、型にふさわしい処理ができる

func main() {
	// var i interface{} = 10
	do(10)
	do("Kei")
	do(true)
}
