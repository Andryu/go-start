
package main

// ベンチマークテスト用
// ただ jsonを返却するAPI
// redisなど他のプログラムと比べるため使用する
import (
   "encoding/json"
   "fmt"
   "net/http"
   "log"
    //"os"
    //"github.com/garyburd/redigo/redis"
)

type CommonParams struct {
    Result string `json:"result"`
    Status   string `json:"status"`
}

type ResponseParams CommonParams


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
    return json.Marshal(struct {
            ResponseParams
            Bonus string `json:"bonus"`
        }{
            ResponseParams: ResponseParams(CommonParams{ "success", "200"}),
            Bonus:    bonus,
        })
}

func getRank(w http.ResponseWriter, r *http.Request){
    //r.ParseForm()        // オプション解析
    //bonus := r.FormValue("bonus")

    result, _ := s_ok()
    fmt.Fprintf(w, string(result))
}

func updateRank(w http.ResponseWriter, r *http.Request){
    r.ParseForm()        // オプション解析
    //bonus := r.FormValue("bonus")
    //point := r.FormValue("point")

    result, _ := s_ok()
    fmt.Fprintf(w, string(result))
}


func main() {
    //http.HandleFunc("/rankup", updateRank)
    http.HandleFunc("/rank", getRank)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
