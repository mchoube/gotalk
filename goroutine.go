package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // this line will spawn a goroutine and the program execution will continue

	f2()
}

func f1() {
	i := 0

	for {
		fmt.Printf("f1: %d\n", i)
		i++
		time.Sleep(time.Second)
	}
}

func f2() {
	i := 0

	for i < 5 {
		fmt.Printf("f2: %d\n", i)
		i++
		time.Sleep(2 * time.Second)
	}
}
