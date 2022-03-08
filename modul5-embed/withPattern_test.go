package modul5_embed

import (
	"embed"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed files/fil*.txt
var filess embed.FS

func TestFiless(t *testing.T) {
	dir, err := filess.ReadDir("files")
	if err != nil {
		panic(err)
	}

	for _, entry := range dir {
		fmt.Println(entry)
		fmt.Println(entry.Name())
		fmt.Println(entry.IsDir())
		file, _ := filess.ReadFile("files/" + entry.Name())
		fmt.Println(string(file))
		fmt.Println("------------")
	}
}
