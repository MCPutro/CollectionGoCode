package modul1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestHttp1(t *testing.T) {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	fmt.Println("-----------------------------------------")
	fmt.Println(resp.Body)
	fmt.Println("-----------------------------------------")

	/**cara 1**/
	/**bs := make([]byte, 99999) //alokasi memori
	resp.Body.Read(bs)        //mamasukkan body ke memori yang sudah di alokasikan
	fmt.Println(string(bs))**/

	/**cara 2**/
	//io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

type logWriter struct{}

func (l logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return 1, nil
}
