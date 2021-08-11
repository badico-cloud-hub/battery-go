package batterygo_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/badico-cloud-hub/battery-go/battery"
	"github.com/badico-cloud-hub/battery-go/storages"
)

func setupBatteryArguments() []BatteryArgument {
	max := 10

	var args []BatteryArgument

	for i := 0; i < max; i++ {
		args = append(args, BatteryArgument{
			Key:   fmt.Sprintf("someTestingKey%d", i),
			Value: fmt.Sprintf("someTestingValue%d", i),
		})
	}

	return args
}

func expectStorageToHasBatteryArgumentsMock(t *testing.T, storage storages.Storage, batteryArgumentsMock []BatteryArgument) {
	for _, arg := range batteryArgumentsMock {
		value, err := storage.Get(arg.Key)
		if err != nil {
			t.Error(err)
		}

		if value != arg.Value {
			t.Errorf("wanted to %s key %s got %s", arg.Value, arg.Key, value.(string))
		}
	}
}

func TestBattery(t *testing.T) {

	batteryArgumentsMock := setupBatteryArguments()

	t.Run("should call ReloadPage 5 times", func(t *testing.T) {
		fixedInterval := 1
		storage := storages.NewGoCacheStorage()
		battery := NewBattery(storage, time.Duration(fixedInterval))

		wantedNumCalls := 5

		var gotNumCalls int

		go battery.Init(func() []BatteryArgument {
			gotNumCalls++

			return batteryArgumentsMock
		})

		wait := time.Duration(fixedInterval) * time.Duration(wantedNumCalls-1) * time.Second
		time.Sleep(wait)

		battery.Stop()

		if gotNumCalls != wantedNumCalls {
			t.Errorf("wanted ticks %d got %d", wantedNumCalls, gotNumCalls)
		}

		expectStorageToHasBatteryArgumentsMock(t, storage, batteryArgumentsMock)

	})

	t.Run("storage should have expected items", func(t *testing.T) {
		fixedInterval := 1
		storage := storages.NewGoCacheStorage()
		battery := NewBattery(storage, time.Duration(fixedInterval))

		go battery.Init(func() []BatteryArgument {

			return []BatteryArgument{}
		})

		battery.Dispatch <- batteryArgumentsMock

		expectStorageToHasBatteryArgumentsMock(t, storage, batteryArgumentsMock)

		battery.Stop()
	})
}
