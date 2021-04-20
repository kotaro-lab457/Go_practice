package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")
	fmt.Println("hello foo")
}

func main() {

	//foo()
	// hello foo
	// world foo   foo()関数が終わった後に、deferの処理が来る。

	// defer ... 遅延実行
	// defer fmt.Println("world")
	// fmt.Println("hello")
	// hello
	// world

	// fmt.Println("run")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("success")
	// スタッティングデバー？
	// run
	// success
	// 3
	// 2
	// 1

	file, _ := os.Open("./loop.go") // ファイルを開く
	defer file.Close()              // 遅延でファイルを閉じる
	data := make([]byte, 100)       // バイト配列？
	file.Read(data)                 // ファイルを読み込む
	fmt.Println(string(data))       // 文字列で表示
}
