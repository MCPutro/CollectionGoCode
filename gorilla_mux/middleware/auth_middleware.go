package middleware

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

//Implement Handler interface
type authMiddleware struct {
	Route *mux.Router
}

func NewAuthMiddleware(route *mux.Router) *authMiddleware {
	return &authMiddleware{Route: route}
}

func (a *authMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if "RAHASIA" == r.Header.Get("Auth") {
		//<> bisa ditambahkan log before action

		// ok than lanjut kan
		a.Route.ServeHTTP(w, r) // Handler.ServeHTTP(w, r)

		//<> bisa ditambahkan log after action
	} else {
		// error than return response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		//resp, _ := json.Marshal(
		//	web.Response{
		//		Code:   http.StatusUnauthorized,
		//		Status: "UNAUTHORIZED",
		//	})

		w.Header().Add("Content-Type", "application/json")
		//w.Write(resp)
		fmt.Fprint(w, "UNAUTHORIZED")

	}
}
