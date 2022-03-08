package modul3_Goroutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestTicker1(t *testing.T) {
	ticker := time.NewTicker(2 * time.Second)

	x := 0

	go func() {
		if x == 5 {
			ticker.Stop()
		}
	}()

	for tick := range ticker.C {
		fmt.Println(tick)
		x++
	}

}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}

func TestCore(t *testing.T) {
	cpu := runtime.NumCPU()
	fmt.Println(cpu)

	thread := runtime.GOMAXPROCS(-1)
	fmt.Println(thread)
}
