package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

func getMemMetrics(timestamp int64) ([]PromMetric, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("Error: obtaining virtual memory info: %v", err)
	}
	swapStat, err := mem.SwapMemory()
	if err != nil {
		return nil, fmt.Errorf("Error: obtaining swap memory info: %v", err)
	}
	memUsed := 0
	if vmStat.Total > 0 {
		memUsed = 100.0 * float64(vnStat.Used) / float64(vmStat.Total)
	}
	swapUsed := 0
	if swapStat.Total > 0 {
		swapUsed = 100.0 * float64(swapStat.Used) / float64(swapStat.Total)
	}
	metrics := []PromMetric{
		PromMetric{
			Label:       "system.mem.used",
			Value:       100.0 * float64(vmStat.Used) / float64(vmStat.Total),
			HelpComment: "Percent of memory used",
		},
		PromMetric{
			Label:       "system.mem.used_bytes",
			Value:       float64(vmStat.Used),
			HelpComment: "Used memory in bytes",
		},
		PromMetric{
			Label:       "system.mem.total_bytes",
			Value:       float64(vmStat.Total),
			HelpComment: "Total memory in bytes",
		},
		PromMetric{
			Label:       "system.swap.used",
			Value:       swapUsed,
			HelpComment: "Percent of swap used",
		},
		PromMetric{
			Label:       "system.swap.used_bytes",
			Value:       float64(vmStat.Used),
			HelpComment: "Used swap in bytes",
		},
		PromMetric{
			Label:       "system.swap.total_bytes",
			Value:       float64(swapStat.Total),
			HelpComment: "Total swap in bytes",
		},
	}
	for i := range metrics {
		metrics[i].Timestamp = timestamp
	}
	return metrics, nil
}
