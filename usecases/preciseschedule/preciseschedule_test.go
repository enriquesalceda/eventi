package preciseschedule_test

import (
	"errors"
	"eventi/modules/scheduler"
	"eventi/usecases/preciseschedule"
	"github.com/aws/aws-sdk-go/aws"
	awsscheduler "github.com/aws/aws-sdk-go/service/scheduler"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPreciseSchedule(t *testing.T) {
	t.Run("Successful precise scheduling", func(t *testing.T) {
		eventiScheduler := scheduler.New()
		fakeEventBridge := newFakeEventBridgeSpec("precise-schedule-arn")
		result := preciseschedule.New(eventiScheduler, fakeEventBridge)(
			preciseschedule.Input{
				ClientToken: "bd6dccce-e27a-11ee-87f6-e7571459c4c5",
				Description: "This a schedule description",
				GroupName:   "my-group",
				Name:        "bd6dccce-e27a-11ee-87f6-e7571459c4c5",
				At:          "at(2000-01-01T00:00:00)",
				Target:      "target-arn",
				DeadLetter:  "dead-letter-arn",
				RoleArn:     "arn:aws:iam::123456789012:role/schedulerroletoinvoketarget",
			},
		)

		require.NoError(t, result)
		eventBridgeScheduleInput, err := eventiScheduler.ToAWS()
		require.NoError(t, err)
		require.Equal(
			t,
			&awsscheduler.CreateScheduleInput{
				ClientToken:           aws.String("bd6dccce-e27a-11ee-87f6-e7571459c4c5"),
				Description:           aws.String("This a schedule description"),
				GroupName:             aws.String("my-group"),
				Name:                  aws.String("bd6dccce-e27a-11ee-87f6-e7571459c4c5"),
				ActionAfterCompletion: aws.String("DELETE"),
				FlexibleTimeWindow: &awsscheduler.FlexibleTimeWindow{
					Mode: aws.String("OFF"),
				},
				Target: &awsscheduler.Target{
					Arn: aws.String("target-arn"),
					DeadLetterConfig: &awsscheduler.DeadLetterConfig{
						Arn: aws.String("dead-letter-arn"),
					},
					RoleArn: aws.String("arn:aws:iam::123456789012:role/schedulerroletoinvoketarget"),
				},
				ScheduleExpression: aws.String("at(2000-01-01T00:00:00)"),
			},
			eventBridgeScheduleInput,
		)
	})

	t.Run("errors when client token is not present", func(t *testing.T) {
		eventiScheduler := scheduler.New()
		fakeEventBridge := newFakeEventBridgeSpec("precise-schedule-arn")
		result := preciseschedule.New(eventiScheduler, fakeEventBridge)(
			preciseschedule.Input{
				Description: "This a schedule description",
				GroupName:   "my-group",
				Name:        "bd6dccce-e27a-11ee-87f6-e7571459c4c5",
				At:          "at(2000-01-01T00:00:00)",
				Target:      "target-arn",
				DeadLetter:  "dead-letter-arn",
				RoleArn:     "arn:aws:iam::123456789012:role/schedulerroletoinvoketarget",
			},
		)

		require.Error(t, result)
		require.ErrorContains(
			t,
			errors.New("InvalidParameter: 1 validation error(s) found.\n- minimum field size of 1, CreateScheduleInput.ClientToken.\n"),
			result.Error(),
		)
	})

	t.Run("errors when group name is not present", func(t *testing.T) {
		eventiScheduler := scheduler.New()
		fakeEventBridge := newFakeEventBridgeSpec("precise-schedule-arn")
		result := preciseschedule.New(eventiScheduler, fakeEventBridge)(
			preciseschedule.Input{
				ClientToken: "bd6dccce-e27a-11ee-87f6-e7571459c4c5",
				Description: "This a schedule description",
				Name:        "bd6dccce-e27a-11ee-87f6-e7571459c4c5",
				At:          "at(2000-01-01T00:00:00)",
				Target:      "target-arn",
				DeadLetter:  "dead-letter-arn",
				RoleArn:     "arn:aws:iam::123456789012:role/schedulerroletoinvoketarget",
			},
		)

		require.Error(t, result)
		require.ErrorContains(
			t,
			errors.New("InvalidParameter: 1 validation error(s) found.\n- minimum field size of 1, CreateScheduleInput.GroupName.\n"),
			result.Error(),
		)
	})

	t.Run("errors when target arn is not present", func(t *testing.T) {
		eventiScheduler := scheduler.New()
		fakeEventBridge := newFakeEventBridgeSpec("precise-schedule-arn")
		result := preciseschedule.New(eventiScheduler, fakeEventBridge)(
			preciseschedule.Input{
				ClientToken: "bd6dccce-e27a-11ee-87f6-e7571459c4c5",
				Description: "This a schedule description",
				GroupName:   "my-group",
				Name:        "bd6dccce-e27a-11ee-87f6-e7571459c4c5",
				At:          "at(2000-01-01T00:00:00)",
				DeadLetter:  "dead-letter-arn",
				RoleArn:     "arn:aws:iam::123456789012:role/schedulerroletoinvoketarget",
			},
		)

		require.Error(t, result)
		require.ErrorContains(
			t,
			errors.New("InvalidParameter: 1 validation error(s) found.\n- minimum field size of 1, CreateScheduleInput.Target.Arn.\n"),
			result.Error(),
		)
	})
}

type fakeEventBridgeClient struct {
	arn string
}

func newFakeEventBridgeSpec(arn string) *fakeEventBridgeClient {
	return &fakeEventBridgeClient{arn: arn}
}

func (f fakeEventBridgeClient) CreateSchedule(input *awsscheduler.CreateScheduleInput) (*awsscheduler.CreateScheduleOutput, error) {
	err := input.Validate()
	if err != nil {
		return &awsscheduler.CreateScheduleOutput{}, err
	}

	return &awsscheduler.CreateScheduleOutput{ScheduleArn: &f.arn}, nil
}
