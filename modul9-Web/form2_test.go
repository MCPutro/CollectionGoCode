package modul9_Web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"testing"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Fprintf(w, "Hello world")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Fprintf(w, "Mehhhhh?")
}

func input(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/input.gohtml")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("Your Name:", r.Form["name"])
		fmt.Println("Your Address:", r.Form["address"])
		fmt.Fprintf(w, "Your name is %s and you are located in %s", r.Form["name"], r.Form["address"])
	}
}

func TestMainServer(t *testing.T) {
	http.HandleFunc("/", index)
	http.HandleFunc("/me", me)
	http.HandleFunc("/input/test", input)
	err := http.ListenAndServe(":6969", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
