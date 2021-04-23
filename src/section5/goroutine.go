package main

import (
	"fmt"
	"sync"
)

func goro(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
	wg.Done()
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	// 関数の先頭に「go」とつけるだけで、平行処理が出来るようになる。
	go goro("world", &wg)
	// go で宣言したのは、前の処理が終わってから処理される
	normal("hello")
	// time.Sleep(200 * time.Microsecond)

	wg.Wait()
}
