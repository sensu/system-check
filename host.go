package main

import (
	"fmt"

	"github.com/shirou/gopsutil/host"
)

func getHostMetrics(timestamp int64) ([]PromMetric, error) {

	uptime, err := host.Uptime()
	if err != nil {
		return nil, fmt.Errorf("Error: obtaining uptime: %v", err)
	}
	metrics := []PromMetric{
		PromMetric{
			Label:       "system.host.uptime",
			Value:       float64(uptime),
			HelpComment: "Host uptime in seconds",
		},
	}
	for i := range metrics {
		metrics[i].Timestamp = timestamp
	}
	return metrics, nil
}
