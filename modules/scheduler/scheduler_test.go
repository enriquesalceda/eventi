package scheduler_test

import (
	"eventi/modules/scheduler"
	"github.com/aws/aws-sdk-go/aws"
	awsscheduler "github.com/aws/aws-sdk-go/service/scheduler"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScheduler(t *testing.T) {
	t.Run("NewScheduler", func(t *testing.T) {
		schdl := scheduler.New(
			"bd6dccce-e27a-11ee-87f6-e7571459c4c5",
			"This a schedule description",
			"my-group",
			"bd6dccce-e27a-11ee-87f6-e7571459c4c5",
		)

		require.Equal(
			t,
			&awsscheduler.CreateScheduleInput{
				ClientToken: aws.String("bd6dccce-e27a-11ee-87f6-e7571459c4c5"),
				Description: aws.String("This a schedule description"),
				GroupName:   aws.String("my-group"),
				Name:        aws.String("bd6dccce-e27a-11ee-87f6-e7571459c4c5"),
			},
			schdl,
		)
	})
}
