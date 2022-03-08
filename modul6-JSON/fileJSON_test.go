package modul6_JSON

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type customer struct { //atribut harus di awali huruf besar
	Id       string `json:"Id"`
	Username string `json:"user_name"`
	Hehe     string
}

func TestReadFile(t *testing.T) {
	open, _ := os.Open("sample.json")
	decoder := json.NewDecoder(open)

	c := customer{}

	decoder.Decode(&c)

	fmt.Println(c.Username)
}

func TestWriteFile(t *testing.T) {

	create, _ := os.Create("sample.json")
	encoder := json.NewEncoder(create)

	newCustomer := customer{
		Id:       "123",
		Username: "hahah",
	}

	_ = encoder.Encode(newCustomer)

	fmt.Println(newCustomer)
}
