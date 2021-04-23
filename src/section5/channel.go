package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// c に向かって sumの値をchannelしてあげる（<-）送信
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go goroutine1(s, c)
	go goroutine1(s, c)
	// c の値が入ったら受信し、x へ代入される。ブロッキング
	x := <-c
	fmt.Println(x)
	// 15
	y := <-c
	fmt.Println(y)
}
