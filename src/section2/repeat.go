package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {

	// 関数呼び出し
	result := by2(10)
	if result == "ok" {
		fmt.Println("great")
	}
	// great

	// 簡略化 result2は中でしか使えない..
	if result2 := by2(10); result2 == "ok" {
		fmt.Println("great2")
	}
	// great2

	//num := 9
	// ４を２で割る = true
	// if num%2 == 0 {
	// 	fmt.Println(("by 2"))
	// } else if num%3 == 0 {
	// 	fmt.Println("by 3")
	// } else {
	// 	fmt.Println("else")
	// }

	x, y := 10, 20
	if x == 10 && y == 20 {
		fmt.Println("&&")
	}
	// &&

	if x == 10 || y == 10 {
		fmt.Println("||")
	}
	// ||
}
