package main

import (
	"fmt"
	//"log"
	"./logs"
	"net/http"
	//"log"
	"bytes"
	"flag"
	"github.com/facebookgo/grace/gracehttp"
	"text/template"
)

type LineOfLog struct {
	RemoteAddr  string
	ContentType string
	Path        string
	Query       string
	Method      string
	Body        string
}

var TemplateOfLog = `
Remote address:   {{.RemoteAddr}}
Content-Type:     {{.ContentType}}
HTTP method:      {{.Method}}
Path: {{.Path}}
Query string: {{.Query}}
Body: {{.Body}}

`

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bufbody := new(bytes.Buffer)
		bufbody.ReadFrom(r.Body)
		body := bufbody.String()

		line := LineOfLog{
			r.RemoteAddr,
			r.Header.Get("Content-Type"),
			r.URL.Path,
			r.URL.RawQuery,
			r.Method, body,
		}
		tmpl, err := template.New("line").Parse(TemplateOfLog)
		if err != nil {
			panic(err)
		}

		bufline := new(bytes.Buffer)
		err = tmpl.Execute(bufline, line)
		if err != nil {
			panic(err)
		}

		logs.Logger.Info(bufline.String())
		handler.ServeHTTP(w, r)
	})
}

var pool = newPool()
var (
	address0 = flag.String("a0", ":8080", "Zero address to bind to.")
)

func main() {
	// seelog
	//logs.Logger.Info("Start server")

	// API Auth
	auth := new(Authentication)
	auth.SetKeyForApp("access_token")
	fmt.Println(auth.Authenticate("jMI3uIk1j-7PFNGStR9JKrDP8QqZfZ4LFVbZb_0yhK4%3D", "message"))

	// router
	//http.HandleFunc("/", Index)
	//err := http.ListenAndServe(":9090", Log(http.DefaultServeMux))
	//if err != nil {
	//  //logs.Logger.Critical("Server err:%v", err)
	//}

	// gracefull
	flag.Parse()
	gracehttp.Serve(
		&http.Server{Addr: *address0, Handler: newHandler()},
	)

}

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	return mux
}
