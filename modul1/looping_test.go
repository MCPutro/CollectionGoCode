package modul1

import (
	"fmt"
	"testing"
)

func TestLooping(t *testing.T) {
	fmt.Println("--------------------------------------------------- 0")

	datas := []string{"apple", "grape", "banana", "melon"} //slice
	for _, data := range datas {
		fmt.Print(data, " ")
	}
	fmt.Println()

	fmt.Println("--------------------------------------------------- 1")

	buah := [3]string{"apple", "grape", "banana"}
	for i := 0; i < len(buah); i++ {
		fmt.Println(i, buah[i])
	}

	fmt.Println("--------------------------------------------------- 2")

	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	fmt.Println("--------------------------------------------------- 3")

	for _, each := range []string{"wick", "hunt", "bourne"} {
		fmt.Println(each)
	}

	fmt.Println("--------------------------------------------------- 4")

	for x, each := range []string{"test A", "test B", "test C"} {
		func(who string, n int) {
			var data = fmt.Sprintf("hello %s", who)
			fmt.Println(n, data)
		}(each, x)
	}

	fmt.Println("--------------------------------------------------- 5")

	counter := 1
	for counter <= 10 {
		fmt.Println("loop", counter)
		counter++
	}

}
