package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	v2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-plugin-sdk/sensu"
	"github.com/shirou/gopsutil/v3/cpu"
)

// Config represents the check plugin config.
type Config struct {
	sensu.PluginConfig
	Interval int
	Example  string
}

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "system-check",
			Short:    "Cross-platform System Check",
			Keyspace: "sensu.io/plugins/system-check/config",
		},
	}

	options = []*sensu.PluginConfigOption{
		&sensu.PluginConfigOption{
			Path:      "example",
			Env:       "CHECK_EXAMPLE",
			Argument:  "example",
			Shorthand: "e",
			Default:   "",
			Usage:     "An example string configuration option",
			Value:     &plugin.Example,
		},
		&sensu.PluginConfigOption{
			Path:      "interval",
			Env:       "SYSTEM_CHECK_INTERVAL",
			Argument:  "interval",
			Shorthand: "i",
			Default:   3,
			Usage:     "An example string configuration option",
			Value:     &plugin.Interval,
		},
	}
)

func main() {
	useStdin := false
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Printf("Error check stdin: %v\n", err)
		panic(err)
	}
	//Check the Mode bitmask for Named Pipe to indicate stdin is connected
	if fi.Mode()&os.ModeNamedPipe != 0 {
		log.Println("using stdin")
		useStdin = true
	}

	check := sensu.NewGoCheck(&plugin.PluginConfig, options, checkArgs, executeCheck, useStdin)
	check.Execute()
}

func checkArgs(event *v2.Event) (int, error) {
	/*
		if len(plugin.Example) == 0 {
			return sensu.CheckStateWarning, fmt.Errorf("--example or CHECK_EXAMPLE environment variable is required")
		}
	*/
	return sensu.CheckStateOK, nil
}

type PromMetric struct {
	Label       string
	Tags        []string
	Timestamp   int64
	Value       float64
	HelpComment string
}

func (m PromMetric) Output() string {
	return fmt.Sprintf("# HELP %s\n%s{%s} %v %v", m.HelpComment, m.Label, strings.Join(m.Tags, ","), m.Timestamp, m.Value)
}

func executeCheck(event *v2.Event) (int, error) {

	info, err := cpu.Info()
	if err != nil {
		return sensu.CheckStateCritical, fmt.Errorf("Error: obtaining CPU info: %v", err)
	} else if len(info) < 1 {
		return sensu.CheckStateCritical, fmt.Errorf("Error: no CPU info found")
	}
	cores := float64(0)
	for _, i := range info {
		cores += float64(i.Cores)
	}

	// take the starting time stat
	start, err := cpu.Times(false)
	if err != nil {
		return sensu.CheckStateCritical, fmt.Errorf("Error: obtaining initial CPU timings: %v", err)
	} else if len(start) < 1 {
		return sensu.CheckStateCritical, fmt.Errorf("Error: no CPU timings found")
	}
	initialTimings := start[0]

	//Now lets wait an interval
	duration, err := time.ParseDuration(fmt.Sprintf("%ds", plugin.Interval))
	if err != nil {
		return sensu.CheckStateCritical, fmt.Errorf("Error: parsing duration: %v", err)
	}
	time.Sleep(duration)

	// take the ending time stats
	final, err := cpu.Times(false)
	if err != nil {
		return sensu.CheckStateCritical, fmt.Errorf("Error: obtaining final CPU timings: %v", err)
	}
	// Final timing
	finalTimings := final[0]
	timestamp := time.Now().Unix()
	// Difference in total time across interval
	// Note: total time for N cores is N x total time for single core
	diffTotal := finalTimings.Total() - initialTimings.Total()
	idlePct := ((finalTimings.Idle - initialTimings.Idle) / diffTotal) * 100

	metrics := []PromMetric{
		PromMetric{
			Label:       "system.cpu.idle",
			Value:       idlePct,
			HelpComment: "Percent of time cpu was idle",
		},
		PromMetric{
			Label:       "system.cpu.used",
			Value:       100 - idlePct,
			HelpComment: "Percent of time cpu was used",
		},
		PromMetric{
			Label:       "system.cpu.user",
			Value:       ((finalTimings.User - initialTimings.User) / diffTotal) * 100,
			HelpComment: "Percent of time cpu was used by normal processes in user mode",
		},
		PromMetric{
			Label:       "system.cpu.system",
			Value:       ((finalTimings.System - initialTimings.System) / diffTotal) * 100,
			HelpComment: "Percent of time cpu used by processes executed in kernel mode",
		},
		PromMetric{
			Label:       "system.cpu.nice",
			Value:       ((finalTimings.Nice - initialTimings.Nice) / diffTotal) * 100,
			HelpComment: "Percent of time cpu used by niced processes in user mode",
		},
		PromMetric{
			Label:       "system.cpu.iowait",
			Value:       ((finalTimings.Iowait - initialTimings.Iowait) / diffTotal) * 100,
			HelpComment: "Percent of time cpu waiting for I/O to complete",
		},
		PromMetric{
			Label:       "system.cpu.irq",
			Value:       ((finalTimings.Irq - initialTimings.Irq) / diffTotal) * 100,
			HelpComment: "Percent of time cpu servicing interrupts",
		},
		PromMetric{
			Label:       "system.cpu.sortirq",
			Value:       ((finalTimings.Softirq - initialTimings.Softirq) / diffTotal) * 100,
			HelpComment: "Percent of time cpu servicing software interrupts",
		},
		PromMetric{
			Label:       "system.cpu.stolen",
			Value:       ((finalTimings.Steal - initialTimings.Steal) / diffTotal) * 100,
			HelpComment: "Percent of time cpu serviced virtual hosts operating systems",
		},
		PromMetric{
			Label:       "system.cpu.guest",
			Value:       ((finalTimings.Guest - initialTimings.Guest) / diffTotal) * 100,
			HelpComment: "Percent of time cpu serviced guest operating system",
		},
		PromMetric{
			Label:       "system.cpu.guest_nice",
			Value:       ((finalTimings.GuestNice - initialTimings.GuestNice) / diffTotal) * 100,
			HelpComment: "Percent of time cpu serviced niced guest operating system",
		},
	}
	for i := range metrics {
		metrics[i].Timestamp = timestamp
		fmt.Println("")
		fmt.Println(metrics[i].Output())
	}
	return sensu.CheckStateOK, nil
}
