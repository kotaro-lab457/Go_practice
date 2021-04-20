package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./loop.go")
	if err != nil {
		log.Fatalln("error!")
	}
	// ファイルがなかったら、エラー
	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error")
	}

	fmt.Println(count, string(data))

	// 出力結果
	// 	100 package main 「100」は読み込んだ数(count,コード)
	// import "fmt"
	// func main() {

	//         // forループ処理
	//         for i := 0; i < 10; i++ {

	//                 // �

	// 簡略化 nilは、ゼロ値として扱い型の指定がない。ない値と考えたほうがいいかな。。
	if err = os.Chdir("test"); err != nil {
		log.Fatalln("error")
	}
}
