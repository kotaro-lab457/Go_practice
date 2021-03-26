package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	fmt.Println(time.Now().Clock())
	fmt.Println(user.Current())
}
