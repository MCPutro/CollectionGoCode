package gorilla_mux

//using middleware

import (
	"fmt"
	"github.com/MCPutro/CollectionGoCode/gorilla_mux/middleware"
	"github.com/gorilla/mux"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/", checkAPI)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	//http.Handle("/", r)

	//dengan middleware
	http.ListenAndServe(":1212", middleware.NewAuthMiddleware(r))

	//tanpa middleware
	//http.ListenAndServe(":1212", r)
}

func checkAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ProductsHandler")
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ProductsHandler")
}
