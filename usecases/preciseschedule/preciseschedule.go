package preciseschedule

import (
	"eventi/modules/schdlr"
	"eventi/modules/scheduler"
)

type Input struct {
	ClientToken string
	Description string
	GroupName   string
	Name        string
	At          string
	Target      string
	DeadLetter  string
	RoleArn     string
}

func New(scheduler *scheduler.BaseScheduleInput, eventBridge schdlr.Scheduler) func(Input) error {
	return func(i Input) error {
		scheduler.
			WithName(i.Name).
			WithDescription(i.Description).
			WithGroupName(i.GroupName).
			WithClientToken(i.ClientToken).
			WithoutFlexibleTimeWindow().
			DeleteAfterCompletion().
			WithTarget(
				i.Target,
				i.DeadLetter,
			).
			WithRoleArn(i.RoleArn).
			At(i.At)

		eventBridgeScheduleInput, err := scheduler.ToAWS()
		if err != nil {
			return err
		}

		_, err = eventBridge.CreateSchedule(eventBridgeScheduleInput)
		if err != nil {
			return err
		}

		return nil
	}
}
