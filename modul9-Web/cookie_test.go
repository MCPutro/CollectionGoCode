package modul9_Web_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{}
	cookie.Name = "coba-cookie"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Expires = time.Now().Add(5 * time.Minute)

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "ok")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("coba-cookie")
	if err != nil {
		fmt.Fprint(writer, "No Cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestSampleCodeWithCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//Unit test
func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=aku", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s \n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "coba-cookie"
	cookie.Value = "baiklah"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
