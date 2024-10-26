package matches

import (
	"context"
	"sync"
	"time"

	"github.com/Sadzeih/valcompbot/ent"
	"github.com/Sadzeih/valcompbot/ent/scheduledmatch"
)

type Scheduler struct {
	ent *ent.Client
}

func NewScheduler(ent *ent.Client) *Scheduler {
	return &Scheduler{
		ent: ent,
	}
}

func (s *Scheduler) Start() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		timer := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-timer.C:

			}
		}
	}()

	wg.Wait()
}

func (s *Scheduler) getMatches(ctx context.Context) (ent.ScheduledMatches, error) {
	return s.ent.ScheduledMatch.Query().
		Where(
			scheduledmatch.PostedAtIsNil(),
		).
		All(ctx)
}
