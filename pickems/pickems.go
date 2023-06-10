package pickems

import (
	"context"
	"fmt"
	"github.com/Sadzeih/valcompbot/comments"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"log"
	"regexp"
)

const (
	noEventMd = `There is no pickems event running at the moment. Ask me later!`
	joinFmtMd = "[Join the subreddit pickems here.](https://vlr.gg/event/pickem/%d?group=valcomp)"
)

var (
	rankSyntax    = regexp.MustCompile(`^!rank$`)
	pickemsSyntax = regexp.MustCompile(`^!pickems$`)

	Event *int = nil
)

type Service struct {
	redditClient *reddit.Client
	commentsSub  *comments.Subscriber
}

func New(redditClient *reddit.Client, commentsSub *comments.Subscriber) *Service {
	return &Service{
		redditClient: redditClient,
		commentsSub:  commentsSub,
	}
}

func (s *Service) Run() error {
	defer s.commentsSub.Close()

	for {
		select {
		case err := <-s.commentsSub.Errors:
			if err != nil {
				return fmt.Errorf("error while reading comment: %w", err)
			}
		case comm := <-s.commentsSub.Comments:
			switch {
			case rankSyntax.MatchString(comm.Body):
				if Event == nil {
					s.NoEventRunning(comm)
					break
				}
				if err := s.RankComment(comm); err != nil {
					log.Print(err)
				}
				break
			case pickemsSyntax.MatchString(comm.Body):
				if Event == nil {
					s.NoEventRunning(comm)
					break
				}
				s.PickemsComment(comm)
				break
			default:
				// no command was recognized, do nothing
			}
		}
	}
}

func (s *Service) NoEventRunning(comm *reddit.Comment) {
	_, _, err := s.redditClient.Comment.Submit(context.Background(), comm.FullID, noEventMd)
	if err != nil {
		log.Print(fmt.Errorf("could not submit no event comment: %w", err))
	}
}

func (s *Service) PickemsComment(comm *reddit.Comment) {
	_, _, err := s.redditClient.Comment.Submit(context.Background(), comm.FullID, fmt.Sprintf(joinFmtMd, Event))
	if err != nil {
		log.Print(fmt.Errorf("could not submit pickems comment: %w", err))
	}
}
