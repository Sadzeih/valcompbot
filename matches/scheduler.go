package matches

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/Sadzeih/valcompbot/ent"
	"github.com/Sadzeih/valcompbot/ent/scheduledmatch"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type Scheduler struct {
	ent    *ent.Client
	reddit *reddit.Client
}

func NewScheduler(ent *ent.Client, redditClient *reddit.Client) *Scheduler {
	return &Scheduler{
		ent:    ent,
		reddit: redditClient,
	}
}

func (s *Scheduler) Start(ctx context.Context) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		timer := time.NewTicker(10 * time.Second)
		for range timer.C {
			matches, err := s.getMatches(ctx)
			if err != nil {
				log.Println(err)
				continue
			}
			for _, m := range matches {
				vm, err := GetMatch(m.MatchID)
				if err != nil {
					log.Println(err)
					continue
				}
				if vm.Info.Completed == "0" {
					continue
				}
				updateDone := vm.Info.Completed == "1" && m.DoneAt == nil
				if err := s.postMatch(ctx, m, vm, updateDone); err != nil {
					log.Println(err)
				}
			}
		}
	}()

	wg.Wait()
}

func (s *Scheduler) postMatch(ctx context.Context, m *ent.ScheduledMatch, vm *Match, updateDone bool) error {
	if err := PostMatch(ctx, vm, s.ent, s.reddit); err != nil {
		return err
	}
	update := m.Update()
	now := time.Now()
	update.SetPostedAt(now)
	if updateDone {
		update.SetDoneAt(now)
	}

	if _, err := update.Save(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Scheduler) getMatches(ctx context.Context) (ent.ScheduledMatches, error) {
	return s.ent.ScheduledMatch.Query().
		Where(
			scheduledmatch.PostedAtIsNil(),
		).
		All(ctx)
}
