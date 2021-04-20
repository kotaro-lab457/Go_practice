package main

import "fmt"

// 構造体 struct
// 構造体の定義方法
type Vertex struct {
	X int
	Y int
	S string
}

func main() {
	// {} で順番にフィールドの値を渡す方法⏬
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	// {1 2}
	fmt.Println(v.X, v.Y)
	// 1 2

	v2 := Vertex{X: 1}
	fmt.Println(v2)
	// {1 0} Y が 0になっている デフォルトの初期値が入っている
	// string　だと空になる

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)
	// {1 2 test}

	v4 := Vertex{}
	fmt.Println(v4)
	// {0 0 } 空の場合は、デフォルトの値になる

	var v5 Vertex
	fmt.Println(v5)

	v6 := new(Vertex)
	fmt.Println(v6)
	// &{0 0 }

	v7 := &Vertex{} // ポインタの型
	fmt.Printf("%T %v\n", v7, v7)
	// *main.Vertex &{0 0 }
}
