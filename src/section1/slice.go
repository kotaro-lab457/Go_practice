package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n) // 配列オブジェクト

	fmt.Println(n[2:5]) // [3 4 5]
	fmt.Println((n[:])) // 全て

	n[2] = 100
	fmt.Println(n) // 3から100へ変更

	var board = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	// [[0 1 2] [3 4 5] [6 7 8]]
	fmt.Println(board)

	n = append(n, 100, 200) // 追加

	// [1 2 100 4 5 6 100 200]
	fmt.Println(n)
}
