package modul5_embed

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed version.txt
var version string

func Test2String(t *testing.T) {
	fmt.Println(version)
}
