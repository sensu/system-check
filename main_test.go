package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/prometheus/common/expfmt"
	"github.com/sensu/sensu-plugin-sdk/sensu"
	"github.com/stretchr/testify/assert"
)

func TestCheckArgs(t *testing.T) {
	assert := assert.New(t)
	state, err := checkArgs(nil)
	assert.NoError(err)
	assert.Equal(sensu.CheckStateOK, state)

}

func TestCollectMetrics(t *testing.T) {
	assert := assert.New(t)
	plugin.Interval = 3
	timestamp := int64(100)
	metrics, err := collectMetrics(timestamp)
	assert.NoError(err)
	output := ""
	for i := range metrics {
		output = output + "\n"
		output = output + fmt.Sprintf("%s\n", metrics[i].Output())
	}
	assert.Contains(output, "system_cpu_cores{}")
	assert.Contains(output, `system_cpu_idle{cpu="cpu-total"}`)
	assert.Contains(output, `system_cpu_idle{cpu="cpu0"}`)
	assert.Contains(output, `system_cpu_used{cpu="cpu-total"}`)
	assert.Contains(output, `system_cpu_used{cpu="cpu0"}`)
	assert.Contains(output, `system_cpu_user{cpu="cpu-total"}`)
	assert.Contains(output, `system_cpu_user{cpu="cpu0"}`)
	assert.Contains(output, `system_cpu_system{cpu="cpu-total"}`)
	assert.Contains(output, `system_cpu_system{cpu="cpu0"}`)
	assert.Contains(output, `system_cpu_nice{cpu="cpu-total"}`)
	assert.Contains(output, `system_cpu_nice{cpu="cpu0"}`)
	assert.Contains(output, `system_cpu_iowait{cpu="cpu-total"}`)
	assert.Contains(output, `system_cpu_iowait{cpu="cpu0"}`)
	assert.Contains(output, `system_cpu_irq{cpu="cpu-total"}`)
	assert.Contains(output, `system_cpu_irq{cpu="cpu0"}`)
	assert.Contains(output, `system_cpu_sortirq{cpu="cpu-total"}`)
	assert.Contains(output, `system_cpu_sortirq{cpu="cpu0"}`)
	assert.Contains(output, `system_cpu_stolen{cpu="cpu-total"}`)
	assert.Contains(output, `system_cpu_stolen{cpu="cpu0"}`)
	assert.Contains(output, `system_cpu_guest{cpu="cpu-total"}`)
	assert.Contains(output, `system_cpu_guest{cpu="cpu0"}`)
	assert.Contains(output, `system_cpu_guest_nice{cpu="cpu-total"}`)
	assert.Contains(output, `system_cpu_guest_nice{cpu="cpu0"}`)
	assert.Contains(output, "system_mem_used{}")
	assert.Contains(output, "system_mem_used_bytes{}")
	assert.Contains(output, "system_mem_total_bytes{}")
	assert.Contains(output, "system_swap_used{}")
	assert.Contains(output, "system_swap_used_bytes{}")
	assert.Contains(output, "system_swap_total_bytes{}")
	assert.Contains(output, "system_load_load1{}")
	assert.Contains(output, "system_load_load5{}")
	assert.Contains(output, "system_load_load15{}")
	assert.Contains(output, "system_load_load1_per_cpu{}")
	assert.Contains(output, "system_load_load5_per_cpu{}")
	assert.Contains(output, "system_load_load15_per_cpu{}")
	assert.Contains(output, "system_host_uptime{}")
	assert.Contains(output, "system_host_processes{}")

	var parser expfmt.TextParser
	_, err = parser.TextToMetricFamilies(strings.NewReader(output))
	assert.NoError(err)
	fmt.Println(output)

}
