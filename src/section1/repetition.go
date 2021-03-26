package main

import "fmt"

func main() {
	n := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(n)
	// map[apple:100 banana:200]

	fmt.Println(n["apple"])
	// 100  appleの数値

	n["banana"] = 300
	fmt.Println(n)
	// map[apple:100 banana:300] 300へ変換

	n["new"] = 500
	fmt.Println(n)
	// map[apple:100 banana:300 new:500]

	fmt.Println(n["nothing"])
	// 0

	v, ok := n["apple"]
	fmt.Println(v, ok)
	// 100 true 値が入っているから「 true 」

	v2, ok2 := n["nothing"]
	fmt.Println(v2, ok2)
	// 0 false

	n2 := make(map[string]int)
	n2["pc"] = 5000
	fmt.Println((n2))
	// map[pc:500]

	// var n3 map[string]int
	// n3["pc"] = 5000
	// fmt.Println(m3)
	// error undefinedのエラーが出る
}
