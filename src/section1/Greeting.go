package main

// フォーマットパッケージ
import (
	"fmt"
)

// init関数が先に出力される
func init() {
	fmt.Println("init!")
}

func bazz() {
	fmt.Println(("Good morning"))
}

//  エントリーポイント
func main() {
	bazz()                      // 関数名を入れることで呼び出せる
	fmt.Println("Hello world!") // fmtパッケージのPrintln関数
}

// コンパイルするため、ビルド（build）を行う。「 go run Greeting.go 」
