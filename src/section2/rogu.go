package main

import (
	"fmt"
	"log"
	"os"
)

// func LoggingSettings(logFile, string) {
// 	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	multiLogFile := io.MultiWriter(os.Stdout, logfile)
// 	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
// 	log.SetOutput((multiLogFile))
// }

func main() {
	// LoggingSettings("test.log")
	_, err := os.Open("fafafa")
	if err != nil {
		log.Fatalln("error", err)
		// 2021/04/20 10:08:06 error open fafafa: no such file or directory
	}
	log.Println("logging")
	// 日付　時間　logging
	log.Printf("%T %v", "test", "test")
	// 日付　時間　string型 test

	// log.Fatalf("%T %v", "test", "test")
	log.Fatalln("error!!")
	// log.Fatalln() でログの出力が終了し、これ以降は出力されない。。

	fmt.Println("ok!") // 出力されない
}
