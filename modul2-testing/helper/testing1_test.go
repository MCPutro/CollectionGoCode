package helper

import (
	"fmt"
	modul2testing "github.com/MCPutro/CollectionGoCode/modul2-testing"
	"testing"
)

/**untuk testing ada beberapa aturan:
	1. nama file harus sama dengan file yg akan di test & di tambahakn "_test"
	2. nama func harus diawali "Test", nama func tidak harus sama, tp untuk memudahakan bisa disamakan
	3. harus ada param t *testing.T & no return
**/

func TestSayhi(t *testing.T) {
	tes := modul2testing.SayHi("haha")

	if tes != "hi,haha" {
		//panic("error om") /**tidak baik digunakan saat testing krn program langsung berhenti**/
		t.Fail() //saat gagal testing, code berikut nya tetap di eksekusi, unit slanjutnya tetap benlanjut
	}

	fmt.Println("testing TestSayhi selese")
}

func TestSayhi1(t *testing.T) {
	tes := modul2testing.SayHi("haha")

	if tes != "hi,haha1" {
		//panic("error om") /**tidak baik digunakan saat testing krn program langsung berhenti**/
		t.FailNow() //saat gagal testing, code berikut tidak di eksekusi, unit slanjutnya tetap benlanjut
	}
	fmt.Println("testing TestSayhi1 selese")
}

func TestSayhi2(t *testing.T) {
	tes := modul2testing.SayHi("haha")

	if tes != "hi,haha1" {
		t.Error("harus nya hi,haha") //setelah dieksekusi akan otomatis langsung memanggil t.Fail
	}
	fmt.Println("testing TestSayhi2 selese")
}

func TestSayhi3(t *testing.T) {
	tes := modul2testing.SayHi("haha")

	if tes != "hi,haha1" {
		t.Fatal("harus nya hi,haha") //setelah dieksekusi akan otomatis langsung memanggil t.FailNow
	}
	fmt.Println("testing TestSayhi3 selese")
}
