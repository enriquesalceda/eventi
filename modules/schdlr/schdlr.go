package schdlr

import awsscheduler "github.com/aws/aws-sdk-go/service/scheduler"

type Scheduler interface {
	CreateSchedule(input *awsscheduler.CreateScheduleInput) (*awsscheduler.CreateScheduleOutput, error)
}
