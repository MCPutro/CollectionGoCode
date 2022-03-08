package modul1

import (
	"fmt"
	"reflect"
	"testing"
)

type Data interface{}

type response struct {
	status, message string
	data            Data
}

type userRespose struct {
	nama string
	id   int
}

func TestInterfaceKosong(t *testing.T) {
	ur := userRespose{
		nama: "user1",
		id:   1,
	}

	resp := response{
		status:  "success",
		message: "",
		data:    ur,
	}

	fmt.Println(resp)

	fmt.Println(resp.data)

	x := resp.data

	fmt.Println(reflect.TypeOf(x) == reflect.TypeOf(userRespose{}))

	var backUpUser userRespose

	switch v := x.(type) {
	case userRespose:
		fmt.Println("ini data user", v.nama)
		backUpUser = v
	case int:
		fmt.Println("ini data integer")
	default:
		fmt.Println("kagak tau om")
	}

	fmt.Println("---------------------")
	fmt.Println(backUpUser.nama)
}
