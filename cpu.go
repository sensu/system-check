package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func getCPUMetrics(timestamp int64, interval int) ([]PromMetric, error) {

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

	// take the starting time stat
	cpuStartTimes, err := cpu.Times(true)
	if err != nil {
		return nil, fmt.Errorf("Error: obtaining initial CPU timings: %v", err)
	} else if len(cpuStartTimes) < 1 {
		return nil, fmt.Errorf("Error: no CPU timings found")
	}
	totalStartTimes, err := cpu.Times(false)
	if err != nil {
		return nil, fmt.Errorf("Error: obtaining initial CPU timings: %v", err)
	} else if len(totalStartTimes) < 1 {
		return nil, fmt.Errorf("Error: no CPU timings found")
	}
	cpuStartTimes = append(cpuStartTimes, totalStartTimes...)

	//Now lets wait an interval
	duration, err := time.ParseDuration(fmt.Sprintf("%ds", interval))
	if err != nil {
		return nil, fmt.Errorf("Error: parsing duration: %v", err)
	}
	time.Sleep(duration)

	// take the ending time stats
	cpuFinalTimes, err := cpu.Times(true)
	if err != nil {
		return nil, fmt.Errorf("Error: obtaining final CPU timings: %v", err)
	} else if len(cpuFinalTimes) < 1 {
		return nil, fmt.Errorf("Error: no CPU timings found")
	}
	totalFinalTimes, err := cpu.Times(false)
	if err != nil {
		return nil, fmt.Errorf("Error: obtaining final CPU timings: %v", err)
	} else if len(totalFinalTimes) < 1 {
		return nil, fmt.Errorf("Error: no CPU timings found")
	}
	cpuFinalTimes = append(cpuFinalTimes, totalFinalTimes...)

	metrics := []PromMetric{
		PromMetric{
			Label:       "system_cpu_cores",
			Value:       cores,
			Type:        "gauge",
			HelpComment: "Number of cpu cores on the system",
		},
	}
	for i := range cpuStartTimes {
		includeComments := false
		if cpuStartTimes[i].CPU != cpuFinalTimes[i].CPU {
			return nil, fmt.Errorf("Error: CPU name mismatch between start and final timing collection")
		}
		cpuName := cpuStartTimes[i].CPU
		initialTimings := cpuStartTimes[i]
		finalTimings := cpuFinalTimes[i]
		diffTotal := finalTimings.Total() - initialTimings.Total()
		idlePct := ((finalTimings.Idle - initialTimings.Idle) / diffTotal) * 100
		if i == 0 {
			includeComments = true
		}
		metrics = append(metrics, PromMetric{
			Label:           "system_cpu_idle",
			Value:           idlePct,
			Type:            "gauge",
			HelpComment:     "Percent of time all cpus were idle",
			IncludeComments: includeComments,
			Tags: []string{
				fmt.Sprintf(`cpu="%s"`, cpuName),
			},
		})
	}
	for i := range cpuStartTimes {
		includeComments := false
		if cpuStartTimes[i].CPU != cpuFinalTimes[i].CPU {
			return nil, fmt.Errorf("Error: CPU name mismatch between start and final timing collection")
		}
		cpuName := cpuStartTimes[i].CPU
		initialTimings := cpuStartTimes[i]
		finalTimings := cpuFinalTimes[i]
		diffTotal := finalTimings.Total() - initialTimings.Total()
		idlePct := ((finalTimings.Idle - initialTimings.Idle) / diffTotal) * 100
		if i == 0 {
			includeComments = true
		}
		metrics = append(metrics, PromMetric{
			Label:           "system_cpu_used",
			Value:           100 - idlePct,
			Type:            "gauge",
			HelpComment:     "Percent of time all cpus were used",
			IncludeComments: includeComments,
			Tags: []string{
				fmt.Sprintf(`cpu="%s"`, cpuName),
			},
		})
	}
	for i := range cpuStartTimes {
		includeComments := false
		if cpuStartTimes[i].CPU != cpuFinalTimes[i].CPU {
			return nil, fmt.Errorf("Error: CPU name mismatch between start and final timing collection")
		}
		cpuName := cpuStartTimes[i].CPU
		initialTimings := cpuStartTimes[i]
		finalTimings := cpuFinalTimes[i]
		diffTotal := finalTimings.Total() - initialTimings.Total()
		if i == 0 {
			includeComments = true
		}
		metrics = append(metrics, PromMetric{
			Label:           "system_cpu_user",
			Value:           ((finalTimings.User - initialTimings.User) / diffTotal) * 100,
			Type:            "gauge",
			HelpComment:     "Percent of time total cpu was used by normal processes in user mode",
			IncludeComments: includeComments,
			Tags: []string{
				fmt.Sprintf(`cpu="%s"`, cpuName),
			},
		})
	}
	for i := range cpuStartTimes {
		includeComments := false
		if cpuStartTimes[i].CPU != cpuFinalTimes[i].CPU {
			return nil, fmt.Errorf("Error: CPU name mismatch between start and final timing collection")
		}
		cpuName := cpuStartTimes[i].CPU
		initialTimings := cpuStartTimes[i]
		finalTimings := cpuFinalTimes[i]
		diffTotal := finalTimings.Total() - initialTimings.Total()
		if i == 0 {
			includeComments = true
		}
		metrics = append(metrics, PromMetric{
			Label:           "system_cpu_system",
			Value:           ((finalTimings.System - initialTimings.System) / diffTotal) * 100,
			Type:            "gauge",
			HelpComment:     "Percent of time all cpus used by processes executed in kernel mode",
			IncludeComments: includeComments,
			Tags: []string{
				fmt.Sprintf(`cpu="%s"`, cpuName),
			},
		})
	}
	for i := range cpuStartTimes {
		includeComments := false
		if cpuStartTimes[i].CPU != cpuFinalTimes[i].CPU {
			return nil, fmt.Errorf("Error: CPU name mismatch between start and final timing collection")
		}
		cpuName := cpuStartTimes[i].CPU
		initialTimings := cpuStartTimes[i]
		finalTimings := cpuFinalTimes[i]
		diffTotal := finalTimings.Total() - initialTimings.Total()
		if i == 0 {
			includeComments = true
		}
		metrics = append(metrics, PromMetric{
			Label:           "system_cpu_nice",
			Value:           ((finalTimings.Nice - initialTimings.Nice) / diffTotal) * 100,
			Type:            "gauge",
			HelpComment:     "Percent of time all cpus used by niced processes in user mode",
			IncludeComments: includeComments,
			Tags: []string{
				fmt.Sprintf(`cpu="%s"`, cpuName),
			},
		})
	}
	for i := range cpuStartTimes {
		includeComments := false
		if cpuStartTimes[i].CPU != cpuFinalTimes[i].CPU {
			return nil, fmt.Errorf("Error: CPU name mismatch between start and final timing collection")
		}
		cpuName := cpuStartTimes[i].CPU
		initialTimings := cpuStartTimes[i]
		finalTimings := cpuFinalTimes[i]
		diffTotal := finalTimings.Total() - initialTimings.Total()
		if i == 0 {
			includeComments = true
		}
		metrics = append(metrics, PromMetric{
			Label:           "system_cpu_iowait",
			Value:           ((finalTimings.Iowait - initialTimings.Iowait) / diffTotal) * 100,
			Type:            "gauge",
			HelpComment:     "Percent of time all cpus waiting for I/O to complete",
			IncludeComments: includeComments,
			Tags: []string{
				fmt.Sprintf(`cpu="%s"`, cpuName),
			},
		})
	}
	for i := range cpuStartTimes {
		includeComments := false
		if cpuStartTimes[i].CPU != cpuFinalTimes[i].CPU {
			return nil, fmt.Errorf("Error: CPU name mismatch between start and final timing collection")
		}
		cpuName := cpuStartTimes[i].CPU
		initialTimings := cpuStartTimes[i]
		finalTimings := cpuFinalTimes[i]
		diffTotal := finalTimings.Total() - initialTimings.Total()
		if i == 0 {
			includeComments = true
		}
		metrics = append(metrics, PromMetric{
			Label:           "system_cpu_irq",
			Value:           ((finalTimings.Irq - initialTimings.Irq) / diffTotal) * 100,
			Type:            "gauge",
			HelpComment:     "Percent of time all cpus servicing interrupts",
			IncludeComments: includeComments,
			Tags: []string{
				fmt.Sprintf(`cpu="%s"`, cpuName),
			},
		})
	}
	for i := range cpuStartTimes {
		includeComments := false
		if cpuStartTimes[i].CPU != cpuFinalTimes[i].CPU {
			return nil, fmt.Errorf("Error: CPU name mismatch between start and final timing collection")
		}
		cpuName := cpuStartTimes[i].CPU
		initialTimings := cpuStartTimes[i]
		finalTimings := cpuFinalTimes[i]
		diffTotal := finalTimings.Total() - initialTimings.Total()
		if i == 0 {
			includeComments = true
		}
		metrics = append(metrics, PromMetric{
			Label:           "system_cpu_sortirq",
			Value:           ((finalTimings.Softirq - initialTimings.Softirq) / diffTotal) * 100,
			Type:            "gauge",
			HelpComment:     "Percent of time all cpus servicing software interrupts",
			IncludeComments: includeComments,
			Tags: []string{
				fmt.Sprintf(`cpu="%s"`, cpuName),
			},
		})
	}
	for i := range cpuStartTimes {
		includeComments := false
		if cpuStartTimes[i].CPU != cpuFinalTimes[i].CPU {
			return nil, fmt.Errorf("Error: CPU name mismatch between start and final timing collection")
		}
		cpuName := cpuStartTimes[i].CPU
		initialTimings := cpuStartTimes[i]
		finalTimings := cpuFinalTimes[i]
		diffTotal := finalTimings.Total() - initialTimings.Total()
		if i == 0 {
			includeComments = true
		}
		metrics = append(metrics, PromMetric{
			Label:           "system_cpu_stolen",
			Value:           ((finalTimings.Steal - initialTimings.Steal) / diffTotal) * 100,
			Type:            "gauge",
			HelpComment:     "Percent of time all cpus serviced virtual hosts operating systems",
			IncludeComments: includeComments,
			Tags: []string{
				fmt.Sprintf(`cpu="%s"`, cpuName),
			},
		})
	}
	for i := range cpuStartTimes {
		includeComments := false
		if cpuStartTimes[i].CPU != cpuFinalTimes[i].CPU {
			return nil, fmt.Errorf("Error: CPU name mismatch between start and final timing collection")
		}
		cpuName := cpuStartTimes[i].CPU
		initialTimings := cpuStartTimes[i]
		finalTimings := cpuFinalTimes[i]
		diffTotal := finalTimings.Total() - initialTimings.Total()
		if i == 0 {
			includeComments = true
		}
		metrics = append(metrics, PromMetric{
			Label:           "system_cpu_guest",
			Value:           ((finalTimings.Guest - initialTimings.Guest) / diffTotal) * 100,
			Type:            "gauge",
			HelpComment:     "Percent of time all cpus serviced guest operating system",
			IncludeComments: includeComments,
			Tags: []string{
				fmt.Sprintf(`cpu="%s"`, cpuName),
			},
		})
	}
	for i := range cpuStartTimes {
		includeComments := false
		if cpuStartTimes[i].CPU != cpuFinalTimes[i].CPU {
			return nil, fmt.Errorf("Error: CPU name mismatch between start and final timing collection")
		}
		cpuName := cpuStartTimes[i].CPU
		initialTimings := cpuStartTimes[i]
		finalTimings := cpuFinalTimes[i]
		diffTotal := finalTimings.Total() - initialTimings.Total()
		if i == 0 {
			includeComments = true
		}
		metrics = append(metrics, PromMetric{
			Label:           "system_cpu_guest_nice",
			Value:           ((finalTimings.GuestNice - initialTimings.GuestNice) / diffTotal) * 100,
			Type:            "gauge",
			HelpComment:     "Percent of time all cpus serviced niced guest operating system",
			IncludeComments: includeComments,
			Tags: []string{
				fmt.Sprintf(`cpu="%s"`, cpuName),
			},
		})
	}
	for i := range metrics {
		metrics[i].Timestamp = timestamp
	}
	return metrics, nil
}
