package main

import (
	"fmt"
)

var pool = newPool()

func main() {
	// api 認証
	auth := new(Authentication)
	auth.SetKeyForApp("access_token")
	fmt.Println(auth.Authenticate("jMI3uIk1j-7PFNGStR9JKrDP8QqZfZ4LFVbZb_0yhK4%3D", "message"))
}
