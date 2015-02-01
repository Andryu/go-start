package main

// 基本的な接続方法とクエリパラメータの解析
import (
   "fmt"
   "net/http"
   "strings"
   "log"
)

func sayHelloName(w http.ResponseWriter, r *http.Request){
    log.Println(r.URL)
    q := r.URL.Query()
    fmt.Println(q)
    r.ParseForm()        // オプション解析
    fmt.Println(r.Form)  // サーバのプリント情報に出力
    fmt.Println(r.FormValue("hoge"))
    //id := params.Get("Id")
    //fmt.Println(id)
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key : ", k)
        fmt.Println("val : ", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!")
}

func main() {
    http.HandleFunc("/", sayHelloName)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

