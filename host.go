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
			Label:           "system_host_uptime",
			Value:           float64(uptime),
			Type:            "counter",
			HelpComment:     "Host uptime in seconds",
			IncludeComments: true,
		},
		PromMetric{
			Label:           "system_host_processes",
			Value:           float64(info.Procs),
			Type:            "gauge",
			HelpComment:     "Number of host processes",
			IncludeComments: true,
		},
	}
	for i := range metrics {
		metrics[i].Timestamp = timestamp
	}
	return metrics, nil
}
