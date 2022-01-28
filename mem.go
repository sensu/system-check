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
	memUsed := float64(0)
	if vmStat.Total > 0 {
		memUsed = 100.0 * float64(vmStat.Used) / float64(vmStat.Total)
	}
	swapUsed := float64(0)
	if swapStat.Total > 0 {
		swapUsed = 100.0 * float64(swapStat.Used) / float64(swapStat.Total)
	}
	metrics := []PromMetric{
		PromMetric{
			Label:           "system_mem_used",
			Value:           memUsed,
			Type:            "gauge",
			HelpComment:     "Percent of memory used",
			IncludeComments: true,
		},
		PromMetric{
			Label:           "system_mem_used_bytes",
			Value:           float64(vmStat.Used),
			Type:            "gauge",
			HelpComment:     "Used memory in bytes",
			IncludeComments: true,
		},
		PromMetric{
			Label:           "system_mem_total_bytes",
			Value:           float64(vmStat.Total),
			Type:            "gauge",
			HelpComment:     "Total memory in bytes",
			IncludeComments: true,
		},
		PromMetric{
			Label:           "system_swap_used",
			Value:           swapUsed,
			Type:            "gauge",
			HelpComment:     "Percent of swap used",
			IncludeComments: true,
		},
		PromMetric{
			Label:           "system_swap_used_bytes",
			Value:           float64(swapStat.Used),
			Type:            "gauge",
			HelpComment:     "Used swap in bytes",
			IncludeComments: true,
		},
		PromMetric{
			Label:           "system_swap_total_bytes",
			Value:           float64(swapStat.Total),
			Type:            "gauge",
			HelpComment:     "Total swap in bytes",
			IncludeComments: true,
		},
	}
	for i := range metrics {
		metrics[i].Timestamp = timestamp
	}
	return metrics, nil
}
