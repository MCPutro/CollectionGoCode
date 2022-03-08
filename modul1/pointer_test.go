package modul1

import (
	"fmt"
	"testing"
)

type Address struct {
	city, province, country string
}

func coba1() {
	alamat1 := Address{"bandung", "jawa barat", "indonesia"}
	alamat2 := &alamat1 // simbol & mengartikan alamat2 merefer ke lokasi memori alamat1

	alamat2.city = "jakarta"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}

func coba2() {
	alamat1 := Address{"bandung", "jawa barat", "indonesia"}
	alamat2 := &alamat1 // simbol & mengartikan alamat2 merefer ke lokasi memori alamat1

	alamat2 = &Address{"malang", "jawa timur", "indonesia"} //menset variable alamat2 ke memori yg baru yaitu malang sedangkan alamat1 tetap ke bandung

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}

func coba3() {
	alamat1 := Address{"bandung", "jawa barat", "indonesia"}
	alamat2 := &alamat1 // simbol & mengartikan alamat2 merefer ke lokasi memori alamat1

	*alamat2 = Address{"semarang", "jawa tengah", "indonesia"}

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}

func coba4() {
	alamat1 := new(Address) //new digunakan untuk membuat/mengembalikan (return) pointer dengan data kosong (tanpa nilai awal)
	alamat1.city = "surabaya"

	var alamat2 *Address = alamat1
	alamat2.city = "xxxx"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}

func TestPointer(t *testing.T) {
	coba1()
	fmt.Println("======================")
	coba2()
	fmt.Println("======================")
	coba3()
	fmt.Println("======================")
	coba4()
}
