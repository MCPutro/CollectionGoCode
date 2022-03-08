package modul1

import (
	"fmt"
	"runtime"
	"testing"
)

func printOut(n int, text string) {
	for i := 0; i < n; i++ {
		fmt.Println((i + 1), text)
		//time.Sleep(1)
	}
}

func TestGoroutine(t *testing.T) {
	runtime.GOMAXPROCS(2)

	go printOut(5, "halo")
	printOut(5, "apa kabar")

	var input string
	fmt.Scanln(&input)
}
