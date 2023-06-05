package pickems

import "text/template"

const (
	leaderboardsFmt = `/r/ValorantCompetitive Pickems Leaderboards

| User | ValComp Rank | Global Rank | Link |
|------|--------------|-------------|------|
{{- range $index, $result := .Pickems }}
| {{ $result.Rank.User.RedditID | printf "/u/%s" }} | $result.Rank.Local | $result.Rank.Absolute | $result.Rank.Link |
{{- end }}`

	rankCommentMd = `You are ranked [#{{ .Global.Absolute }}]({{ .Link }}) globally on [VLR.gg](https://vlr.gg) with a score of {{ .Score }}. Which brings you to the top {{ .Global.Percentile }}% of users!

You are ranked [#{{ .Local.Absolute }}]({{ .Rank.Link }}) (top {{ .Local.Percentile }}%) on the subreddit leaderboard.

[^(Join the subreddit pickems here.)](https://vlr.gg/event/pickem/%d?group_id=valcomp)`

	unrankedCommentMd = `You either do not have a rank yet or have not joined the subreddit group.

[Join the subreddit pickems here!](https://vlr.gg/event/pickem/%d?group_id=valcomp)
`
)

var (
	leaderboardsTmpl = template.Must(template.New("leaderboardsTmpl").Parse(leaderboardsFmt))
	rankTmpl         = template.Must(template.New("rankTmpl").Parse(rankCommentMd))
)
