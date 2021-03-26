package main

import "fmt"

func main() {
	b := []byte{72, 73}
	fmt.Println(b)
	fmt.Println(string(b))
	// [72 73]
	// HI 文字列

	// アスキーコードを元に、アルファベットを示している🔼🔼

	c := []byte("HI")
	fmt.Println(c)
	// [72 73]　へ変換
	fmt.Println(string(c))
	// HI 文字列変換
}

// ネットワーク関係て使いそうだそう
