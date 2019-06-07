# spot-ninja

[![Go Report Card](https://goreportcard.com/badge/github.com/getninjas/spot-ninja)](https://goreportcard.com/report/github.com/getninjas/spot-ninja)
[![Build Status](https://travis-ci.org/getninjas/spot-ninja.svg?branch=master)](https://travis-ci.org/getninjas/spot-ninja)
[![GoDoc](https://godoc.org/github.com/getninjas/spot-ninja?status.svg)](https://godoc.org/github.com/getninjas/spot-ninja)

Available translations:

- [pt-BR](./README_PT_BR.md)

This software is responsible for all the intelligence in increasing and decreasing AutoScalingGroups size based on Spotfleet health.

## License

spot ninja is released under the terms of the Apache License. See LICENSE file for more information or see [apache](https://www.apache.org/licenses/LICENSE-2.0).

## Requirements

- Docker >=v1.13.1
- Docker Compose >=1.23.1
- Go >= 1.11.2
- aws-sdk-go Module
- yaml.v2 Module
- Polices:
  - CloudWatchReadOnlyAccess
  - AmazonEC2SpotFleetAutoscaleRole
  - AutoScalingConsoleFullAccess

## Architecture

![spot-ninja](assets/spot_EN-US.png)

## ENVs configuration

To configure the spot-ninja, you need to configure some environment variables. The variables are:

- AWS_REGION
  - Default us-east-1
- DATA_POINT_CONFIG
  - Default 30
- STATISTIC_CONFIG
  - Default Average
- FLEET_TYPE_CONFIG
  - Default FleetRequestId
- ID_METRIC_CONFIG
  - Default metric
- SCAN_TYPE_CONFIG
  - Default TimestampDescending
- METRIC_NAME_CONFIG
  - Default PendingCapacity
- NAMESPACE_CONFIG
  - Default AWS/EC2Spot
- UNIT_CONFIG
  - Default Count
- TIME_TO_CONFIG
  - Default 5
- SPEED_CONFIG
  - Default 10
- DIVIDER_CONFIG
  - Default 4
- FLEET_IGNORED
  - Default none
- ENABLE_EVENTS_ON_SQS
  - Default false
- SQS_URL
  - Default none
- PREFIX
  - Default ecs-
- TIME_TO_LIVE
  - In seconds Default 15 min.

## Installing with docker-compose

1 - Build

```bash
docker-compose build
```

2 - Run spot-ninja

```bash
docker-compose up
```

## Project Structure

- cmd/
  - main spot-ninja
- config/
  - general configs to spot ninja
- pkg/
  - general libs

### References

- [gopherize](https://www.gopherize.me/)
- [cloudcraft](https://cloudcraft.co/)
- [project-layout](https://github.com/golang-standards/project-layout)
- [regex101](https://regex101.com/r/FwSMp7/1/)
- [sdk-for-go](https://docs.aws.amazon.com/sdk-for-go/api/)
