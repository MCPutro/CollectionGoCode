package modul1

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	fmt.Println(time.Now().Local())
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())
	//fmt.Println(time.da)

	wib := time.Date(2021, 11, 9, 18, 11, 10, 0, time.FixedZone("WIB", 7*60*60))
	fmt.Println("wib :", wib)
	fmt.Println("wib :", wib.Local())

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println("utc :", utc)
	fmt.Println("utc :", utc.Local())

	parse, _ := time.Parse(time.RFC3339, "2006-01-02T15:34:05+07:30")
	fmt.Println("parse :", parse)
	fmt.Println("parse :", parse.Local())

	formatDate := "02-Jan-2006" // angka harus sesuai dengan contoh tamplate
	parse2, _ := time.Parse(formatDate, "30-Mar-2023")
	fmt.Println("parse2 :", parse2)
	fmt.Println("parse2 :", parse2.Local())
}
