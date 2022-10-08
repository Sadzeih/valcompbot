package highlighter

import (
	"context"
	"fmt"
	"strings"

	"github.com/Sadzeih/valcompbot/config"
	"github.com/Sadzeih/valcompbot/ent"
	"github.com/Sadzeih/valcompbot/ent/highlightedcomment"
	"github.com/Sadzeih/valcompbot/ent/pinnedcomment"
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

func (h *Highlighter) FindTeam(tag string) (string, error) {
	if teamsResp == nil {
		resp, err := h.sheets.Spreadsheets.Values.Get(spreadsheetID, teamsRange).Do()
		if err != nil {
			return "", err
		}
		teamsResp = resp
	}

	for _, row := range teamsResp.Values {
		if len(row) == 2 && row[1].(string) == tag {
			return row[0].(string), nil
		}
	}

	return strings.Trim(tag, ":"), nil
}

func (h *Highlighter) ParseFlair(flair string) (*Flair, error) {
	verified := false

	// Get fullname + role
	roleName := []byte{}
	for _, submatches := range flairRegex.FindAllSubmatchIndex([]byte(flair), -1) {
		roleName = flairRegex.Expand(roleName, []byte(flairTemplate), []byte(flair), submatches)
	}

	// handle teams
	icons := []byte{}
	for _, submatches := range flairRegex.FindAllSubmatchIndex([]byte(flair), -1) {
		icons = flairRegex.Expand(roleName, []byte("$1"), []byte(flair), submatches)
	}

	teams := make([]string, 0)
	for _, icon := range iconRegex.FindAllString(string(icons), -1) {
		team := ""
		if icon == ":verified:" {
			verified = true
			continue
		}

		// look for Esports Orgs
		team, err := h.FindTeam(string(icon))
		if err != nil {
			return nil, err
		}
		if team == "" {
			// look in here
			switch icon {
			case ":Riot:":
				team = "Riot Games"
				verified = true
			case ":VCT:":
				team = "VCT"
				verified = true
			case ":VAL:":
				team = "VALORANT"
				verified = true
			}
		}
		if team == "" {
			continue
		}
		teams = append(teams, team)
		roleName = []byte(strings.ReplaceAll(string(roleName), fmt.Sprintf(", %s", team), ""))
	}

	// rewriting the flair into a readable string
	rewrittenFlair := ""

	// adding teams
	for _, team := range teams {
		rewrittenFlair += fmt.Sprintf("%s, ", team)
	}

	return &Flair{
		Verified: verified,
		Text:     rewrittenFlair + string(roleName),
		Teams:    teams,
		Type:     h.findType(rewrittenFlair+string(roleName), teams),
	}, nil
}

func (h *Highlighter) findType(flair string, teams []string) string {
	// if user has pro player in flair, it's a pro player
	if strings.Contains(flair, "Pro Player") {
		return "proplayer"
	}

	// if it's not a pro but has a team flair, then it's an org staff
	if len(teams) == 1 {
		return "orgstaff"
	}

	// Look for Riot Games
	for _, team := range teams {
		if team == "Riot Games" {
			return team
		}
	}

	// other prominent members of the community
	return "other"
}

func (h *Highlighter) saveComment(flair *Flair, comment *reddit.Comment) error {
	return h.ent.HighlightedComment.Create().
		SetCommentID(comment.FullID).
		SetBody(comment.Body).
		SetAuthor(comment.Author).
		SetFlair(flair.Text).
		SetAuthorType(flair.Type).
		SetParentID(comment.PostID).
		SetTimestamp(comment.Created.Time).
		SetLink("https://reddit.com" + comment.Permalink).
		Exec(context.Background())
}

func (h *Highlighter) Run() error {
	commentsCh, errsCh, closeFunc := h.reddit.Stream.Comments(config.Get().RedditSubreddit, reddit.StreamDiscardInitial)
	defer closeFunc()

	for {
		select {
		case err := <-errsCh:
			return err
		case comment := <-commentsCh:
			flair, err := h.ParseFlair(comment.AuthorFlairText)
			if err != nil {
				return err
			}
			fmt.Println(flair)
			if !flair.Verified {
				continue
			}

			if err := h.saveComment(flair, comment); err != nil {
				return err
			}

			highlight := ""

			// loop over all user types
			for userType, prefix := range prefixMap {
				// grab all existing comments on this thread
				comments, err := h.ent.HighlightedComment.
					Query().
					Where(
						highlightedcomment.And(
							highlightedcomment.ParentID(comment.PostID),
							highlightedcomment.AuthorType(userType),
						),
					).
					Order(ent.Desc(highlightedcomment.FieldTimestamp)).
					All(context.Background())
				if err != nil {
					return err
				}

				// skip if no comments exist for this type
				if len(comments) > 0 {
					highlight += fmt.Sprintf("%s\n\n", prefix)
					for _, litComment := range comments {
						highlight += fmt.Sprintf("* [Comment by /u/%s -- %s, on %s](%s):\n\n", litComment.Author, litComment.Flair, litComment.Timestamp.Format("Jan 2 2006 03:04PM"), litComment.Link)
						body := litComment.Body
						if len(body) >= 300 {
							body = body[:297] + "..."
						}
						bodyLines := strings.Split(body, "\n")
						for _, line := range bodyLines {
							highlight += fmt.Sprintf("\t> %s\n", line)
						}
						highlight += "\n"
					}
				}
			}

			// find the pinned comment for the thread
			pinned, err := h.ent.PinnedComment.Query().
				Where(pinnedcomment.ParentID(comment.PostID)).
				Only(context.Background())
			if ent.IsNotFound(err) {
				// if we don't find the pinned comment, create it and save it in DB
				newComment, _, err := h.reddit.Comment.Submit(context.Background(), comment.PostID, highlight)
				if err != nil {
					return err
				}
				_, err = h.reddit.Moderation.DistinguishAndSticky(context.Background(), newComment.FullID)
				if err != nil {
					return err
				}
				err = h.ent.PinnedComment.Create().
					SetCommentID(newComment.FullID).
					SetParentID(newComment.ParentID).
					Exec(context.Background())
				if err != nil {
					return err
				}
				continue
			} else if err != nil {
				return err
			}

			// edit the pinned comment
			_, _, err = h.reddit.Comment.Edit(context.Background(), pinned.CommentID, highlight)
			if err != nil {
				return err
			}
		}
	}
}
