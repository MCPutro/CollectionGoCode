package modul3_Goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {

	timer := time.NewTimer(5 * time.Second)
	fmt.Println("1>", time.Now())

	time2 := <-timer.C
	fmt.Println("2>", time2)
}

func TestTimerAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println("1>", time.Now())

	time2 := <-channel
	fmt.Println("2>", time2)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("2>", time.Now())
		group.Done()
	})
	fmt.Println("1>", time.Now())

	group.Wait()

}
