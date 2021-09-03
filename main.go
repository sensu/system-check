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
	Label           string
	Tags            []string
	Timestamp       int64
	Value           float64
	Type            string
	IncludeComments bool
	HelpComment     string
}

func (m PromMetric) Output() string {
	comments := ""
	if m.IncludeComments {
		comments = fmt.Sprintf("# HELP %s [%s] %s\n# TYPE %s %s\n", m.Label, strings.ToUpper(m.Type), m.HelpComment, m.Label, strings.ToUpper(m.Type))
	}
	return fmt.Sprintf("%s%s{%s} %v %v", comments, strings.ReplaceAll(m.Label, ".", "_"), strings.Join(m.Tags, ","), m.Value, m.Timestamp)
}

func collectMetrics(timestamp int64) ([]PromMetric, error) {
	metrics := []PromMetric{}

	if m, err := getCPUMetrics(timestamp, plugin.Interval); err != nil {
		return nil, err
	} else {
		metrics = append(metrics, m...)
	}
	if m, err := getMemMetrics(timestamp); err != nil {
		return nil, err
	} else {
		metrics = append(metrics, m...)
	}
	if m, err := getLoadMetrics(timestamp); err != nil {
		return nil, err
	} else {
		metrics = append(metrics, m...)
	}
	if m, err := getHostMetrics(timestamp); err != nil {
		return nil, err
	} else {
		metrics = append(metrics, m...)
	}

	return metrics, nil
}

func executeCheck(event *v2.Event) (int, error) {
	milliEpoch := time.Now().UnixNano() / 1000000
	metrics, err := collectMetrics(milliEpoch)
	if err != nil {
		return sensu.CheckStateCritical, err
	}
	for i := range metrics {
		fmt.Println("")
		fmt.Println(metrics[i].Output())
	}
	return sensu.CheckStateOK, nil
}
