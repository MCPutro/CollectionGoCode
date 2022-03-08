package modul5_embed

import (
	"embed"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed files/file1.txt
//go:embed files/file2.txt
//go:embed files/file3.txt
var files embed.FS

func TestDe(t *testing.T) {
	file1, err1 := files.ReadFile("files/file1.txt")
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(string(file1))

	file2, err2 := files.ReadFile("files/file2.txt")
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(string(file2))

	file3, err3 := files.ReadFile("files/file3.txt")
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(string(file3))
}
