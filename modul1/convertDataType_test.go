package modul1

import (
	"fmt"
	"testing"
)

func TestConvertData(t *testing.T) {
	var nilai32 int32 = 10000
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	name := "abcdefghijk"

	for i := 0; i < len(name); i++ {
		fmt.Println(">", string(name[i]))
	}

	type noKTP string

	var haha noKTP = "kokok"
	fmt.Println(haha)
}
