package structure

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func TestQueryData(t *testing.T) {
	tests := []struct {
		name string
		want *cloudwatch.MetricDataQuery
	}{
		{"Test", &cloudwatch.MetricDataQuery{
			Id: aws.String("metric"),
			MetricStat: &cloudwatch.MetricStat{
				Metric: &cloudwatch.Metric{
					Dimensions: []*cloudwatch.Dimension{
						{
							Name:  aws.String("FleetRequestId"),
							Value: aws.String("42"),
						},
					},
					MetricName: aws.String("PendingCapacity"),
					Namespace:  aws.String("AWS/EC2Spot"),
				},
				Period: aws.Int64(30),
				Stat:   aws.String("Average"),
				Unit:   aws.String("Count"),
			},
		},
		},
	}
	query := CloudDimension("42")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryData(query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRegion(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"test", "us-east-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetRegion()
			if got := os.Getenv("AWS_REGION"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSpotConfig(t *testing.T) {
	envTest := "42"
	tests := []struct {
		name string
		want string
	}{
		{"test", envTest},
	}
	os.Setenv("SPOT_CONFIG", envTest)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSpotConfig(); got != tt.want {
				t.Errorf("getSpotConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFleettype(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"test", "FleetRequestId"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fleettype(); got != tt.want {
				t.Errorf("Fleettype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIDMetric(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"test", "metric"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IDMetric(); got != tt.want {
				t.Errorf("IDMetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnit(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"test", "Count"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unit(); got != tt.want {
				t.Errorf("Unit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeConfig(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"test", 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeConfig(); got != tt.want {
				t.Errorf("TimeConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpeed(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"test", 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Speed(); got != tt.want {
				t.Errorf("Speed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivider(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"test", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divider(); got != tt.want {
				t.Errorf("Divider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTrigger(t *testing.T) {
	tests := []struct {
		name string
		data int64
		want bool
	}{
		{"test", 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTrigger(tt.data); got != tt.want {
				t.Errorf("GetTrigger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckLock(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Test", false},
	}
	time := time.Now()
	timeNow := time.Unix()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckLock(timeNow); got != tt.want {
				t.Errorf("CheckLock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeLock(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Test", true},
	}

	value := int64(42)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeLock(value); got != tt.want {
				t.Errorf("MakeLock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFallbackCheck(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Test", false},
	}
	trigger := true
	lock := "olamundo"
	scalling := "olamundo"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FallbackCheck(trigger, lock, scalling); got != tt.want {
				t.Errorf("FallbackCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIgnoreFleet(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Test", false},
	}

	iamRole := "spotfleet-fake"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IgnoreFleet(iamRole); got != tt.want {
				t.Errorf("IgnoreFleet() = %v, want %v", got, tt.want)
			}
		})
	}
}
