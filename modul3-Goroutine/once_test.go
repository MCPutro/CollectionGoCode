package modul3_Goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func addCounter() {
	counter++
}

func Test_once(t *testing.T) {

	//hanya 1x panggil function,
	//jika ada banyak goroutine yg memanggil func yg sama, maka hanya goroutine pertama yg bisa meng-eksekusi function tersebut
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		func() {
			group.Add(1)
			once.Do(addCounter) //hanya bisa untuk func tanpa parameter
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(counter)
}
