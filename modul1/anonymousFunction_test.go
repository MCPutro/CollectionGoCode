package modul1

import (
	"fmt"
	"testing"
)

func testData(who string, isValid func(string) bool) {
	if isValid(who) {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}
}

func TestAnonymousFunc(t *testing.T) {
	//di looping
	for x, each := range []string{"test A", "test B", "test C"} {
		//anonymous function
		func(who string, n int) {
			var data = fmt.Sprintf("hello %s", who)
			fmt.Println(n, data)
		}(each, x)
	}

	//di variable
	validate := func(name string) bool {
		return name == "admin"
	}

	fmt.Println(validate("admin"))

	testData("admin", func(s string) bool {
		return s == "admin"
	})

}
