package modul1

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexP(t *testing.T) {
	//var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")
	//
	//fmt.Println(regex.MatchString("eko"))
	//fmt.Println(regex.MatchString("eto"))
	//fmt.Println(regex.MatchString("eDo"))
	//
	//search := regex.FindAllString("eko eka edo eto eyo eki", -1) //-1 artinya mencari sebanyak2 nya
	//fmt.Println(search)

	emailPattern := regexp.MustCompile("^([a-zA-Z0-9\\.-\\_]{4,})+@([\\w-]+\\.)+[\\w-]{2,4}$")

	fmt.Println(emailPattern.MatchString("hahaas@gmail.com"))
	fmt.Println(emailPattern.MatchString("_@gmail.com"))

	listEmail := []string{"hahaas@gmail.com", "hahaas@gmail.com", "12@gmail.com"}

	for _, s := range listEmail {
		if emailPattern.MatchString(s) {
			fmt.Println(s)
		}
	}

	xxx := regexp.MustCompile("e([a-z])o")

	search := xxx.FindAllString("eko eka edo eto eyo eki", -1) //-1 artinya mencari sebanyak2 nya
	fmt.Println(search)
}
