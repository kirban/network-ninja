package ping

import (
	probing "github.com/prometheus-community/pro-bing"
	"time"
)

type PingResult struct {
	Resource  string        `json:"resource"`
	Success   bool          `json:"success"`
	Timestamp time.Time     `json:"timestamp"`
	Latency   time.Duration `json:"latency"`
}

func SendPing(target string) *PingResult {
	var pinger, err = probing.NewPinger("www.google.com")

	if err != nil {
		panic(err)
	}

	pinger.Count = 3
	pinger.Interval = 1 * time.Second
	pinger.Timeout = 10 * time.Second
	//pinger.TTL =

	err = pinger.Run()

	if err != nil {
		panic(err)
	}

	stats := pinger.Statistics()

	return &PingResult{
		Resource:  target,
		Success:   (stats.PacketsSent - stats.PacketsRecv) == 0,
		Timestamp: time.Now(),
		Latency:   stats.AvgRtt,
	}
}
