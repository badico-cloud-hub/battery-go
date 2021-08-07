package batterygo

import "time"

type timeoutIntervalConfig struct {
	event   string
	seconds time.Duration
}

func timeInterval(
	interval timeoutIntervalConfig,
	dispatch chan string,
	quit chan struct{},
) {
	ticker := time.NewTicker(interval.seconds * time.Second)
	for {
		select {
		case <-ticker.C:
			dispatch <- interval.event
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
