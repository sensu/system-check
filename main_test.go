package main

import (
	"fmt"
	"testing"

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
	assert.Contains(output, "system.cpu.cores{} 100")
	assert.Contains(output, "system.cpu.idle{} 100")
	assert.Contains(output, "system.cpu.used{} 100")
	assert.Contains(output, "system.cpu.user{} 100")
	assert.Contains(output, "system.cpu.system{} 100")
	assert.Contains(output, "system.cpu.nice{} 100")
	assert.Contains(output, "system.cpu.iowait{} 100")
	assert.Contains(output, "system.cpu.irq{} 100")
	assert.Contains(output, "system.cpu.sortirq{} 100")
	assert.Contains(output, "system.cpu.stolen{} 100")
	assert.Contains(output, "system.cpu.guest{} 100")
	assert.Contains(output, "system.cpu.guest_nice{} 100")
	assert.Contains(output, "system.mem.used{} 100")
	assert.Contains(output, "system.mem.used_bytes{} 100")
	assert.Contains(output, "system.mem.total_bytes{} 100")
	assert.Contains(output, "system.swap.used{} 100")
	assert.Contains(output, "system.swap.used_bytes{} 100")
	assert.Contains(output, "system.swap.total_bytes{} 100")
	assert.Contains(output, "system.load.load1{} 100")
	assert.Contains(output, "system.load.load5{} 100")
	assert.Contains(output, "system.load.load15{} 100")
	assert.Contains(output, "system.load.load1_per_cpu{} 100")
	assert.Contains(output, "system.load.load5_per_cpu{} 100")
	assert.Contains(output, "system.load.load15_per_cpu{} 100")
	assert.Contains(output, "system.host.uptime{} 100")
	assert.Contains(output, "system.host.processes{} 100")
	fmt.Println(output)

}
