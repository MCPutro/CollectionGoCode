package modul1

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	//map[key-type]value-type //key is uniq
	person := map[string]string{
		"id1": "11111",
		"id3": "3333",
		"id2": "222",
	}

	person["id4"] = "ookkk"

	fmt.Println(person)
	fmt.Println(person["id2"]) // print value

	for key, value := range person {
		fmt.Println(key, "-", value)
	}

	color2 := map[string]string{}
	color2["color2"] = "hitam"

	fmt.Println(color2)
}
