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
	loadStats, err := load.Avg()
	if err != nil {
		return nil, fmt.Errorf("Error: obtaining load stas: %v", err)
	}

	metrics := []PromMetric{
		PromMetric{
			Label:       "system.load.load1",
			Value:       loadStats.Load1,
			HelpComment: "System load averaged over 1 minute, high load value dependant on number of cpus in system",
		},
		PromMetric{
			Label:       "system.load.load5",
			Value:       loadStats.Load5,
			HelpComment: "System load averaged over 5 minute, high load value dependent on number of cpus in system",
		},
		PromMetric{
			Label:       "system.load.load15",
			Value:       loadStats.Load15,
			HelpComment: "System load averaged over 15 minute, high load value dependent on number of cpus in system",
		},
		PromMetric{
			Label:       "system.load.load1_per_cpu",
			Value:       loadStats.Load1 / cores,
			HelpComment: "System load averaged over 1 minute normalized by cpu count, values > 1 means system may be overloaded",
		},
		PromMetric{
			Label:       "system.load.load5_per_cpu",
			Value:       loadStats.Load5 / cores,
			HelpComment: "System load averaged over 5 minute normalized by cpu count, values > 1 means system may be overloaded",
		},
		PromMetric{
			Label:       "system.load.load15_per_cpu",
			Value:       loadStats.Load15 / cores,
			HelpComment: "System load averaged over 15 minute normalized by cpu count, values > 1 means system may be overloaded",
		},
	}
	for i := range metrics {
		metrics[i].Timestamp = timestamp
	}
	return metrics, nil
}
