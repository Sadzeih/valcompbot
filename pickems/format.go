package pickems

import "text/template"

const (
	leaderboardsFmt = `/r/ValorantCompetitive Pickems Leaderboards

| User | ValComp Rank | Global Rank | Link |
|------|--------------|-------------|------|
{{- range $index, $result := .Pickems }}
| {{ printf "/u/" }}{{ $result.User.RedditID }} | {{ $result.Rank.Local.Absolute }} | {{ $result.Rank.Global.Absolute }} | [link]({{ $result.Link }}) |
{{- end }}`

	rankCommentMd = `You are ranked [#{{ .Rank.Global.Absolute }}]({{ .Link }}) globally on [VLR.gg](https://vlr.gg) with a score of {{ .Score }}. Which brings you to the top {{ .Rank.Global.Percentile }}%% of users!

You are ranked [#{{ .Rank.Local.Absolute }}]({{ .Link }}) (top {{ .Rank.Local.Percentile }}%%) on the subreddit leaderboard.

[^(Join the subreddit pickems here.)](https://vlr.gg/event/pickem/%d?group=valcomp)`

	unrankedCommentMd = `You either do not have a rank yet or have not joined the subreddit group.

[Join the subreddit pickems here!](https://vlr.gg/event/pickem/%d?group=valcomp)
`
	pickemsFmtMd = `[Click here to see /u/%s's pickems](%s)
	
[Join the subreddit pickems here.](https://vlr.gg/event/pickem/%d?group=valcomp)`
)

var (
	leaderboardsTmpl = template.Must(template.New("leaderboardsTmpl").Parse(leaderboardsFmt))
	rankTmpl         = template.Must(template.New("rankTmpl").Parse(rankCommentMd))
)
