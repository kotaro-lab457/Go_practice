package main

import "fmt"

func main() {

	// makeは空を返す
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	// new はポインタを返す
	var p *int = new(int)
	fmt.Printf("%T\n", p)

	// new() ... 初期化せず、０になる。返り値がアドレスでの表示
	// var p *int = new(int)
	// fmt.Println(*p)
	// // 0
	// *p++
	// fmt.Println(*p)
	// // 1

	// var p2 *int
	// fmt.Println(p2)
	// // <nil>
	// *p2++
	// fmt.Println(p2)
	// panicエラー nilには、数値がないので足すことができません
}
