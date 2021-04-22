package main

import "fmt"

// 出力結果を変えてやる。。
type Person struct {
	Name string
	Age  int
}

// 出力を変換（Personのstruct）
func (p Person) String() string {
	return fmt.Sprintf("Mr. %v", p.Name)
	// Mr.Kei
}

func main() {
	kei := Person{"Kei", 23}
	// {Kei 23}
	fmt.Println(kei)

}
