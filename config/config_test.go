package config

import (
	"os"
	"testing"
)

func TestDataPointConfig(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"Teste", 42},
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
		{"Teste", "Average"},
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
		{"Teste", "FleetRequestId"},
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
		{"Teste", "metric"},
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
		{"Teste", "TimestampDescending"},
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
		{"Teste", "PendingCapacity"},
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
		{"Teste", "AWS/EC2Spot"},
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
		{"Teste", "Count"},
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
		{"Teste", 5},
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
		{"Teste", 10},
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
		{"Teste", 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DividerConfig(); got != tt.want {
				t.Errorf("DividerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
