package main

/*
  関数の戻り値を interface{}にすることで、どんなstructの型にも対応
/*
import (
    "fmt"
    "encoding/json"
)

type structA struct{
    Result string `json:"result"`
    Status string `json:"status"`
}

type structB struct{
    *structA     `json"structA"`
    Name string `json:"name"`
}

func get() structA {
    fmt.Println("hello world")
    return structA{ "success", "200" }
}

func getB() interface{} {
    fmt.Println("hello world B")
    return structB{ &structA{"success", "200"}, "Name is Tom" }
}

func main() {
  res, e := json.Marshal(get())
  if e != nil{
      fmt.Println(e)
  }
  fmt.Println(string(res))

  result, err := json.Marshal(getB())
  if err != nil {
      fmt.Println(err)
  }
  fmt.Println(string(result))
}
