package modul5_embed

import (
	_ "embed"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed Taishiro_Toyomitsu_Hero_Costume.png
var logo []byte

func Test(t *testing.T) {
	err := ioutil.WriteFile("file_baru.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
