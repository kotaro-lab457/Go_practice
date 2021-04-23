package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 100
	// fmt.Println(ch) 0xc0000220e0
	fmt.Println(len(ch))
	// 1
	ch <- 200
	fmt.Println(len(ch))
	// 2

	// ch <- 300
	// fmt.Println(len(ch)) bufferedの数が２で指定されているからエラーが出る。

	// x := <-ch チャネルを取り出す 100
	// fmt.Println(x)

	// チャネルの読み込みを防ぐ
	close(ch)

	for c := range ch {
		fmt.Println(c)
	}
}
