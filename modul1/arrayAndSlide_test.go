package modul1

import (
	"fmt"
	"testing"
)

func TestArrayAndSlice(t *testing.T) {
	dataArrayA := [...]string{"haha", "hhoho"} //ukuran array di hitung berdasarkan data elemen yang diisikan, contoh di smaping maka ukuran adalah 2

	fmt.Println(dataArrayA)

	dataArray := [5]string{}

	for i := 0; i < len(dataArray); i++ {
		dataArray[i] = fmt.Sprint("data array ke-", i)
	}

	for _, s := range dataArray {
		fmt.Println(s)
	}

	dataSlise := []string{}

	for i := 0; i < 7; i++ {
		//append(dataSlise, "asd")

		dataSlise = append(dataSlise, fmt.Sprint("data slide ke", i))
	}

	fmt.Println(">", cap(dataSlise))
}
