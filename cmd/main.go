package main

import (
	"fmt"
	"github.com/kirban/network-ninja/internal/config"
	_ "github.com/kirban/network-ninja/internal/ping"
)

func main() {
	// 1) read data from config.yaml

	var c config.Config

	c.Load()

	// 2) we need to get ping_interval & max_concurrent_pings

	// 3) in loop:
	// every N (ping_interval) seconds run K (targets) <= M(max_concurrent_pings) goroutines
	// if K <= M - run K goroutines
	// if K > M
	// например N = 1s, M = 5, K = 10
	//		numChunks = K / M
	//		timeBetweenChunks = N / numChunks
	// N = 1s, M = 5, K = 23
	//		numChunks = 23 / 5 = 4.6 = 5
	//		timeBetweenChunks = 1 / 5 = 0.2s

	// 4) save to storage

	// 5) collect aggregated metrics (latency, packet loss, etc.) and save to storage

}
