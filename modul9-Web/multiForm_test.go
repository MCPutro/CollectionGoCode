package modul9_Web

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestServer(t *testing.T) {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9999"
	}
	fmt.Println("server running in port ", PORT)

	myRoute := mux.NewRouter()

	myRoute.HandleFunc("/form2", multiPartForm).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+PORT, myRoute))
}

func multiPartForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err := r.ParseMultipartForm(32 << 20) // 32Mb
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	alias := r.FormValue("alias")

	uploadedFile, handler, err := r.FormFile("file")

	//fmt.Println(uploadedFile)
	//fmt.Println(handler)
	//fmt.Println(alias)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd() //mendapatkan path dari file.go yg di exec
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := handler.Filename
	if alias != "" {
		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
	}

	path := filepath.Join(dir, "files")
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	fileLocation := filepath.Join(path, filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("done"))
}
