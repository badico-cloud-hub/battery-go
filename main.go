package main

import (
	"fmt"
	"math/rand"

	batterygo "github.com/badico-cloud-hub/battery-go/battery"
	"github.com/badico-cloud-hub/battery-go/storages"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func updateBatteryStorage() []batterygo.BatteryArgument {
	Key := "foo"
	Value := randSeq(10)
	fmt.Println("update storage: ", Value)

	b := batterygo.BatteryArgument{
		Key,
		Value,
	}

	return []batterygo.BatteryArgument{b}
}

func main() {
	storage, err := storages.NewRedisStorage()
	if (err != nil) {
		fmt.Println("err", err)
		return
	}
	battery := batterygo.NewBattery(storage, 3)
	go battery.Init(updateBatteryStorage)
	for {
		// value, err := storage.Get("init")
		// if err != nil {
		// 	fmt.Println("err", err)
		// 	return
		// }

		foo, errf := storage.Get("foo")
		if errf != nil {
			fmt.Println("err", errf)
			// return
		} else {
			fmt.Println("running", foo)
		}
		fmt.Println("running")
	}
}
