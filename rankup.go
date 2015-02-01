package main

import (
   "encoding/json"
   "fmt"
   "net/http"
   "log"
    //"os"
    "github.com/garyburd/redigo/redis"
)

type CommonParams struct {
    result string
    code   string
}

type ResponseParams CommonParams

//type ResponseParams struct {
//    common CommonParams
//    bonus  string
//}

func e404() ([]byte, error) {
    res := CommonParams{ "failed", "404" }
    return json.Marshal(res)
}

func s_ok() ([]byte, error) {
    res := CommonParams{ "success", "200"}
    log.Println(res)
    return json.Marshal(res)
}

func s_bonus(bonus string) ([]byte, error) {
    //res := ResponseParams{
    //    CommonParams{ " success", "200"},
    //    bonus,
    //}
    //return json.Marshal(res)
    return json.Marshal(struct {
            ResponseParams
            bonus string
        }{
            ResponseParams: ResponseParams(CommonParams{ "success", "200"}),
            bonus:    bonus,
        })
}


func updateRank(w http.ResponseWriter, r *http.Request){
    r.ParseForm()        // オプション解析
    bonus := r.FormValue("bonus")
    point := r.FormValue("point")

    c := pool.Get()
    defer c.Close()

    _, err := c.Do("SET", "bonus:" + bonus, point)
    if err != nil {
        e, _ := e404()
        fmt.Fprintf(w, string(e))
        return
    }
    result, err := s_ok()
    fmt.Println(string(result))
    fmt.Fprintf(w, string(result))
}

func getRank(w http.ResponseWriter, r *http.Request){
    r.ParseForm()        // オプション解析
    bonus := r.FormValue("bonus")

    c := pool.Get()
    defer c.Close()

    res, err := redis.String(c.Do("GET", "bonus:" + bonus))
    if err != nil {
        e, _ := e404()
        fmt.Fprintf(w, string(e))
        return
    }
    result, err := s_bonus(res)
    fmt.Println(string(result))
    fmt.Fprintf(w, string(result))
}

func newPool() *redis.Pool {
    return &redis.Pool{
        MaxIdle: 80,
        MaxActive: 12000, // max_number of connections
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", ":6379")
            if err != nil {
                panic(err.Error())
            }
            return c, err
        },
    }
}

var pool = newPool()

func main() {
    http.HandleFunc("/rankup", updateRank)
    http.HandleFunc("/rank", getRank)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
