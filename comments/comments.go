package comments

import (
	"github.com/Sadzeih/valcompbot/config"
	"github.com/google/uuid"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type Topic struct {
	commentsChan <-chan *reddit.Comment
	errsChan     <-chan error

	closeFunc func()

	subscribers map[string]*Subscriber
}

type Subscriber struct {
	id       string
	closed   bool
	Comments chan *reddit.Comment
	Errors   chan error
}

func New(client *reddit.Client) *Topic {
	commentsCh, errsCh, closeFunc := client.Stream.Comments(config.Get().RedditSubreddit, reddit.StreamDiscardInitial)

	return &Topic{
		commentsChan: commentsCh,
		errsChan:     errsCh,
		closeFunc:    closeFunc,
		subscribers:  map[string]*Subscriber{},
	}
}

func (t *Topic) Subscribe() *Subscriber {
	id := uuid.NewString()
	commentsChan := make(chan *reddit.Comment)
	errsChan := make(chan error)
	t.subscribers[id] = &Subscriber{
		id:       id,
		Comments: commentsChan,
		Errors:   errsChan,
	}

	return t.subscribers[id]
}

func (t *Topic) Run() {
	for {
		select {
		case err := <-t.errsChan:
			if err != nil {
				for _, sub := range t.subscribers {
					if sub.closed {
						close(sub.Comments)
						close(sub.Errors)
						delete(t.subscribers, sub.id)
						break
					}
					sub.Errors <- err
				}
			}
		case comm := <-t.commentsChan:
			for _, sub := range t.subscribers {
				if sub.closed {
					close(sub.Comments)
					close(sub.Errors)
					delete(t.subscribers, sub.id)
					break
				}
				sub.Comments <- comm
			}
		}
	}
}

func (t *Topic) Close() {
	for _, sub := range t.subscribers {
		close(sub.Comments)
		close(sub.Errors)
		delete(t.subscribers, sub.id)
	}
}

func (s *Subscriber) Close() {
	s.closed = true
}
