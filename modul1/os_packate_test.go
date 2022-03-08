package modul1

import (
	"fmt"
	"os"
	"testing"
)

func TestPackageOS(t *testing.T) {
	//go run os_packate_test.go [param1] [param2]

	args := os.Args

	fmt.Println(args)
	fmt.Println(args[1])

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("error :", err)
	}
	fmt.Println(hostname)

}
