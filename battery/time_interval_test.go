package batterygo_test

import (
	"testing"
	"time"

	. "github.com/badico-cloud-hub/battery-go/battery"
)

func TestTimeInterval(t *testing.T) {
	setupTest := func(t *testing.T, wantedEvent string, wantedInterval time.Duration) {
		quit := make(chan struct{})
		timeout := make(chan string)

		go TimeInterval(
			NewTimeoutIntervalConfig(wantedEvent, wantedInterval),
			timeout,
			quit,
		)

		start := time.Now()

		gotEvent := <-timeout

		gotInterval := time.Now().Sub(start)

		quit <- struct{}{}

		if gotEvent != wantedEvent {
			t.Errorf("Wanted event (%s) got event (%s)", wantedEvent, gotEvent)
		}

		if gotInterval < wantedInterval {
			t.Errorf("wanted Interval was %s Event dispatch after  %s", wantedInterval, gotInterval)
		}
	}

	t.Run("should dispatch event after 2 seconds", func(t *testing.T) {
		wantedEvent := "some_event"
		wantedInterval := time.Duration(2)

		setupTest(t, wantedEvent, wantedInterval)
	})

	t.Run("should dispatch event after 4 seconds", func(t *testing.T) {
		wantedEvent := "another_event"
		wantedInterval := time.Duration(4)

		setupTest(t, wantedEvent, wantedInterval)
	})
}

func TestTimeIntervalStop(t *testing.T) {
	runTimeIntervalAndStop := func(t *testing.T, wantedTicks int) int {
		quit := make(chan struct{})
		timeout := make(chan string)
		endRoutine := make(chan struct{})
		tickerStopped := make(chan struct{})

		fixedInterval := time.Duration(1)

		go TimeInterval(
			NewTimeoutIntervalConfig("some_event", fixedInterval),
			timeout,
			quit,
		)

		var gotTicks int
		go func() {
			for {
				select {
				case <-timeout:
					gotTicks++
				case <-endRoutine:
					tickerStopped <- struct{}{}
					return
				}
			}
		}()

		second := fixedInterval * time.Duration(wantedTicks) * time.Second

		msAdjust := (fixedInterval * 100) * time.Millisecond

		wait := second + msAdjust
		time.Sleep(wait)

		quit <- struct{}{}

		endRoutine <- struct{}{}

		<-tickerStopped

		return gotTicks
	}

	t.Run("should stop timeInterval after 4 ticks", func(t *testing.T) {
		wantedTicks := 4

		gotTicks := runTimeIntervalAndStop(t, wantedTicks)

		if gotTicks != wantedTicks {
			t.Errorf("wanted ticks %d got %d", wantedTicks, gotTicks)
		}
	})
}
