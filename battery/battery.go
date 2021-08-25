package batterygo

import (
	"time"

	"github.com/badico-cloud-hub/battery-go/storages"
)

const RELOAD_EVENT = "reload"

type Battery struct {
	storage     storages.Storage
	interval    time.Duration
	Dispatch    chan []BatteryArgument
	quitTimer   chan struct{}
	stopBattery chan struct{}
}

type BatteryArgument struct {
	Key   string
	Value interface{}
}

func NewBattery(storage storages.Storage, interval time.Duration) *Battery {
	Dispatch := make(chan []BatteryArgument)
	quitTimer := make(chan struct{})
	stopBattery := make(chan struct{})

	return &Battery{
		storage,
		interval,
		Dispatch,
		quitTimer,
		stopBattery,
	}
}

func (b *Battery) Get(key string) (interface{}, error) {
	value, err := b.storage.Get(key)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (b *Battery) Init(reloadStorage func() []BatteryArgument) {
	timeout := make(chan string)
	go timeInterval(
		TimeoutIntervalConfig{
			seconds: b.interval,
			event:   RELOAD_EVENT,
		},
		timeout,
		b.quitTimer,
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
		case <-b.stopBattery:
			return
		}
	}
}

func (b *Battery) Stop() {
	b.quitTimer <- struct{}{}
	b.stopBattery <- struct{}{}
}
