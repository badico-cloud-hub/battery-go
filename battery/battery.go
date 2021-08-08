package batterygo

import (
	"time"

	"github.com/badico-cloud-hub/battery-go/storages"
)

const RELOAD_EVENT = "reload"

type Battery struct {
	storage  storages.Storage
	interval time.Duration
	Dispatch chan []BatteryArgument
}

type BatteryArgument struct {
	Key   string
	Value interface{}
}

func NewBattery(storage storages.Storage, interval time.Duration) *Battery {
	Dispatch := make(chan []BatteryArgument)
	return &Battery{
		storage,
		interval,
		Dispatch,
	}
}

func (b *Battery) Get(key string) interface{} {
	value, err := b.storage.Get(key)
	if err != nil {
		return err
	}
	return value
}
func (b *Battery) Init(reloadStorage func() []BatteryArgument) chan []BatteryArgument {
	quit := make(chan struct{})
	timeout := make(chan string)
	go timeInterval(
		timeoutIntervalConfig{
			seconds: b.interval,
			event:   RELOAD_EVENT,
		},
		timeout,
		quit,
	)

	argsFirstLoad := reloadStorage()
	for _, arg := range argsFirstLoad {
		b.storage.Set(arg.Key, arg.Value)
	}
	for {
		select {
		case action := <-timeout:
			if action == RELOAD_EVENT {
				argReload := reloadStorage()
				for _, arg := range argReload {
					b.storage.Set(arg.Key, arg.Value)
				}
			}
		case args := <-b.Dispatch:
			for _, arg := range args {
				b.storage.Set(arg.Key, arg.Value)
			}
		}
	}

}
