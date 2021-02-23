package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	go doSomething(done)

	duration := 50 * time.Millisecond
	// do not use time.After in select case
	// to avoid OOM
	timer := time.NewTimer(duration)
	defer timer.Stop()

LOOP:
	for {
		timer.Reset(duration)
		select {
		case <-done:
			fmt.Println("task done!")
			break LOOP
		case <-timer.C:
			fmt.Println("times up!")
			break LOOP
		}
	}
}

func doSomething(done chan<- struct{}) {
	// do nothing
	time.Sleep(20 * time.Millisecond)
	done <- struct{}{}
}
