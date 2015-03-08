package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

type postHelloInput struct {
	ID   int
	Name string
}

type postHelloOutput struct {
	ID     int
	Result string
}

func postHello(w rest.ResponseWriter, req *rest.Request) {
	params := req.URL.Query()
	fmt.Println("params------")
	fmt.Println(params)

	id := params.Get("ID")
	Name := params.Get("Name")
	fmt.Println(id)
	fmt.Println(Name)
	fmt.Println("---------------")
	input := postHelloInput{}
	//err := req.DecodeJsonPayload(&input)

	//// そもそも入力の形式と違うとここでエラーになる
	//if err != nil {
	//    fmt.Println(err)
	//    rest.Error(w, err.Error(), http.StatusInternalServerError)
	//    return
	//}

	// 適当なバリデーション
	//if input.Name == "" {
	//    rest.Error(w, "Name is required", 400)
	//    return
	//}

	log.Printf("%#v", input)

	// 結果を返す部分
	w.WriteJson(&postHelloOutput{
		input.ID,
		"Hello, " + input.Name,
	})
}

func main() {
	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
	}
	handler.SetRoutes(
		&rest.Route{"POST", "/hello", postHello},
		&rest.Route{"GET", "/hello", postHello},
	)
	log.Printf("Server started")
	http.ListenAndServe(":9999", &handler) // ポート9999で立ち上がる
}
