package modul1

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

var delimiter string = "; "

func save2file(fileName string, d []string) error {
	return ioutil.WriteFile(fileName, []byte(strings.Join(d, delimiter)), 0666)
}

func loadFromFile(fileName string) []string {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), delimiter)
	return s
}

func TestFile1(t *testing.T) {
	filename := "fileData.txt"

	data := []string{"satu", "dua", "tiga"}

	err := save2file(filename, data)
	if err != nil {
		return
	}

	file := loadFromFile(filename)
	fmt.Println(file)
	fmt.Println(len(file))

}

func TestFile2(t *testing.T) {

	//write to file
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("fileDat1.txt", d1, 0644)
	if err != nil {
	}

	//append
	f, err := os.OpenFile("fileDat1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
	}
	defer f.Close()
	if _, err := f.WriteString("text to append1\n"); err != nil {
		log.Println(err)
	}
}
