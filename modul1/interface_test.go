package modul1

import (
	"fmt"
	"testing"
)

type orang interface {
	printNama()
	getNama() string
}

type siswa struct {
	nism int
	nama string
}

func (s siswa) printNama() {
	//TODO implement me
	fmt.Println("nama saya", s.nama)
}

func (s siswa) getNama() string {
	//TODO implement me
	return s.nama
}

func TestInterface(t *testing.T) {
	A := siswa{
		nism: 1,
		nama: "lala",
	}

	fmt.Println(A.getNama())
	A.printNama()
}
