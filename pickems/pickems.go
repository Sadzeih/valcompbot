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
	if Event == nil {
		s.NoEventRunning(comm)
		return
	}

	rankResp, err := s.getRankFromVLR(comm)
	if err != nil {
		log.Print(err)
		return
	}

	_, _, err := s.redditClient.Comment.Submit(context.Background(), comm.FullID, fmt.Sprintf(pickemsFmtMd, comm.Author, *rankResp.Link, *Event))
	if err != nil {
		log.Print(fmt.Errorf("could not submit pickems comment: %w", err))
	}
}
