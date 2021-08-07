package main

import (
	"fmt"
	"time"

	"github.com/badico-cloud-hub/battery-go/storages"
)

func main() {
	quit := make(chan struct{})
	dispatch := make(chan string)
	myStorage := storages.New()
	go timeInterval(
		timeoutIntervalConfig{
			seconds: 2,
			event:   "A",
		},
		dispatch,
		quit,
	)
	go timeInterval(
		timeoutIntervalConfig{
			seconds: 1,
			event:   "read",
		},
		dispatch,
		quit,
	)

	go timeInterval(
		timeoutIntervalConfig{
			seconds: 8,
			event:   "B",
		},
		dispatch,
		quit,
	)

	go timeInterval(
		timeoutIntervalConfig{
			seconds: 40,
			event:   "C",
		},
		dispatch,
		quit,
	)
	var i int
	for action := range dispatch {
		switch action {
		case "A":
			fmt.Println("A")
			myStorage.Set("A", i)
		case "B":
			fmt.Println("B")
			myStorage.Set("B", i)
		case "read":
			fmt.Println("read", myStorage)
		default:
			fmt.Println("end", myStorage)
			quit <- struct{}{}
			quit <- struct{}{}
			quit <- struct{}{}
			time.Sleep(30 * time.Second)
			fmt.Println("stoping")
			return
		}
		i = i + 1
	}
	fmt.Println("finish!")
}
