package modul1

import (
	"fmt"
	"testing"
)

func say(who string) string {
	return fmt.Sprint("hi, ", who)
}

func TestFuncAsVariable(t *testing.T) {
	param := say //menyimpan func ke dalam variable

	fmt.Println(param("mcp"))

}
