package modul1

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	data := list.New()

	data.PushBack("2")
	data.PushBack("1")
	data.PushFront("7")

	data1 := list.New()
	data1.PushFront("a")
	data1.PushFront("b")
	data1.PushFront("c")

	data.PushBackList(data1)

	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
