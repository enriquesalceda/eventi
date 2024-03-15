package scheduler_test

import (
	"eventi/modules/scheduler"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScheduler(t *testing.T) {
	t.Run("New schedule", func(t *testing.T) {
		t.Run("Base schedule input", func(t *testing.T) {
			baseScheduleInput := scheduler.New(
				"bd6dccce-e27a-11ee-87f6-e7571459c4c5",
				"This a schedule description",
				"my-group",
				"bd6dccce-e27a-11ee-87f6-e7571459c4c5",
			)

			require.Equal(
				t,
				&scheduler.BaseScheduleInput{
					ClientToken: "bd6dccce-e27a-11ee-87f6-e7571459c4c5",
					Description: "This a schedule description",
					GroupName:   "my-group",
					Name:        "bd6dccce-e27a-11ee-87f6-e7571459c4c5",
				},
				baseScheduleInput,
			)
		})

		t.Run("Create a precise schedule without flexible time window", func(t *testing.T) {
			baseScheduleInput := scheduler.New(
				"bd6dccce-e27a-11ee-87f6-e7571459c4c5",
				"This a schedule description",
				"my-group",
				"bd6dccce-e27a-11ee-87f6-e7571459c4c5",
			).WithoutFlexibleTimeWindow()

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
				},
				baseScheduleInput,
			)
		})

		t.Run("Create a precise schedule without flexible time window deleting after completion", func(t *testing.T) {
			baseScheduleInput := scheduler.
				New(
					"bd6dccce-e27a-11ee-87f6-e7571459c4c5",
					"This a schedule description",
					"my-group",
					"bd6dccce-e27a-11ee-87f6-e7571459c4c5",
				).
				WithoutFlexibleTimeWindow().
				DeleteAfterCompletion()

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
				},
				baseScheduleInput,
			)
		})

		t.Run("Create a precise schedule without flexible time window deleting after completion with target", func(t *testing.T) {
			baseScheduleInput := scheduler.
				New(
					"bd6dccce-e27a-11ee-87f6-e7571459c4c5",
					"This a schedule description",
					"my-group",
					"bd6dccce-e27a-11ee-87f6-e7571459c4c5",
				).
				WithoutFlexibleTimeWindow().
				DeleteAfterCompletion().
				WithTarget("target-arn", "dead-letter-arn")

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
				},
				baseScheduleInput,
			)
		})
	})
}
