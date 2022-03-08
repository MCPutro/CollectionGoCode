package modul1

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCoba(t *testing.T) {

	/**v32 := "-354634382"
	s1, err1 := strconv.ParseInt(v32, 10, 32)
	if err1 == nil {
		fmt.Printf("%T, %v\n", s1, s1)
		fmt.Println("--------------------------------------1")
	} else {
		fmt.Println("error :", err1)
	}

	s2, err2 := strconv.ParseInt(v32, , 32)
	if err2 == nil {
		fmt.Printf("%T, %v\n", s2, s2)
		fmt.Println("--------------------------------------2")
	} else {
		fmt.Println("error :", err2)
	}

	fmt.Println("===========================================================")
	v64 := "-3546343826724305832"
	if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
		fmt.Println("--------------------------------------3")
	}
	if s, err := strconv.ParseInt(v64, 16, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
		fmt.Println("--------------------------------------4")
	}**/
	var str = "100"
	var num, err = strconv.ParseInt(str, 16, 16)

	if err == nil {
		fmt.Println(num) // 10
	} else {
		fmt.Println("error :", err)
	}
}
