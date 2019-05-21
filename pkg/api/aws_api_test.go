package api

import (
	"reflect"
	"testing"

	"github.com/getninjas/spot-ninja/pkg/structure"
)

func Test_sessiontStartEc2(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "2016-11-15"},
	}

	client := sessiontStartEc2()
	apiVersion := client.APIVersion

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := apiVersion; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sessiontStartEc2() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_sessionStartCloudWatch(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "2010-08-01"},
	}

	client := sessionStartCloudWatch()
	apiVersion := client.APIVersion

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := apiVersion; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sessiontStartEc2() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_sessionStartAutoScaling(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "2011-01-01"},
	}

	client := sessionStartAutoScaling()
	apiVersion := client.APIVersion

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := apiVersion; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sessiontStartEc2() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestListInstanceIds(t *testing.T) {
	structure.GetRegion()
	tests := []struct {
		name string
		want bool
	}{
		{"Test", true},
	}

	scalingName := "eks-training-room"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listInstanceIds(scalingName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListInstanceIds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sendSqsMessage(t *testing.T) {
	structure.GetRegion()
	instanceID := "i-42424242424242424242"

	tests := []struct {
		name string
		want bool
	}{
		{"Test", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sendSqsMessage(instanceID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sendSqsMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
