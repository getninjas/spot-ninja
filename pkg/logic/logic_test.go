package logic

import (
	"testing"
)

func Test_decreaseLogic(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Test", true},
	}
	scaling := "exampleScaling"
	var count int64
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decreaseLogic(scaling, count); got != tt.want {
				t.Errorf("decreaseLogic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scaleUp(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Test", false},
	}
	scaling := "exampleScaling"
	var count int64
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scaleUp(scaling, count); got != tt.want {
				t.Errorf("scaleUp() = %v, want %v", got, tt.want)
			}

		})
	}
}
