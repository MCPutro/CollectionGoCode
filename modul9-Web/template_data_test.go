package modul9_Web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTMLFileWithMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml")) //load semua file di dalam directory

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"body":  "Batman",
	}

	t.ExecuteTemplate(writer, "simple2.gohtml", data) //name == filename
}

func TestSimpleHTMLFileWithMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFileWithMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Address struct {
	Street string
}

type Person struct {
	Name      string
	AddressXX Address
	Hobbies   []string
}

func (p Person) SayHallo(name string, name1 string) string {
	return "Hello, " + name + name1
}

//atribut struct harus public
func SimpleHTMLFileWithStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml")) //load semua file di dalam directory

	p := Person{
		Name:      "aku",
		AddressXX: Address{Street: "haha"},
		Hobbies:   []string{"makan", "minum", "tidur"},
	}

	t.ExecuteTemplate(writer, "simple3.gohtml", p) //name == filename
}

func TestSimpleHTMLFileWithStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFileWithStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
