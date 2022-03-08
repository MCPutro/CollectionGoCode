package modul9_Web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//import template
//{{template "name"}} untuk meng-import template tanpa mengirim data apapun ke template yg diimport
//{{template "name" .value}} untuk meng-import template dg mengirim data ke template yg diimport

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
		"./templates/body.gohtml",
	))

	data := map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "okok",
	}

	t.ExecuteTemplate(writer, "okok", data)
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
