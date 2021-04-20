package main

import (
	"fmt"
)

func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100

	var k int = 50
	one(&n)
	fmt.Println(&n)

	fmt.Println(*&k)
	// fmt.Println(n)

	// // Addressをmemoryに変換
	// // p　ポインタ変数
	// // *int　ポインタ型
	// var p *int = &n

	// // Address（アドレス出力：16進数)へ変換
	// fmt.Println(&n)
	// // 0xc0000ac008 アドレス

	// fmt.Println(*p)
	// 100

}
