package modul1

import (
	"fmt"
	"testing"
)

type user struct {
	id, name string
	umur     int
}

func (u user) printOut() {
	fmt.Println("name : ", u.name)
	fmt.Println("id : ", u.id)
	fmt.Println("umur : ", u.umur)

}

func TestStruct(t *testing.T) {
	saya := user{
		id:   "A1",
		name: "nama1",
		umur: 1,
	}

	fmt.Println(saya)

	var saya1 user
	saya1.id = "A2"
	saya1.name = "nama2"
	saya1.umur = 5

	fmt.Println(saya1)

	fmt.Println("-----------------")
	saya.printOut()

}
