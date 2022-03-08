package modul1

import (
	"fmt"
	"testing"
)

func test(datas ...int) {
	for _, data := range datas {
		fmt.Println("data = ", data)
	}
}

func TestVariadicFunc(t *testing.T) {
	test(1, 2, 4, 3)

	dataSlice := []int{1, 2, 3, 4}
	test(dataSlice...)
}
