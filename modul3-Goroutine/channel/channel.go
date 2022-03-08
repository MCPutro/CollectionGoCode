package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	/**var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")
	sayHelloTo("xxx xxx")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)**/

	names := []string{"haha", "hihi", "hohoh"}

	printDeh := func(who string) {
		say := fmt.Sprintf("hi %s", who)
		messages <- say //insert to channel
	}

	for _, name := range names {
		go printDeh(name)
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(<-messages)
	}

}
