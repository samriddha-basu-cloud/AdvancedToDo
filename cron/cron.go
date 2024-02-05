package cron

import (
	"github.com/robfig/cron/v3"
)

func ScheduleTaskPriorityUpdate() {
	c := cron.New()

	// Schedule the task priority update cron job
	c.AddFunc("0 0 * * *", func() {
		// Implement logic to update task priorities based on due dates
		// Example: db.UpdateTaskPriorities()
	})

	c.Start()
}

func ScheduleVoiceCalling() {
	c := cron.New()

	// Schedule the voice calling cron job
	c.AddFunc("0 0 * * *", func() {
		// Implement logic to perform voice calling using Twilio
		// Example: db.InitiateVoiceCalls()
	})

	c.Start()
}
