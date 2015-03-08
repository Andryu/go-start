package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// 環境変数設定
	os.Setenv("hoge", "fuga")
	// 環境変数取得
	fmt.Println("hoge is", os.Getenv("hoge"))

	// 環境変数をリストで取得
	for i, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Printf("%d: name=%s value=%s\n", i, pair[0], pair[1])

	}
}
