package main

import "fmt"

// 自分でエラーを出しに行く。。

type userNotFound struct {
	Username string
}

// エラーを作成
func (e *userNotFound) Error() string {
	// エラーメッセージ
	return fmt.Sprintf("username: %v", e.Username)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &userNotFound{Username: "kei"}
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	} else {
		// エラーじゃない結果
		fmt.Println("err")
	}
}
