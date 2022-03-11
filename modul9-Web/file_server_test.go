package modul9_Web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

// http://localhost:8080/static/index.html||.css
func TestFileServer(t *testing.T) {

	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	//mux.Handle("/static", fileServer) // hasil -> /resources/static/index.css
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // http.StripPrefix menghapus prefix /static

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources
var resources embed.FS

// http://localhost:8080/static/index.html||.css
func TestFileServerGolangEmbed(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	//mux.Handle("/static", fileServer)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
