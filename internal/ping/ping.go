package ping

import (
	"fmt"
	probing "github.com/prometheus-community/pro-bing"
	"time"
)

type PingResult struct {
	Resource  string        `json:"resource"`
	Success   bool          `json:"success"`
	Timestamp time.Time     `json:"timestamp"`
	Latency   time.Duration `json:"latency"`
}

func SendPing(target string, ping_interval int) *PingResult {
	var pinger, err = probing.NewPinger(target)

	if err != nil {
		panic(err)
	}

	pinger.Count = 3
	pinger.Interval = time.Duration(ping_interval) * time.Second

	err = pinger.Run()

	if err != nil {
		panic(err)
	}

	stats := pinger.Statistics()

	fmt.Println("sent ", stats.PacketsSent)
	fmt.Println("recieved ", stats.PacketsRecv)

	result := &PingResult{
		Resource:  target,
		Success:   (stats.PacketsSent - stats.PacketsRecv) == 0,
		Timestamp: time.Now(),
		Latency:   stats.AvgRtt,
	}

	fmt.Printf("stats %+v\n", stats)

	return result
}
