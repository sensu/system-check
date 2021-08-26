package main

import (
	"fmt"

	"github.com/shirou/gopsutil/host"
)

func getHostMetrics(timestamp int64) ([]PromMetric, error) {

	uptime, err := host.Uptime()
	if err != nil {
		return nil, fmt.Errorf("Error: obtaining host uptime: %v", err)
	}
	info, err := host.Info()
	if err != nil {
		return nil, fmt.Errorf("Error: obtaining host info: %v", err)
	}
	metrics := []PromMetric{
		PromMetric{
			Label:       "system.host.uptime",
			Value:       float64(uptime),
			HelpComment: "Host uptime in seconds",
		},
		PromMetric{
			Label:       "system.host.processes",
			Value:       float64(info.Procs),
			HelpComment: "Number of host processes",
		},
	}
	for i := range metrics {
		metrics[i].Timestamp = timestamp
	}
	return metrics, nil
}
