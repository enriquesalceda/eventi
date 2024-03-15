package preciseschedule_test

import (
	awsscheduler "github.com/aws/aws-sdk-go/service/scheduler"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPreciseSchedule(t *testing.T) {
	t.Run("Precise scheduling", func(t *testing.T) {
		useCase := New(
			"target-arn",
			"dead-letter-arn",
			newFakeEventBridgeSpec("precise-schedule-arn"),
		)

		result := useCase(
			precisescchedule.Input{
				ClientToken: "bd6dccce-e27a-11ee-87f6-e7571459c4c5",
				Description: "This a schedule description",
				GroupName:   "my-group",
				Name:        "bd6dccce-e27a-11ee-87f6-e7571459c4c5",
				At:          "at(2000-01-01T00:00:00)",
			},
		)

		require.NoError(t, result)
		require.Equal(t, "precise-schedule-arn", result)

	})
}

type fakeEventBridgeClient struct {
	arn string
}

func (f fakeEventBridgeClient) CreateSchedule(input *awsscheduler.CreateScheduleInput) (*awsscheduler.CreateScheduleOutput, error) {
	return &awsscheduler.CreateScheduleOutput{ScheduleArn: &f.arn}, nil
}

func newFakeEventBridgeSpec(arn string) *fakeEventBridgeClient {
	return &fakeEventBridgeClient{arn: arn}
}
