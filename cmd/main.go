package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kirban/network-ninja/internal/config"
	"github.com/kirban/network-ninja/internal/ping"
	_ "github.com/kirban/network-ninja/internal/ping"
	"log"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	PingResults := make(chan ping.PingResult, 100)

	// 1) read data from config.yaml
	var c config.Config

	c.Load()

	// 2) we need to get ping_interval & max_concurrent_pings
	ping_interval := c.App.Ping_interval
	max_concurrent_pings := c.App.Max_concurrent_pings
	targets := len(c.Resources)

	// run send ping goroutine for each target
	for _, target := range c.Resources {
		go ping.SendPing(target.Address, target.Timeout)
	}

	for {
		select {
		case result := <-PingResults:
			// process result
			//ProcessPingResult(result, db, metricsExporter, config.Alerts)
			fmt.Printf("PingResult %+v", result)
		}
	}

	// 4) save to storage

	// 5) collect aggregated metrics (latency, packet loss, etc.) and save to storage

}
