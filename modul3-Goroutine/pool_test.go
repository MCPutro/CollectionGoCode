package modul3_Goroutine

import (
	"fmt"
	"sync"
	"testing"
)

//design pattern pool object

func Test_pool(t *testing.T) {
	group := &sync.WaitGroup{}
	pool := sync.Pool{
		New: func() interface{} {
			return "kosong"
		},
	}

	pool.Put("satu")
	pool.Put("2")
	pool.Put("3")

	for i := 0; i < 10; i++ {
		go func() {
			defer group.Done()
			group.Add(1)

			data := pool.Get()
			fmt.Println(">", data)
			pool.Put(data)

		}()
	}

	group.Wait()
	//time.Sleep(11 * time.Second)
	fmt.Println("selesai")
}

func Test_aja(t *testing.T) {
	defer fmt.Println("defer 1")

	func() {
		defer fmt.Println("defer 3")
		fmt.Println("defer 4")
	}()

	fmt.Println("defer 2")
}
