package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	say("First!") // say runs... we wait

	// A goroutine is a lightweight thread managed by the Go runtime.
	go say("world") // say starts running
	say("hello") // does not wait for second say to finish
/*
	Goroutines run in the same address space, so access to shared memory must be synchronized.
	The sync package provides useful primitives, although you won't need them much in Go as there are other primitives.
*/
}
