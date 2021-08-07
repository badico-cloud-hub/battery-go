package main

import (
	"fmt"
	"math/rand"

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

func updateBatteryStorage() []BatteryArgument {
	key := "foo"
	value := randSeq(10)
	fmt.Println("update storage: ", value)

	b := BatteryArgument{
		key,
		value,
	}

	return []BatteryArgument{b}
}

func main() {
	storage := storages.New()
	battery := NewBattery(storage, 3)
	go battery.Init(updateBatteryStorage)
	for {
		value, err := storage.Get("init")
		if err != nil {
			fmt.Println("err", err)
			return
		}

		foo, errf := storage.Get("foo")
		if errf != nil {
			fmt.Println("err", errf)
			// return
		}
		fmt.Println("running", value, foo)
	}
}
