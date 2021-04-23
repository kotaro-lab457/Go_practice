package main

import (
	"fmt"
	"sync"
)

// ch(引数)
func producer(ch chan int, i int) {
	// ② ch に値が渡ってきたのを処理する。
	ch <- i * 2
}

// ③ チャネルから渡ってきた値を range　で回す。
func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("process", i*1000)
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// producer
	// ①10回分 wg,Add(1)を呼び出し。producer の ch に値を渡す。
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	go consumer(ch, &wg)
	wg.Wait()
	close(ch)
	// consume rが goroutine としてまだ動いているため chを探しちゃう。
}
