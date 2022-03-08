package modul1

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStrconvFunc(t *testing.T) {
	/**
	- base :
		biner 2
		oktal 8
		decimal 10
		hexadecimal 16
	 mengacu pada data yg akan di convert
	- bitSize : tujuan akhir
	**/
	var str = "12d680"
	var num, err = strconv.ParseInt(str, 16, 32)

	if err == nil {
		fmt.Println(num) // 10
	} else {
		fmt.Println("error :", err)
	}

	/**base adalah tujuan akhir (convert to) **/
	hasil := strconv.FormatInt(1234560, 16)
	fmt.Println(hasil)

}
