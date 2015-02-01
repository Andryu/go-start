package main

import (
   "fmt"
   "net/http"
   "log"
    //"os"
    "github.com/garyburd/redigo/redis"
)

type ResponseParam struct {
    result string
    code   string
}

func updateRank(w http.ResponseWriter, r *http.Request){
    r.ParseForm()        // オプション解析
    bonus := r.FormValue("bonus")
    point := r.FormValue("point")

    c := pool.Get()
    defer c.Close()

    test, _ := c.Do("SET", "bonus:" + bonus, point)
    fmt.Println(test)
}

func getRank(w http.ResponseWriter, r *http.Request){
    r.ParseForm()        // オプション解析
    bonus := r.FormValue("bonus")

    c := pool.Get()
    defer c.Close()

    test, _ := redis.String(c.Do("GET", "bonus:" + bonus))
    fmt.Println(test)
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
