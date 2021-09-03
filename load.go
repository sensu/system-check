package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
)

func getLoadMetrics(timestamp int64) ([]PromMetric, error) {

	info, err := cpu.Info()
	if err != nil {
		return nil, fmt.Errorf("Error: obtaining CPU info: %v", err)
	} else if len(info) < 1 {
		return nil, fmt.Errorf("Error: no CPU info found")
	}
	cores := float64(0)
	for _, i := range info {
		cores += float64(i.Cores)
	}
	var loadStats *load.AvgStat
	for i := 1; i < 10; i++ {
		loadStats, err = load.Avg()
	}
	if err != nil {
		return nil, fmt.Errorf("Error: obtaining load stas: %v", err)
	}

	metrics := []PromMetric{
		PromMetric{
			Label:           "system_load_load1",
			Value:           loadStats.Load1,
			Type:            "gauge",
			HelpComment:     "System load averaged over 1 minute, high load value dependant on number of cpus in system",
			IncludeComments: true,
		},
		PromMetric{
			Label:           "system_load_load5",
			Value:           loadStats.Load5,
			Type:            "gauge",
			HelpComment:     "System load averaged over 5 minute, high load value dependent on number of cpus in system",
			IncludeComments: true,
		},
		PromMetric{
			Label:           "system_load_load15",
			Value:           loadStats.Load15,
			Type:            "gauge",
			HelpComment:     "System load averaged over 15 minute, high load value dependent on number of cpus in system",
			IncludeComments: true,
		},
		PromMetric{
			Label:           "system_load_load1_per_cpu",
			Value:           loadStats.Load1 / cores,
			Type:            "gauge",
			HelpComment:     "System load averaged over 1 minute normalized by cpu count, values > 1 means system may be overloaded",
			IncludeComments: true,
		},
		PromMetric{
			Label:           "system_load_load5_per_cpu",
			Value:           loadStats.Load5 / cores,
			Type:            "gauge",
			HelpComment:     "System load averaged over 5 minute normalized by cpu count, values > 1 means system may be overloaded",
			IncludeComments: true,
		},
		PromMetric{
			Label:           "system_load_load15_per_cpu",
			Value:           loadStats.Load15 / cores,
			Type:            "gauge",
			HelpComment:     "System load averaged over 15 minute normalized by cpu count, values > 1 means system may be overloaded",
			IncludeComments: true,
		},
	}
	for i := range metrics {
		metrics[i].Timestamp = timestamp
	}
	return metrics, nil
}
