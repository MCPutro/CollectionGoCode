package modul1

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	//1
	name := "HAHA"
	switch name {
	case "haha":
		fmt.Println("opsi 1")
	case "ok":
		fmt.Println("opsi 2")
	default:
		fmt.Println("opsi 3")
	}

	//2
	switch panjang := 4; panjang >= 9 {
	case true:
		fmt.Println("benar")
	case false:
		fmt.Println("salah")
	}

	//3
	name1 := "okelah11"
	panjang1 := -5
	switch {
	case name1 == "okelah":
		fmt.Println("aaaa")
	case panjang1 > 1:
		fmt.Println("bbbb")
	default:
		fmt.Println("default")
	}
}
