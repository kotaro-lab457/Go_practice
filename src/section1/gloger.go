package main

import "fmt"

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// 円周率
func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	// 1
	// 2
	// 3

	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	// 1
	// 2
	// 3
	// 4

	c1 := circleArea(3.14)
	fmt.Println(c1(2))
	// 12.56
	fmt.Println(c1(3))
	// 28.259999999999998
}
