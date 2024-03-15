package scheduler

type BaseScheduleInput struct {
	ClientToken string
	Description string
	GroupName   string
	Name        string
}

func New(clientToken string, description string, groupName string, name string) *BaseScheduleInput {
	return &BaseScheduleInput{
		ClientToken: clientToken,
		Description: description,
		GroupName:   groupName,
		Name:        name,
	}
}
