package modul1

import (
	"container/ring"
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	data := ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = fmt.Sprint("data ke ", i)
		data = data.Next()
	}

	data.Do(func(i interface{}) {
		fmt.Println(i)
	})

}
