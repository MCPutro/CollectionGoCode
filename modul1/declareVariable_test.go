package modul1

import (
	"fmt"
	"testing"
)

func TestCreateVariable(t *testing.T) {
	//1
	var data string
	data = "haha"
	fmt.Println(data)

	//2
	var data2 = "koko"
	fmt.Println(data2)

	//3
	data3 := "okelah"
	fmt.Println(data3)

	//4 constant hanya bisa di assign 1x
	const name = "data3"
	const name2 string = "ad"
	fmt.Println(name)
	fmt.Println(name2)

	//5 constant hanya bisa di assign 1x
	const (
		a        = "qqq"
		b string = "aaa"
	)
	fmt.Println(a)
	fmt.Println(b)

}
