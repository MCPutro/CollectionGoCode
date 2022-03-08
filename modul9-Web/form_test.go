package modul9_Web_test

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func checkForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	value1 := r.FormValue("param1")
	fmt.Println(value1)
	value11 := r.PostFormValue("param1")
	fmt.Println(value11)
	value2 := r.PostFormValue("param2")
	fmt.Println(value2)

	fmt.Fprint(w, "ok: "+value1+" -> "+value2)
}

func TestSampleForm(t *testing.T) {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9999"
	}
	fmt.Println("server running in port ", PORT)

	myRoute := mux.NewRouter()
	myRoute.HandleFunc("/form", checkForm).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+PORT, myRoute))
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("param1=aku&param2=adalah")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	checkForm(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
