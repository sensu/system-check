package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	v2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-plugin-sdk/sensu"
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
			Path:      "interval",
			Env:       "SYSTEM_CHECK_INTERVAL",
			Argument:  "interval",
			Shorthand: "i",
			Default:   3,
			Usage:     "Interval in seconds over which to calculate cpu usage metrics",
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
	timestamp := time.Now().Unix()

	metrics, err := getCPUMetrics(timestamp, plugin.Interval)
	if err != nil {
		return sensu.CheckStateCritical, err
	}
	for i := range metrics {
		fmt.Println("")
		fmt.Println(metrics[i].Output())
	}
	metrics, err = getMemMetrics(timestamp)
	if err != nil {
		return sensu.CheckStateCritical, err
	}
	for i := range metrics {
		fmt.Println("")
		fmt.Println(metrics[i].Output())
	}
	metrics, err = getLoadMetrics(timestamp)
	if err != nil {
		return sensu.CheckStateCritical, err
	}
	for i := range metrics {
		fmt.Println("")
		fmt.Println(metrics[i].Output())
	}
	metrics, err = getHostMetrics(timestamp)
	if err != nil {
		return sensu.CheckStateCritical, err
	}
	for i := range metrics {
		fmt.Println("")
		fmt.Println(metrics[i].Output())
	}
	return sensu.CheckStateOK, nil
}
