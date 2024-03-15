package scheduler

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
