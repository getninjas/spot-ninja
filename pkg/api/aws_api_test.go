package api

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

var (
	fleetID      = "sfr-424242-9101-4242-42-42424242"
	fleetStatus  = "active"
	IamFleetRole = "arn:aws:iam::42:role/42-production"
)

type mockEC2Svc struct {
	ec2iface.EC2API
}

func (m *mockEC2Svc) DescribeSpotFleetRequests(input *ec2.DescribeSpotFleetRequestsInput) (*ec2.DescribeSpotFleetRequestsOutput, error) {
	return &ec2.DescribeSpotFleetRequestsOutput{
		SpotFleetRequestConfigs: []*ec2.SpotFleetRequestConfig{
			{
				SpotFleetRequestId:    &fleetID,
				SpotFleetRequestState: &fleetStatus,
				SpotFleetRequestConfig: &ec2.SpotFleetRequestConfigData{
					IamFleetRole: &IamFleetRole,
				},
			},
		},
	}, nil

}
func Test_FleetID(t *testing.T) {
	svc := &mockEC2Svc{}
	id := GetSpotID(svc)

	expected := []string{
		"sfr-424242-9101-4242-42-42424242",
	}

	if !reflect.DeepEqual(expected, id) {
		t.Errorf("expected %q to eq %q", expected, id)
	}
}

func TestScalingName(t *testing.T) {
	svc := &mockEC2Svc{}
	id := ScalingName(fleetID, svc)

	expected := "arn:aws:iam::42:role/42-production"

	if !reflect.DeepEqual(expected, id) {
		t.Errorf("expected %q to eq %q", expected, id)
	}
}
