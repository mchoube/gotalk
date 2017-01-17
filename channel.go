package main

import (
	"fmt"
	"time"
)

func f1(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "done." // sending data over channel
}

func main() {
	ch := make(chan string)
	go f1(ch)
	fmt.Println("Waiting for goroutine...")
	// this will block until data is available on channel and then read the data
	fmt.Println(<-ch)
}
