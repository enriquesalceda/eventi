package scheduler_test

import (
	"errors"
	"eventi/modules/scheduler"
	"github.com/aws/aws-sdk-go/aws"
	awsscheduler "github.com/aws/aws-sdk-go/service/scheduler"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScheduler(t *testing.T) {
	t.Run("New Precision Schedule", func(t *testing.T) {
		t.Run("with name, description, group name, client token, target, delete after completion and without flexible time window", func(t *testing.T) {
			baseScheduleInput := scheduler.
				New().
				WithName("bd6dccce-e27a-11ee-87f6-e7571459c4c5").
				WithDescription("This a schedule description").
				WithGroupName("my-group").
				WithClientToken("bd6dccce-e27a-11ee-87f6-e7571459c4c5").
				WithoutFlexibleTimeWindow().
				DeleteAfterCompletion().
				WithTarget(
					"target-arn",
					"dead-letter-arn",
				).
				At("2000-01-01T00:00:00").
				WithRoleArn("arn:aws:iam::123456789012:role/schedulerroletoinvoketarget")

			require.Equal(
				t,
				&scheduler.BaseScheduleInput{
					ClientToken: "bd6dccce-e27a-11ee-87f6-e7571459c4c5",
					Description: "This a schedule description",
					GroupName:   "my-group",
					Name:        "bd6dccce-e27a-11ee-87f6-e7571459c4c5",
					FlexibleTimeWindow: &scheduler.FlexibleTimeWindow{
						Mode: "OFF",
					},
					ActionAfterCompletion: "DELETE",
					Target: &scheduler.Target{
						Arn:                 "target-arn",
						DeadLetterConfigArn: "dead-letter-arn",
					},
					ScheduleExpression: "2000-01-01T00:00:00",
					RoleArn:            "arn:aws:iam::123456789012:role/schedulerroletoinvoketarget",
				},
				baseScheduleInput,
			)
		})

		t.Run("ToAWS", func(t *testing.T) {
			t.Run("valid", func(t *testing.T) {
				precisionAwsBaseScheduleInput, err := scheduler.
					New().
					WithName("bd6dccce-e27a-11ee-87f6-e7571459c4c5").
					WithDescription("This a schedule description").
					WithGroupName("my-group").
					WithClientToken("bd6dccce-e27a-11ee-87f6-e7571459c4c5").
					WithoutFlexibleTimeWindow().
					DeleteAfterCompletion().
					WithTarget(
						"target-arn",
						"dead-letter-arn",
					).
					At("2000-01-01T00:00:00").
					WithRoleArn("arn:aws:iam::123456789012:role/schedulerroletoinvoketarget").
					ToAWS()

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
						ScheduleExpression: aws.String("2000-01-01T00:00:00"),
					},
					precisionAwsBaseScheduleInput,
				)
			})

			t.Run("invalid", func(t *testing.T) {
				_, err := scheduler.
					New().
					WithName("bd6dccce-e27a-11ee-87f6-e7571459c4c5").
					WithDescription("This a schedule description").
					WithGroupName("my-group").
					WithClientToken("bd6dccce-e27a-11ee-87f6-e7571459c4c5").
					WithoutFlexibleTimeWindow().
					DeleteAfterCompletion().
					WithTarget(
						"target-arn",
						"dead-letter-arn",
					).
					At("2000-01-01T00:00:00").
					ToAWS()

				require.Error(t, err)
				require.ErrorContains(
					t,
					errors.New(
						"InvalidParameter: 1 validation error(s) found.\n- minimum field size of 1, CreateScheduleInput.Target.RoleArn.\n",
					),
					err.Error(),
				)
			})
		})
	})
}
