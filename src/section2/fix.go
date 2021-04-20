package main

import "fmt"

func thirdPartyConnectDB() {
	// panic() ... 強制終了
	panic("Unable to connect database!")
}

func save() {
	defer func() {
		// recover() ... 強制終了（panic）を止めて、継続する
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()

	// 関数を上におくと、プログラムが終了してしまう。←thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("ok?")
}
