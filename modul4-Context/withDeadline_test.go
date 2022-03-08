package modul4_Context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func createCounter3(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				fmt.Println("di time out")
				return
			default:
				destination <- counter //memasukkan counter ke dalam channel(destination)
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestWithDeadline(t *testing.T) {
	fmt.Println("Total Gourotine:", runtime.NumGoroutine())

	parentContext := context.Background()
	ctx, cancel := context.WithDeadline(parentContext, time.Now().Add(5*time.Second))
	defer cancel()

	destination := createCounter3(ctx)

	for i := range destination {
		fmt.Println("Counter", i)
		if i == 10 {
			break
		}
	}

	time.Sleep(3 * time.Second)

	fmt.Println("Total Gourotine:", runtime.NumGoroutine())
}
