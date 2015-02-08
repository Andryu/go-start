package main

import (
	"fmt"
)

var pool = newPool()

func main() {
	auth := new(Authentication)
	auth.SetKeyForApi("1")
	fmt.Println(auth.Authenticate("i19IcCmVwVmMVz2x4hhmqbgl1KeU0WnXBgoDYFeWNgs%3D", "message"))
}
