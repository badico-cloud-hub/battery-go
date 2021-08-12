package batterygo

import "time"

type TimeoutIntervalConfig struct {
	event   string
	seconds time.Duration
}

func NewTimeoutIntervalConfig(event string, seconds time.Duration) TimeoutIntervalConfig {
	return TimeoutIntervalConfig{
		event,
		seconds,
	}
}
func TimeInterval(
	interval TimeoutIntervalConfig,
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
