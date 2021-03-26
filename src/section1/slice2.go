package main

import "fmt"

func main() {
	n := make([]int, 3, 5)

	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	// en=3 cap=5 value=[0 0 0]%

	n = append(n, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	// len=5 cap=5 value=[0 0 0 0 0]

	n = append(n, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	// len=10 cap=10 value=[0 0 0 0 0 1 2 3 4 5]

	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)
	// len=3 cap=3 value=[0 0 0]

	b := make([]int, 0) //メモリーに確保
	var c []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)
	// len=0 cap=0 value=[]
	// len=0 cap=0 value=[]

	// for文
	c = make([]int, 3)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	// [0 0 0 0]
	// [0 0 0 0 1]
	// [0 0 0 0 1 2]
	// [0 0 0 0 1 2 3]
	// [0 0 0 0 1 2 3 4]
}
