package api

import (
	"fmt"
	"regexp"
	"time"

	"github.com/getninjas/spot-ninja/config"

	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/getninjas/spot-ninja/pkg/structure"
)

// waitTimeDrain will controll how many (seconds) time to wait to drain result
const waitTimeDrain int = 300

// Create a session on AWS EC2
func sessiontStartEc2() *ec2.EC2 {
	return ec2.New(session.New())
}

// Create a session on AWS CloudWatch
func sessionStartCloudWatch() *cloudwatch.CloudWatch {
	return cloudwatch.New(session.New())
}

// Create a session on Auto Scaling
func sessionStartAutoScaling() *autoscaling.AutoScaling {
	return autoscaling.New(session.New())
}

// Create a session on SQS
func sessionSqs() *sqs.SQS {
	return sqs.New(session.New())
}

// GetSpotID show the spot IDs from AWS account
func GetSpotID() []string {
	spotID := []string{}
	client := sessiontStartEc2()
	input := &ec2.DescribeSpotFleetRequestsInput{}

	dataResult, err := client.DescribeSpotFleetRequests(input)
	if err != nil {
		fmt.Println("error, describe-spotfleetrequest, ", err)
	}

	for _, loop := range dataResult.SpotFleetRequestConfigs {
		if *loop.SpotFleetRequestState == "active" {
			iamFleet := *loop.SpotFleetRequestConfig.IamFleetRole
			if !structure.IgnoreFleet(iamFleet) {
				spotID = append(spotID, *loop.SpotFleetRequestId)
			}
		}
	}
	return spotID
}

// QueryDataRequest request all necessary data from AWS
func QueryDataRequest(data *cloudwatch.MetricDataQuery) int64 {
	var result float64
	client := sessionStartCloudWatch()

	timeNow := time.Now()
	timePrev := timeNow.Add(time.Duration(-structure.TimeConfig()) * time.Minute)

	request := &cloudwatch.GetMetricDataInput{
		EndTime:           aws.Time(timeNow),
		StartTime:         aws.Time(timePrev),
		MaxDatapoints:     structure.DataPoint(),
		ScanBy:            structure.ScanType(),
		MetricDataQueries: []*cloudwatch.MetricDataQuery{data},
	}

	dataResult, err := client.GetMetricData(request)
	if err != nil {
		fmt.Println("error, describe-metadata, ", err)
	}

	for index, loop := range dataResult.MetricDataResults {
		result = *loop.Values[index]
	}
	if result >= 1 {
		return int64(result) / structure.Divider()
	}
	return 0
}

// ScalingName get the real name from
func ScalingName(id string) string {
	client := sessiontStartEc2()
	var scaling string
	input := &ec2.DescribeSpotFleetRequestsInput{
		SpotFleetRequestIds: []*string{
			aws.String(id),
		},
	}
	regex := "^arn:aws:iam::([\\d]+):role/[a-zA-Z]+-"
	getRegex, _ := regexp.Compile(regex)

	dataResult, err := client.DescribeSpotFleetRequests(input)
	if err != nil {
		fmt.Println("error, describe-spotfleet, ", err)
	}

	for _, loop := range dataResult.SpotFleetRequestConfigs {
		if *loop.SpotFleetRequestId == id {
			scaling = *loop.SpotFleetRequestConfig.IamFleetRole
		}
	}

	return getRegex.ReplaceAllString(scaling, "ecs-")
}

// IncreaseGroup increases the self-dimensioning group
func IncreaseGroup(scaling string, capacity int64) bool {

	client := sessionStartAutoScaling()
	var desired int64
	input := &autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: []*string{
			aws.String(scaling),
		},
	}

	describe, err := client.DescribeAutoScalingGroups(input)
	if err != nil {
		fmt.Println("error, describe-AutoScalingGroups, ", err)
	}
	for _, loop := range describe.AutoScalingGroups {
		desired = *loop.DesiredCapacity
	}

	inputScaling := &autoscaling.UpdateAutoScalingGroupInput{
		AutoScalingGroupName: aws.String(scaling),
		DesiredCapacity:      aws.Int64(desired + capacity),
	}

	_, err = client.UpdateAutoScalingGroup(inputScaling)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case autoscaling.ErrCodeScalingActivityInProgressFault:
				fmt.Println(autoscaling.ErrCodeScalingActivityInProgressFault, aerr.Error())
			case autoscaling.ErrCodeResourceContentionFault:
				fmt.Println(autoscaling.ErrCodeResourceContentionFault, aerr.Error())
			case autoscaling.ErrCodeServiceLinkedRoleFailure:
				fmt.Println(autoscaling.ErrCodeServiceLinkedRoleFailure, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
			return false
		}

		fmt.Println(err.Error())
		return false
	}
	return true
}

// DecreaseGroup decrease the self-dimensioning group
func DecreaseGroup(scaling string) bool {

	listInstanceIds(scaling)
	time.Sleep(time.Duration((waitTimeDrain)) * time.Second)

	client := sessionStartAutoScaling()

	inputScaling := &autoscaling.UpdateAutoScalingGroupInput{
		AutoScalingGroupName: aws.String(scaling),
		DesiredCapacity:      aws.Int64(0),
	}
	result, err := client.UpdateAutoScalingGroup(inputScaling)
	fmt.Println("fallback-return, ", result)
	if err != nil {
		fmt.Println(err)
	}
	return true
}

// listInstanceIds return all instance ID from a scaling group.
func listInstanceIds(scaling string) bool {

	client := sessionStartAutoScaling()
	input := &autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: []*string{
			aws.String(scaling),
		},
	}
	describe, err := client.DescribeAutoScalingGroups(input)
	if err != nil {
		fmt.Println("error, describe-AutoScalingGroups, ", err)
		return false
	}
	for _, loop := range describe.AutoScalingGroups {
		for _, ID := range loop.Instances {
			sendSqsMessage(*ID.InstanceId)
		}
	}
	return true
}

func sendSqsMessage(instanceID string) bool {
	if !config.EnableEventsOnSqs() {
		fmt.Println("Events on SQS disabled")
		return false
	}

	client := sessionSqs()
	qURL := config.SqsURL()

	message := &sqs.SendMessageInput{
		DelaySeconds: aws.Int64(5),
		MessageBody:  aws.String(`{ "detail-type": "Spot Ninja Notification Drain", "source": "spot-ninja", "detail": { "instance-id": ` + instanceID + `, "state": "terminated" } }`),
		QueueUrl:     &qURL,
	}
	_, err := client.SendMessage(message)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
