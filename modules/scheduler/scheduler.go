package scheduler

type BaseScheduleInput struct {
	ClientToken        string
	Description        string
	GroupName          string
	Name               string
	FlexibleTimeWindow *FlexibleTimeWindow
}

type FlexibleTimeWindow struct {
	Mode string
}

func New(clientToken string, description string, groupName string, name string) *BaseScheduleInput {
	return &BaseScheduleInput{
		ClientToken: clientToken,
		Description: description,
		GroupName:   groupName,
		Name:        name,
	}
}

func (b *BaseScheduleInput) WithoutFlexibleTimeWindow() *BaseScheduleInput {
	b.FlexibleTimeWindow = &FlexibleTimeWindow{
		Mode: "OFF",
	}
	return b
}
