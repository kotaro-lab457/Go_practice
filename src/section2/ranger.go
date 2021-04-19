package main

import "fmt"

func main() {
	lan := []string{"python", "go", "java"}

	for i := 0; i < len(lan); i++ {
		fmt.Println(i, lan[i])
	}
	// 0 python
	// 1 go
	// 2 java

	// range ... インデックス番号と配列をまとめてくれる
	for i, v := range lan {
		fmt.Print(i, v)
	}
	// 0python1go2java

	// インデックス番号なし アンダーバー
	for _, v := range lan {
		fmt.Print(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
	// apple 100
	// banana 200

	for k := range m {
		fmt.Println(k)
	}
	// apple
	// banana

	for _, v := range m {
		fmt.Println(v)
	}
	// 100
	// 200
}
