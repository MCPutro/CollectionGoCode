package gorilla_mux

//using auth for specific API

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"testing"
)

func TestServer1(t *testing.T) {
	r := router()
	http.ListenAndServe(":1212", r)
}

func router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	secure := router.PathPrefix("/auth").Subrouter()
	secure.Use(JwtVerify)
	secure.HandleFunc("/api", ApiHandler).Methods("GET")

	admin := router.PathPrefix("/user").Subrouter()
	admin.HandleFunc("/admin", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("OK"))
	}).Methods("GET")

	return router
}

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("x-access-token")

		//json.NewEncoder(w).Encode(r)
		header = strings.TrimSpace(header)

		if header == "" {
			w.WriteHeader(http.StatusForbidden)
			//json.NewEncoder(w).Encode("Missing auth token")
			w.Write([]byte("harus pake auth"))
			return
		} else {
			//json.NewEncoder(w).Encode(fmt.Sprintf("Token found. Value %s", header))
			w.Write([]byte(fmt.Sprintf("Token found. Value %s", header)))
		}
		next.ServeHTTP(w, r)
	})
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("SUCCESS!")
	//w.Write([]byte("SUCCESS"))
	return
}
