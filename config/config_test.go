package config

import (
	"os"
	"strings"
	"testing"
)

func TestDataPointConfig(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"Test", 30},
	}
	os.Setenv("DATAPOINTCONFIG", "42")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DataPointConfig(); got != tt.want {
				t.Errorf("DataPointConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatisticConfig(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "Average"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StatisticConfig(); got != tt.want {
				t.Errorf("StatisticConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFleetTypeConfig(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "FleetRequestId"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FleetTypeConfig(); got != tt.want {
				t.Errorf("FleetTypeConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIDmetricConfig(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "metric"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IDmetricConfig(); got != tt.want {
				t.Errorf("IDmetricConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScantypeConfig(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "TimestampDescending"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ScantypeConfig(); got != tt.want {
				t.Errorf("ScantypeConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetricnameConfig(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "PendingCapacity"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MetricnameConfig(); got != tt.want {
				t.Errorf("MetricnameConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNamespaceConfig(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "AWS/EC2Spot"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NamespaceConfig(); got != tt.want {
				t.Errorf("NamespaceConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitConfig(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "Count"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnitConfig(); got != tt.want {
				t.Errorf("UnitConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeToConfig(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"Test", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeToConfig(); got != tt.want {
				t.Errorf("TimeToConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpeedConfig(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"Test", 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpeedConfig(); got != tt.want {
				t.Errorf("SpeedConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDividerConfig(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"Test", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DividerConfig(); got != tt.want {
				t.Errorf("DividerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getStringEnv(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "example"},
	}
	os.Setenv("Test", "example")

	result, _ := getStringEnv("Test")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := result; got != tt.want {
				t.Errorf("getStringEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIntEnv(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"Test", 42},
	}
	os.Setenv("Test", "42")

	result, _ := getIntEnv("Test")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := result; got != tt.want {
				t.Errorf("getIntEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBoolEnv(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Test", true},
	}
	os.Setenv("Test", "true")

	result, _ := getBoolEnv("Test")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := result; got != tt.want {
				t.Errorf("getBoolEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getArryStringEnv(t *testing.T) {
	strwant := strings.FieldsFunc("42,", func(r rune) bool {
		if r == ',' {
			return true
		}
		return false
	})

	os.Setenv("Test", "42")
	result, _ := getArryStringEnv("Test")

	if result[0] != strwant[0] {
		t.Errorf("getArryStringEnv() = %v, want %v", result[0], strwant[0])
	}
}
