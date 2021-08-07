package batterygo

import (
	"time"

	"github.com/badico-cloud-hub/battery-go/storages"
)

const RELOAD_EVENT = "reload"

type Battery struct {
	storage  storages.Storage
	interval time.Duration
}

type BatteryArgument struct {
	Key   string
	Value interface{}
}

func NewBattery(storage storages.Storage, interval time.Duration) *Battery {
	return &Battery{
		storage, interval,
	}
}

func (b *Battery) Init(reloadStorage func() []BatteryArgument) {
	quit := make(chan struct{})
	dispatch := make(chan string)
	go timeInterval(
		timeoutIntervalConfig{
			seconds: b.interval,
			event:   RELOAD_EVENT,
		},
		dispatch,
		quit,
	)

	argsFirstLoad := reloadStorage()
	for _, arg := range argsFirstLoad {
		b.storage.Set(arg.Key, arg.Value)
	}
	for action := range dispatch {
		if action == RELOAD_EVENT {
			argReload := reloadStorage()
			for _, arg := range argReload {
				b.storage.Set(arg.Key, arg.Value)
			}
		}
	}
}
