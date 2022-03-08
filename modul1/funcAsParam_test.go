package modul1

import (
	"fmt"
	"testing"
)

//1
type Filter func(string2 string) string

func cek2(who string, filter Filter) string {
	return fmt.Sprint("hi ", filter(who))
}

//2
func cek(who string, filte func(string) string) string {
	return fmt.Sprint("hi, ", filte(who))
}

func spamFilter(who2 string) string {
	if who2 == "anjing" {
		return "..."
	}

	return who2
}

func TestFunctionAsParam(t *testing.T) {
	//1.1
	fmt.Println(cek("barbar", spamFilter))
	//1.2
	filter := spamFilter
	fmt.Println(cek("anjing", filter))

	//2.1
	fmt.Println(cek2("okok", spamFilter))
}
