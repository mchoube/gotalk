package main

import "some_lib"

func nonBlocking(doneCh chan bool) {
	some_lib.blockingCall()
	doneCh <- true
}

func main() {
	// do some work

	ch := make(chan bool)
	go nonBlocking(ch) // HL

	// do some other work

	// wait for the blocking call
	<-ch
}
