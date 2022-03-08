package assertion_example

import (
	"fmt"
	modul2testing "github.com/MCPutro/CollectionGoCode/modul2-testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssertion1(t *testing.T) {
	result := modul2testing.SayHi("a ku")

	assert.Equal(t, "hi,aku", result, "ini salah") // otomatis memanggil Fail

	assert.Equal(t, "hi,aku", result)

	fmt.Println("->>>>>>>> done")
}

func TestAssertion2(t *testing.T) {

	SD := modul2testing.School{}
	SD.TambahMahasiswa()

	fmt.Println("->>>>>>>> done")
}
