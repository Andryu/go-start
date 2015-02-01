package main

import (
   "fmt"
   "net/http"
   "log"
    "os"
    "github.com/garyburd/redigo/redis"
)

type ResponseParam struct {
    result string
    code   string
}

func rankUp(w http.ResponseWriter, r *http.Request){
    r.ParseForm()        // オプション解析
    bonus := r.FormValue("bonus")
    point := r.FormValue("point")
    //connect_redis("bonus:" + bonus, point)

    c := pool.Get()
    defer c.Close()

    test, _ := c.Do("SET", "bonus:" + bonus, point)
    fmt.Println(test)
}

func connect_redis(key string, value string){
    c, err := redis.Dial("tcp", "127.0.0.1:6379")

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    defer c.Close()

    c.Do("SET", key, value)
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
    http.HandleFunc("/rankup", rankUp)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
