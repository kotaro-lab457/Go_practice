package main

import "fmt"

// 型　接続
type Human interface {
	Say()
}

// class 定義
type Person struct {
	Name string
}

// Say() の型を呼び出す。。
func (i Person) Say() {
	i.Name = "Mr." + i.Name
	fmt.Println(i.Name)
}

// func DriveCar(human Human) {
// 	if human.Say() == "Mr.Kei" {
// 		fmt.Println("Run")
// 	} else {
// 		fmt.Println("Get out")
// 	}
// }

func main() {
	// インターフェースにPersonを入れる
	var Kei Human = Person{"Kei"}
	Kei.Say()
	//DriveCar(Kei)

}
