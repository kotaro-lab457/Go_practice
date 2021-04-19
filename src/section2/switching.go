package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "mac"
}

func main() {
	//os := getOsName()

	// switch文
	// 補足：関数を呼び出して使用できる
	switch getOsName() {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
		//default:
		//fmt.Println("Default!!")
	}

	// time.Time型 ... Nowなどのメソッドを使用する時間表示
	t := time.Now()
	// Hour() ... 現在時刻
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() > 12:
		fmt.Println("Afternoon")
	}
}
