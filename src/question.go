package main

import "fmt"

// 定義
func q1() {
	// int型に変換し出力してください。
	f := 1.11
	i := int(f)
	fmt.Println(i)

	m := map[string]int{
		"Mike": 20, "Nancy": 24, "Messi": 30,
	}
	fmt.Printf("%T %v", m, m)
}

// ステートメント
func q2() {

	// 一番小さい数を出力
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	var min int
	for i, v := range l {
		if i == 0 {
			min = v
			continue
		}

		if min >= v {
			min = v
		}
	}
	fmt.Println(min)

	// 果物の価格合計値
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	sum := 0
	for _, i := range m {
		sum += i
	}
	fmt.Println(sum)
}
func q3() {
	var i int = 100
	var j int = 200
	var k int = 30
	var p1 *int
	var p2 *int
	var p3 *int
	p1 = &i
	//fmt.Println(*p1)
	p2 = &j
	i = *p1 + *p2
	p3 = &k
	p2 = p3
	fmt.Println(*p2)
	j = *p2 + i
	fmt.Println(j)

}

type Vertex struct {
	X, Y int
}

// func (v Vertex) Plus() int {
// 	return v.X + v.Y
// }

func (v Vertex) String() string {
	return fmt.Sprintf("X is %d! Y is %d!", v.X, v.Y)
}

func q4() {
	v := Vertex{3, 4}
	// fmt.Println(v.Plus())
	fmt.Println(v)
}

func main() {
	// q1()
	// q2()
	// q3()
	q4()
}
