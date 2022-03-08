package modul4_Context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func createCounter1(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				fmt.Println("di cancel")
				return
			default:
				destination <- counter //memasukkan counter ke dalam channel(destination)
				counter++
			}
		}
	}()

	return destination
}

func TestWithCancel(t *testing.T) {
	fmt.Println("Total Gourotine:", runtime.NumGoroutine())

	parentContext := context.Background()
	ctx, cancel := context.WithCancel(parentContext)

	destination := createCounter1(ctx)

	for i := range destination {
		fmt.Println("Counter", i)
		if i == 10 {
			break
		}
	}

	cancel()

	time.Sleep(3 * time.Second)

	fmt.Println("Total Gourotine:", runtime.NumGoroutine())
}
