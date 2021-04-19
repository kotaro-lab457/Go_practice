package main

import "fmt"

func main() {

	// forループ処理
	for i := 0; i < 10; i++ {

		// 継続（continue）
		if i == 3 {
			fmt.Println("continue")
			continue
		}

		// 中断（break）
		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}
	// 0 1 2 continue 4 5 6 7 8 9

	// 簡略 10を越えるまで繰り返す。。（無限ループ）
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)

	// forだけだと無限ループになってしまうので注意
	// for {
	// 	fmt.Println("hello")
	// }
}
