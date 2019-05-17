package structure

import (
	"os"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/getninjas/spot-advisor/config"
)

// time to a instance on demmand stay alive (45 min)
const timeToLive int64 = 2700

// QueryData return struct to query the data to the AWS
func QueryData(dimension *cloudwatch.Dimension) *cloudwatch.MetricDataQuery {

	metricQuerieData := &cloudwatch.MetricDataQuery{
		Id: aws.String(
			IDMetric()),
		MetricStat: &cloudwatch.MetricStat{
			Unit: aws.String(
				Unit()),
			Period: DataPoint(),
			Stat:   Statistic(),
			Metric: &cloudwatch.Metric{
				Dimensions: []*cloudwatch.Dimension{dimension},
				MetricName: MetricName(),
				Namespace:  NameSpace(),
			},
		},
	}
	return metricQuerieData
}

// CloudDimension return struct to query the data to the AWS
func CloudDimension(spootID string) *cloudwatch.Dimension {
	configFile := Fleettype()
	result := &cloudwatch.Dimension{
		Name:  aws.String(configFile),
		Value: aws.String(spootID),
	}
	return result
}

// GetRegion verify if have a region
func GetRegion() {
	if len(os.Getenv("AWS_REGION")) <= 1 {
		os.Setenv("AWS_REGION", "us-east-1")
	}
}

func getSpotConfig() string {
	spot := os.Getenv("SPOT_CONFIG")
	if len(spot) <= 1 {
		panic("env SPOT_CONFIG not found")
	}
	return os.Getenv("SPOT_CONFIG")
}

// DataPoint get the total of past to look. This config come from file
func DataPoint() *int64 {
	dpoint := config.DataPointConfig()
	return &dpoint
}

// Statistic return the average metric type from config file
func Statistic() *string {
	result := config.StatisticConfig()
	return &result
}

// Fleettype return fleet type from config file
func Fleettype() string {
	return config.FleetTypeConfig()
}

// IDMetric show from config the type metric name
func IDMetric() string {
	return config.IDmetricConfig()
}

// ScanType get from config the scan type
func ScanType() *string {
	result := config.ScantypeConfig()
	return &result
}

// MetricName show the type of metric from config file
func MetricName() *string {
	result := config.MetricnameConfig()
	return &result
}

// NameSpace show the product name from AWS
func NameSpace() *string {
	result := config.NamespaceConfig()
	return &result
}

// Unit show the unit name from config file
func Unit() string {
	return config.UnitConfig()
}

//TimeConfig to look at past (in minutes)
func TimeConfig() int64 {
	return config.TimeToConfig() * 2
}

// Speed daemon from config file
func Speed() int64 {
	return config.SpeedConfig()
}

// Divider get from config the number to divider the request
func Divider() int64 {
	return config.DividerConfig()
}

// FallbackCheck verify if have a lock REMOVER!!!
func FallbackCheck(trigger bool, lock, scaling string) bool {
	if trigger == true && lock != scaling {
		return true
	} else if trigger == false && lock == scaling {
		return false
	}
	return false
}

// MakeLock return true if need to increase autoscaling
func MakeLock(lock int64) bool {
	return lock > 0
}

// CheckLock check if can be remove lock.
// Return True if can
func CheckLock(lock int64) bool {
	t := time.Now()
	timeNow := t.Unix()
	result := timeNow - lock
	return result >= timeToLive
}

// GetTrigger wil return true if need to enable fallback
func GetTrigger(data int64) bool {
	return data > 0
}

// IgnoreFleet will return true if the fleed not need to monitor
func IgnoreFleet(iamRole string) bool {
	list := config.FleetIgnored()

	for _, value := range list {
		getRegex, _ := regexp.Compile(value)
		result := getRegex.FindString(iamRole)
		if len(result) > 0 {
			return true
		}
	}
	return false
}
