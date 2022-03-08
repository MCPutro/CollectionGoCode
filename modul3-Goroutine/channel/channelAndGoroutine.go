package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{"http://google.com", "http://facebook.com", "http://golang.org"}

	channel := make(chan string)

	for _, link := range links {
		go check(link, channel)
	}

	//for i := 0; i < len(links); i++ {
	//	fmt.Println(<-c) //out
	//}

	for s := range channel {
		fmt.Println(s) //apa(s)
	}

	//for {
	//	check(<-c, c)
	//}

	//var input string
	//fmt.Scanln(&input)

}

func check(link string, c chan string) { //channel sebagai parameter
	_, err := http.Get(link)
	if err != nil {
		//fmt.Println(link, "is down")
		c <- link //fmt.Sprintf("%s is down", link) //in
		return
	}

	//fmt.Println(link, "is up")
	c <- link //fmt.Sprintf("%s is up", link)
}

func apa(test string) {
	fmt.Println(test)
}
