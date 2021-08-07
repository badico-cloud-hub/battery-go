package main

import (
	"fmt"
)

func main() {
	quit := make(chan struct{})
	dispatch := make(chan string)
	go timeInterval(
		timeoutIntervalConfig{
			seconds: 1,
			event:   "1 second",
		},
		dispatch,
		quit,
	)

	go timeInterval(
		timeoutIntervalConfig{
			seconds: 5,
			event:   "5 seconds",
		},
		dispatch,
		quit,
	)

	for action := range dispatch {
		fmt.Println(action)
	}
	// time.Sleep(60 * 10 * time.Second)
	fmt.Println("finish!")
}
