package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
	// for と close で随時結果を返すようにする。
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	// チャネルを確保する場合は、len()を使用してあげればいい。
	c := make(chan int, len(s))
	go goroutine1(s, c)
	for i := range c {
		fmt.Println(i)
	}
}
