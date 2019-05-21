package config

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

var errMsgVar = errors.New("getenv: env variable not exist")

func getStringEnv(key string) (string, error) {
	env := os.Getenv(key)
	if len(env) == 0 {
		return env, errMsgVar
	}
	return env, nil
}

func getIntEnv(key string) (int64, error) {
	env := os.Getenv(key)
	if len(env) == 0 {
		return 0, errMsgVar
	}
	envResult, err := strconv.ParseInt(env, 10, 64)
	if err != nil {
		return 0, err
	}
	return envResult, nil
}

func getBoolEnv(key string) (bool, error) {
	env := os.Getenv(key)
	if len(env) == 0 {
		return false, errMsgVar
	}
	envResult, err := strconv.ParseBool(env)
	if err != nil {
		return false, err
	}
	return envResult, nil
}

func getArryStringEnv(key string) ([]string, error) {
	env := os.Getenv(key)
	if len(env) == 0 {
		return nil, errMsgVar
	}
	envResult := strings.FieldsFunc(env, func(r rune) bool {
		if r == ',' {
			return true
		}
		return false
	})

	return envResult, nil
}

// DataPointConfig is How much statistical data that CloudWatch will return
func DataPointConfig() int64 {
	env, _ := getIntEnv("DATA_POINT_CONFIG")
	if env == 0 {
		env = 30
	}
	return env
}

// StatisticConfig is the Statistic type to check
func StatisticConfig() string {
	env, _ := getStringEnv("STATISTIC_CONFIG")
	if len(env) == 0 {
		env = "Average"
	}
	return env
}

// FleetTypeConfig is the Fleet request type to query on AWS
func FleetTypeConfig() string {
	env, _ := getStringEnv("FLEET_TYPE_CONFIG")
	if len(env) <= 1 {
		env = "FleetRequestId"
	}
	return env
}

// IDmetricConfig is the ID name to query on CloudWatch
func IDmetricConfig() string {
	env, _ := getStringEnv("ID_METRIC_CONFIG")
	if len(env) <= 1 {
		env = "metric"
	}
	return env
}

// ScantypeConfig are the Scan type
func ScantypeConfig() string {
	env, _ := getStringEnv("SCAN_TYPE_CONFIG")
	if len(env) <= 1 {
		env = "TimestampDescending"
	}
	return env
}

// MetricnameConfig are the Name of metric
func MetricnameConfig() string {
	env, _ := getStringEnv("METRIC_NAME_CONFIG")
	if len(env) <= 1 {
		env = "PendingCapacity"
	}
	return env
}

// NamespaceConfig are the AWS service name
func NamespaceConfig() string {
	env, _ := getStringEnv("NAMESPACE_CONFIG")
	if len(env) <= 1 {
		env = "AWS/EC2Spot"
	}
	return env
}

// UnitConfig are the time unit
func UnitConfig() string {
	env, _ := getStringEnv("UNIT_CONFIG")
	if len(env) <= 1 {
		env = "Count"
	}
	return env
}

// TimeToConfig to look at past (in minutes)
func TimeToConfig() int64 {
	env, _ := getIntEnv("TIME_TO_CONFIG")
	if env <= 1 {
		env = 5
	}
	return env
}

// SpeedConfig Daemon speed(seconds)
func SpeedConfig() int64 {
	env, _ := getIntEnv("SPEED_CONFIG")
	if env <= 1 {
		env = 10
	}
	return env
}

// DividerConfig Divider number of the missing capacity result
func DividerConfig() int64 {
	env, _ := getIntEnv("DIVIDER_CONFIG")
	if env == 0 {
		env = 4
	}
	return env
}

// FleetIgnored return a fleet array list to be ignored by the spot-ninja
func FleetIgnored() []string {
	env, _ := getArryStringEnv("FLEET_IGNORED")
	return env
}

// EnableEventsOnSqs enable or disable sqs post events
func EnableEventsOnSqs() bool {
	env, _ := getBoolEnv("ENABLE_EVENTS_ON_SQS")
	return env
}

// SqsURL return sqs url to post events
func SqsURL() string {
	env, _ := getStringEnv("SQS_URL")
	return env
}

// ScallingPrefix will get the prefix to find the right scalling name
func ScallingPrefix() string {
	env, _ := getStringEnv("PREFIX")
	if len(env) <= 1 {
		env = "ecs-"
	}
	return env
}

// TimeToLive returns (in seconds) the minimum time an instance should
// be alive. The minimum is 60 seconds, default 15 min.
func TimeToLive() int64 {
	env, _ := getIntEnv("TIMETO_LIVE")
	if env == 0 {
		env = 900
	}
	return env
}
