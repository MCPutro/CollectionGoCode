package modul1

import (
	"fmt"
	d "github.com/MCPutro/CollectionGoCode/modul1/helper"
	"testing"
)

type siswaSD struct {
	id, nama string
}

type kelas struct {
	nama        string
	daftarSiswa []siswaSD
}

func (k *kelas) tambahSiswa(id string, nama string) { //method
	k.daftarSiswa = append(k.daftarSiswa, siswaSD{
		id:   id,
		nama: nama,
	})
}

func TestPointerFunc(t *testing.T) {
	kelas5 := kelas{
		nama:        "5",
		daftarSiswa: nil,
	}

	kelas5.tambahSiswa("001", "aku lupa")

	fmt.Println("kelas : ", kelas5.nama)
	fmt.Println("murid :")
	for i, sd := range kelas5.daftarSiswa {
		fmt.Println(i+1, sd.id, sd.nama)
	}

	d.SayHi()

}
