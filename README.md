[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/sensu/system-check)
![Go Test](https://github.com/sensu/system-check/workflows/Go%20Test/badge.svg)
![goreleaser](https://github.com/sensu/system-check/workflows/goreleaser/badge.svg)

# Sensu System Check

## Table of Contents
- [Overview](#overview)
  - [Output Metrics](#output-metrics)
- [Usage examples](#usage-examples)
  - [Help output](#help-output)
  - [Environment variables](#environment-variables)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Check definition](#check-definition)
- [Installation from source](#installation-from-source)
- [Additional notes](#additional-notes)
- [Contributing](#contributing)

## Overview

The [Sensu System Check][1] is a cross-platform [Sensu Metrics Check][7] that provides baseline system metrics in prometheus format. 

### Output Metrics
#### cpu
| Name              | Description   |
|-------------------|---------------|
| system.cpu.cores  | Number of cpu cores on the system |
| system.cpu.idle   | Percent of time cpu was idle |
| system.cpu.used   | Percent of time cpu was used
| system.cpu.user   | Percent of time cpu was used by normal processes in user mode
| system.cpu.system | Percent of time cpu used by processes executed in kernel mode
| system.cpu.nice   | Percent of time cpu used by niced processes in user mode
| system.cpu.iowait | Percent of time cpu waiting for I/O to complete
| system.cpu.irq    | Percent of time cpu servicing interrupts
| system.cpu.sortirq | Percent of time cpu servicing software interrupts
| system.cpu.stolen | Percent of time cpu serviced virtual hosts operating systems
| system.cpu.guest | Percent of time cpu serviced guest operating system
| system.cpu.guest_nice | Percent of time cpu serviced niced guest operating system



#### mem
| Name              | Description   |
|-------------------|---------------|
| system.mem.used | Percent of memory used
| system.mem.used_bytes | Used memory in bytes
| system.mem.total_bytes | Total memory in bytes

#### swap
| Name              | Description   |
|-------------------|---------------|
| system.swap.used | Percent of swap used
| system.swap.used_bytes | Used swap in bytes
| system.swap.total_bytes| Total swap in bytes

#### load
| Name              | Description   |
|-------------------|---------------|
| system.load.load1 | System load averaged over 1 minute, high load value dependant on number of cpus in system
| system.load.load5 | System load averaged over 5 minute, high load value dependent on number of cpus in system
| system.load.load15 | System load averaged over 15 minute, high load value dependent on number of cpus in system
| system.load.load1_per_cpu | System load averaged over 1 minute normalized by cpu count, values > 1 means system may be overloaded
| system.load.load5_per_cpu | System load averaged over 5 minute normalized by cpu count, values > 1 means system may be overloaded
| system.load.load15_per_cpu| System load averaged over 15 minute normalized by cpu count, values > 1 means system may be overloaded

#### host
| Name              | Description   |
|-------------------|---------------|
| system.host.uptime | Host uptime in seconds 
| system.host.processes | Number of host processes 

## Usage examples

### Help output

```
Cross-platform System Check

Usage:
  system-check [flags]
  system-check [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -h, --help           help for system-check
  -i, --interval int   Interval in seconds over which to calculate cpu usage metrics (default 3)

Use "system-check [command] --help" for more information about a command.
```

### Environment variables
|Argument                       |Environment Variable                 |
|-------------------------------|-------------------------------------|
|--interval                     |SYSTEM_CHECK_INTERVAL                |



## Configuration
### Asset registration

[Sensu Assets][11] are the best way to make use of this plugin. If you're not using an asset, please
consider doing so! If you're using sensuctl 5.13 with Sensu Backend 5.13 or later, you can use the
following command to add the asset:

```
sensuctl asset add sensu/system-check
```

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index][12].

### Check definition

```yml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: system-check
  namespace: default
spec:
  command: system-check
  subscriptions:
  - system
  runtime_assets:
  - sensu/system-check
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an Asset. If you would
like to compile and install the plugin from source or contribute to it, download the latest version
or create an executable script from this source.

From the local path of the system-check repository:

```
go build
```

## Additional notes

## Contributing

For more information about contributing to this plugin, see [Contributing][1].

[1]: https://github.com/sensu/system-check
[2]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[3]: https://github.com/sensu/sensu-plugin-sdk
[4]: https://github.com/sensu-plugins/community/blob/master/PLUGIN_STYLEGUIDE.md
[5]: https://github.com/sensu/check-plugin-template/blob/master/.github/workflows/release.yml
[6]: https://github.com/sensu/check-plugin-template/actions
[7]: https://docs.sensu.io/sensu-go/latest/reference/checks/
[8]: https://github.com/sensu/check-plugin-template/blob/master/main.go
[9]: https://bonsai.sensu.io/
[10]: https://github.com/sensu/sensu-plugin-tool
[11]: https://docs.sensu.io/sensu-go/latest/reference/assets/
[12]: https://bonsai.sensu.io/assets/sensu/system-check
