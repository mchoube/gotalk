package main

import "some_lib"

func nonBlocking(doneCh chan bool) {
	some_lib.blockingCall()
	doneCh <- true
}

func nonBlockingWithTimeout(done chan bool, timeout time.Duration) {
	doneCh := make(chan bool)
	timer := time.After(timeout) // HL
	go nonblocking(doneCh)
	select {
	case <-doneCh:
		done <- true
		return
	case <-timer:
		done <- true
		return
	}
}

func main() {
	// do some work

	ch := make(chan bool)
	go nonBlockingWithTimeout(ch, time.Second) // HL

	// do some other work

	// wait for the blocking call
	<-ch
}
