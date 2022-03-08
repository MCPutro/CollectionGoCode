package modul3_Goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func RunHello(name string) {
	fmt.Println("hello", name)
}

func RunHello2(chann chan string, name string) {
	time.Sleep(2 * time.Second)
	chann <- "hello," + name
}

func Test_pertama(t *testing.T) {
	go RunHello("aku")
	fmt.Println("kkk")

	time.Sleep(2 * time.Second)
}

func Test_useChannel(t *testing.T) {
	// tipe data channel harus ada tipe ny, ex: int, string, struct
	// data di channel harus ada yg mengambil

	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "test"
	}()

	data := <-channel
	fmt.Println(data)
	//close(channel)

	time.Sleep(5 * time.Second)
}

func Test_ChannelAsParam(t *testing.T) {
	channel := make(chan string)

	go RunHello2(channel, "au ah")

	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)

}

func onlyIn(chann chan<- string) {
	chann <- "haha"
}

func onlyOut(chann <-chan string) {
	data := <-chann
	fmt.Println(data)
}

func Test_ChannelInOut(t *testing.T) {

	channel := make(chan string)
	defer close(channel)

	go onlyIn(channel)
	go onlyOut(channel)

	time.Sleep(5 * time.Second)

}

func Test_BufferChannel(t *testing.T) {
	// 3 adalah buffer nya,
	// saat channel memiliki buffer, tidak ada ygmengmabil data dari hannel jg gpp
	// akan error jika di insert/mengmabil lebih dari 3 di channel
	chann := make(chan string, 3)

	chann <- "okok"
	chann <- "okok1"
	chann <- "okok2"

	fmt.Println(len(chann))
	fmt.Println(cap(chann))
	fmt.Println(<-chann)
	fmt.Println(<-chann)
	fmt.Println(<-chann)

}

func Test_RangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i <= 3; i++ {
			channel <- "looping ke-" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
}

func Test_range2(t *testing.T) {
	var messages = make(chan string)
	defer close(messages)
	names := []string{"haha", "hihi", "hohoh"}

	printDeh := func(who string) {
		say := fmt.Sprintf("hi %s", who)
		messages <- say //insert to channel
	}

	for _, name := range names {
		go printDeh(name)
	}

	for range names {
		fmt.Println("> ", <-messages)
	}

}

func Test_selectInChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go RunHello2(channel1, "pertama")
	go RunHello2(channel2, "kedua")

	i := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println(data)
			i++
		case data := <-channel2:
			fmt.Println(data)
			i++
		default:
			fmt.Println("tunggu dulu")
		}

		if i == 2 {
			break
		}
	}
}
