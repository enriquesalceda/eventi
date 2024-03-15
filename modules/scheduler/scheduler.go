package scheduler

import (
	"github.com/aws/aws-sdk-go/aws"
	awsscheduler "github.com/aws/aws-sdk-go/service/scheduler"
)

func New(clientToken string, description string, groupName string, name string) *awsscheduler.CreateScheduleInput {
	return &awsscheduler.CreateScheduleInput{
		ClientToken: aws.String(clientToken),
		Description: aws.String(description),
		GroupName:   aws.String(groupName),
		Name:        aws.String(name),
	}
}
