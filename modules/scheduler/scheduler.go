package scheduler

import (
	"github.com/aws/aws-sdk-go/aws"
	awsscheduler "github.com/aws/aws-sdk-go/service/scheduler"
)

type BaseScheduleInput struct {
	ClientToken           string
	Description           string
	GroupName             string
	Name                  string
	ActionAfterCompletion string
	FlexibleTimeWindow    *FlexibleTimeWindow
	Target                *Target
}

type FlexibleTimeWindow struct {
	Mode string
}

type Target struct {
	Arn                 string
	DeadLetterConfigArn string
}

func New(clientToken string, description string, groupName string, name string) *BaseScheduleInput {
	return &BaseScheduleInput{ClientToken: clientToken, Description: description, GroupName: groupName, Name: name}
}

func (b *BaseScheduleInput) WithoutFlexibleTimeWindow() *BaseScheduleInput {
	b.FlexibleTimeWindow = &FlexibleTimeWindow{Mode: "OFF"}
	return b
}

func (b *BaseScheduleInput) DeleteAfterCompletion() *BaseScheduleInput {
	b.ActionAfterCompletion = "DELETE"
	return b
}

func (b *BaseScheduleInput) WithTarget(targetArn string, deadLetterConfigArn string) *BaseScheduleInput {
	b.Target = &Target{Arn: targetArn, DeadLetterConfigArn: deadLetterConfigArn}
	return b
}

func (b *BaseScheduleInput) ToAWS() *awsscheduler.CreateScheduleInput {
	return &awsscheduler.CreateScheduleInput{
		ClientToken:           aws.String(b.ClientToken),
		Description:           aws.String(b.Description),
		GroupName:             aws.String(b.GroupName),
		Name:                  aws.String(b.Name),
		ActionAfterCompletion: aws.String(b.ActionAfterCompletion),
		FlexibleTimeWindow: &awsscheduler.FlexibleTimeWindow{
			Mode: aws.String(b.FlexibleTimeWindow.Mode),
		},
		Target: &awsscheduler.Target{
			Arn: aws.String(b.Target.Arn),
			DeadLetterConfig: &awsscheduler.DeadLetterConfig{
				Arn: aws.String(b.Target.DeadLetterConfigArn),
			},
		},
	}
}
