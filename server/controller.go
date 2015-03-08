package main

// ベンチマークテスト用
// ただ jsonを返却するAPI
// redisなど他のプログラムと比べるため使用する
import (
  "encoding/json"
  "fmt"
  "net/http"
  //"os"
  //"github.com/garyburd/redigo/redis"
  //"./logs"
)

type CommonParams struct {
  Result string `json:"result"`
  Status string `json:"status"`
}

type ResponseParams CommonParams

func e404() ([]byte, error) {
  res := CommonParams{"failed", "404"}
  return json.Marshal(res)
}

func ok() ([]byte, error) {
  res := CommonParams{"success", "200"}
  return json.Marshal(res)
}

func Index(w http.ResponseWriter, r *http.Request) {
  r.ParseForm() // オプション解析

  //logs.Logger.Info("Hogeeeeeeeeeeeeeeeeeeeeeeeee")
  result, _ := ok()
  fmt.Fprintf(w, string(result))
}
