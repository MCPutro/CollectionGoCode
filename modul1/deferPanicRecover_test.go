package modul1

import (
	"fmt"
	"testing"
)

func satu() {
	fmt.Println("satu")

	msg := recover() //dipasang di defer

	fmt.Println("error dengan pesan : ", msg)
}

func dua() {
	defer satu() //defer akan slalu di eksekusi di akhir walaupun ada error logic
	fmt.Println("dua")
	panic("error gan") //aplikasi berhenti

	fmt.Println("xxxx") //jika ada panic maka code di bawahnya tidak akan di eksekusi
}

func TestDeferPanicRecovery(t *testing.T) {
	dua()
	fmt.Println("aaaaaaaaaaaaaaa")
}
