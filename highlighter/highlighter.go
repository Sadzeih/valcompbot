package highlighter

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/Sadzeih/valcompbot/config"
	"github.com/Sadzeih/valcompbot/ent"
	"github.com/Sadzeih/valcompbot/ent/highlightedcomment"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"google.golang.org/api/sheets/v4"
)

type Highlighter struct {
	ent    *ent.Client
	reddit *reddit.Client
	sheets *sheets.Service
}

func New(ctx context.Context, r *reddit.Client, ent *ent.Client) (*Highlighter, error) {
	sheets, err := sheets.NewService(ctx)
	if err != nil {
		return nil, err
	}
	return &Highlighter{
		reddit: r,
		sheets: sheets,
		ent:    ent,
	}, nil
}

func (h *Highlighter) GetUsers() error {
	resp, err := h.sheets.Spreadsheets.Values.Get(spreadsheetID, rioterRange).Do()
	if err != nil {
		return err
	}

	for _, row := range resp.Values {
		// Riot employees logic
		if strings.Contains(row[0].(string), "Retired") {
			// skip retired rioters
			break
		} else if len(row) <= 1 || len(row[3].(string)) <= 8 {
			// skip rows without an actual user
			continue
		}
		rioters.Add(FlairedUser{Username: row[3].(string), Role: row[8].(string)})
	}

	return nil
}

func (h *Highlighter) Run() error {
	commentsCh, errsCh, closeFunc := h.reddit.Stream.Comments(config.Get().RedditSubreddit, reddit.StreamDiscardInitial)
	defer closeFunc()

	for {
		select {
		case err := <-errsCh:
			return err
		case comment := <-commentsCh:
			edited := false
			if rioters.Exists(comment.Author) {
				edited = true
				h.ent.HighlightedComment.Create().
					SetCommentID(comment.FullID).
					SetBody(comment.Body).
					SetAuthor(comment.Author).
					SetAuthorRole(rioters[comment.Author].Role).
					SetAuthorType("rioter").
					SetParentID(comment.ParentID).
					SetLink("https://reddit.com" + comment.Permalink)
			}

			if !edited {
				continue
			}
			highlight := ""
			// grab all existing rioter comments on this thread
			rioterComments, err := h.ent.HighlightedComment.
				Query().
				Where(highlightedcomment.ParentID(comment.ParentID), highlightedcomment.AuthorType("rioter")).
				Order(ent.Asc(highlightedcomment.FieldAuthor)).
				All(context.Background())
			if err != nil {
				return err
			}
			highlight += fmt.Sprintf("%s\n\n", prefix["rioter"])
			for _, rioterComment := range rioterComments {
				highlight += fmt.Sprintf("* [Comment by /u/%s - %s](%s):\n", rioterComment.Author, rioterComment.AuthorRole, rioterComment.Link)
				body := rioterComment.Body[:(int)(math.Min(296, (float64)(len(rioterComment.Body)-1)))] + "..."
				bodyLines := strings.Split(body, "\n")
				for i, line := range bodyLines {
					highlight += fmt.Sprintf("> %s\n", line)
					if i != len(bodyLines)-1 {
						highlight += "\n>\n"
					} else {
						highlight += "\n"
					}
				}
			}
			_, _, err = h.reddit.Comment.Submit(context.Background(), comment.ParentID, highlight)
			if err != nil {
				return err
			}
		}
	}
}
