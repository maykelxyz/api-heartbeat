package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
)

func New(timezone string) (*gocron.Scheduler, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, err
	}
	return gocron.NewScheduler(location), nil
}
