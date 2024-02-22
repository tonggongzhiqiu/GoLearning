package httprouterlearning

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"testing"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "hello, ", ps.ByName("name"))
}

// 可以通过 context 获取 params
func Other(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	fmt.Fprint(w, "hello, the other one is ", params.ByName("name"))
}

type Ori struct{}

func (*Ori) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func TestRouter(t *testing.T) {
	server := http.Server{
		Handler: Ori,
	}
	http.ListenAndServe("localhost:8888", server)

	router := httprouter.New()
	router.GET("/index", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe("localhost:9999", router))
}
