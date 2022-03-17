package gorilla_mux

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"testing"
)

func TestServer3(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/", checkAPI)
	r.Handle("/products", authMiddleware(http.HandlerFunc(ProductsHandler)))
	r.HandleFunc("/articles", ArticlesHandler)
	//http.Handle("/", r)

	//dengan middleware
	//http.ListenAndServe(":1212", middleware.NewAuthMiddleware(r))

	//tanpa middleware
	http.ListenAndServe(":1212", r)
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//logic
		fmt.Println("ini authMiddleware")

		next.ServeHTTP(w, r)
	})
}
