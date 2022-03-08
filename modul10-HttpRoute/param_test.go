package modul10_HttpRoute

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	myRoute := mux.NewRouter()

	myRoute.HandleFunc("/user/{sm}/data", Subscribe).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", myRoute))

}

func Subscribe(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	sm := param["sm"]

	w.Write([]byte(sm))
}

//Unit Test
func TestUnitParams(t *testing.T) {
	//request := httptest.NewRequest(http.MethodGet, "/user/hahaha/data", nil)
	//recorder := httptest.NewRecorder()
	//
	//Subscribe(recorder, request)
	//body, _ := io.ReadAll(recorder.Result().Body)
	//fmt.Println(string(body))
	r := httptest.NewRequest("GET", "http://localhost:8080", nil)
	w := httptest.NewRecorder()

	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"sm": "hahaha",
	}

	// CHANGE THIS LINE!!!
	r = mux.SetURLVars(r, vars)

	Subscribe(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, []byte("hahaha"), w.Body.Bytes())
}
