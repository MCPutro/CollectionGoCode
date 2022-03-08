package modul3_Goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

//manipulasi data primitive di goroutine, contoh goroutine manipulasi counter++
//tp jika struct, lebih baik dengan mutex

func TestAtomic(t *testing.T) {
	var x int32 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				atomic.AddInt32(&x, 1)
				//x = x + 1
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(x)

}
