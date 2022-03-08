package modul1

import (
	"flag"
	"fmt"
	"testing"
)

func TestFlag(t *testing.T) {
	//variable name, default value if null, desc
	hostname := flag.String("host", "localhost", "Put your database host")
	username := flag.String("user", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	//flag.String("host", "localhost", "Put your database host")

	flag.Parse()

	fmt.Println(*hostname)
	fmt.Println(*username)
	fmt.Println(*password)

}
