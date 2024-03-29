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
	ScheduleExpression    string
	RoleArn               string
}

type FlexibleTimeWindow struct {
	Mode string
}

type Target struct {
	Arn                 string
	DeadLetterConfigArn string
}

func New() *BaseScheduleInput { return &BaseScheduleInput{} }

func (b *BaseScheduleInput) WithGroupName(groupName string) *BaseScheduleInput {
	b.GroupName = groupName
	return b
}

func (b *BaseScheduleInput) WithDescription(description string) *BaseScheduleInput {
	b.Description = description
	return b
}

func (b *BaseScheduleInput) WithName(name string) *BaseScheduleInput {
	b.Name = name
	return b
}

func (b *BaseScheduleInput) WithClientToken(clientToken string) *BaseScheduleInput {
	b.ClientToken = clientToken
	return b
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

func (b *BaseScheduleInput) At(scheduleExpression string) *BaseScheduleInput {
	b.ScheduleExpression = scheduleExpression
	return b
}

func (b *BaseScheduleInput) WithRoleArn(roleArn string) *BaseScheduleInput {
	b.RoleArn = roleArn
	return b
}

func (b *BaseScheduleInput) ToAWS() (*awsscheduler.CreateScheduleInput, error) {
	createScheduleInput := &awsscheduler.CreateScheduleInput{
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
			RoleArn: aws.String(b.RoleArn),
		},
		ScheduleExpression: aws.String(b.ScheduleExpression),
	}

	err := createScheduleInput.Validate()
	if err != nil {
		return &awsscheduler.CreateScheduleInput{}, err
	}

	return createScheduleInput, nil
}
