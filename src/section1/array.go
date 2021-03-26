package main

import "fmt"

//配列
func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b [2]int = [2]int{100, 200} // 型の位置が少し変わる
	fmt.Println(b)

	var c []int = []int{100, 200}
	c = append(c, 300) // 代入
	fmt.Println(c)
}
