package main

import (
	"fmt"
	"time"
)


func ping(channel chan string) {
	for i := 0; ; i++{
		channel <- "ping"
	}
}

func pong(channel chan string) {
	for i := 0 ;; i++ {
		channel <- "pong"
	}
}

func print(channel chan string){
	for {
		msg := <-channel
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

func main() {
	pingpong := make(chan string)

	go ping(pingpong)
	go print(pingpong)
	go pong(pingpong)
	
	var input string
	fmt.Scanln(&input)

}
