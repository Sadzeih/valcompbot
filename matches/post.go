package matches

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Sadzeih/valcompbot/config"
	"github.com/Sadzeih/valcompbot/ent"
	"github.com/Sadzeih/valcompbot/ent/trackedevent"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func PostMatch(ctx context.Context, vm *Match, ent *ent.Client, redditClient *reddit.Client) error {
	eID, err := strconv.Atoi(vm.Info.EventID)
	if err != nil {
		return err
	}

	e, err := ent.TrackedEvent.
		Query().
		Where(trackedevent.EventID(eID)).
		Only(ctx)
	if err != nil {
		return err
	}

	title := fmt.Sprintf(titleFmt, vm.Teams[0].Name, vm.Teams[1].Name, e.Name, vm.Info.Series)

	md, err := vm.ToMarkdown()
	if err != nil {
		return err
	}

	_, _, err = redditClient.Post.SubmitText(ctx, reddit.SubmitTextRequest{
		Subreddit: config.Get().RedditSubreddit,
		Title:     title,
		Text:      md,
		// FlairID:   pmtFlairID,
		Spoiler: true,
	})
	if err != nil {
		return err
	}

	return nil
}
