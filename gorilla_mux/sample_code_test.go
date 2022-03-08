package gorilla_mux

import (
	"fmt"
	"github.com/MCPutro/CollectionGoCode/gorilla_mux/middleware"
	"github.com/gorilla/mux"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/", chectAPI)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	//http.Handle("/", r)

	//dengan middleware
	http.ListenAndServe(":1212", middleware.NewAuthMiddleware(r))

	//tanpa middleware
	//http.ListenAndServe(":1212", r)
}

func chectAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ProductsHandler")
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ProductsHandler")
}
