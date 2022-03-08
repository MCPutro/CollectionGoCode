package modul3_Goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsyn(group *sync.WaitGroup, i int) {
	group.Add(1)
	defer group.Done()

	fmt.Println("hello", i)
	time.Sleep(1 * time.Second)
}

func Test_wait_group(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyn(group, i)
	}

	group.Wait()
	fmt.Println("selesai")
}

func Test_wait_group2(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func(x int) {
			group.Add(1)
			defer group.Done()

			fmt.Println("hello2", x)
			//time.Sleep(1 * time.Second)
		}(i)
	}

	group.Wait()
	fmt.Println("selesai")
}
