package matches

import (
	"bytes"
	"html/template"
	"strings"
)

var (
	MatchMd = `# [{{(index .Teams 0).Name}}](https://vlr.gg/team/{{(index .Teams 0).ID}}) {{(index .Teams 0).MapsWon}}-{{(index .Teams 1).MapsWon}} [{{(index .Teams 1).Name}}](https://vlr.gg/team/{{(index .Teams 1).ID}})

[vlr.gg]({{.Info.Link}})


{{ range .Maps }}
{{ if ne (index .Teams 0).IsWinner (index .Teams 1).IsWinner }}
**{{ .Name | Title }}**: {{ (index .Teams 0).RoundsWon }}-{{ (index .Teams 1).RoundsWon }}
{{- end }}
{{- end }}

---

**{{(index .Teams 0).Name}}** | [VLR](https://vlr.gg/team/{{(index .Teams 0).ID}})

**{{(index .Teams 1).Name}}** | [VLR](https://vlr.gg/team/{{(index .Teams 1).ID}})

---

[For spoiler-free VALORANT VoDs check out Juked](https://juked.gg/vods?games[]=20&sort=recent)

**Join the subreddit Discord server by clicking the link in the sidebar!**

---

{{- range $index, $map := .Maps }}
{{- if ne (index .Teams 0).IsWinner (index .Teams 1).IsWinner }}
# Map {{ $index | AddOne }}: {{ $map.Name | Title }}
| **Team** | **ATK** | **DEF** | **Total** |
|:----|:-----|:----|---|
| **{{(index .Teams 0).Name}}** | {{ (index .Teams 0).RoundsWonAttack }} |  {{ (index .Teams 0).RoundsWonDefense }} | {{ (index .Teams 0).RoundsWon }}
|              |          **DEF**            |            **ATK**          |
| **{{(index .Teams 1).Name}}** | {{ (index .Teams 1).RoundsWonAttack }} |  {{ (index .Teams 1).RoundsWonDefense }} | {{ (index .Teams 1).RoundsWon }}
---

| **{{(index .Teams 0).Name}}** | **ACS** | **K** | **D** | **A** |
|------------------|---------|-------|-------|-------|
{{- range $playerIdx, $player := (index .Teams 0).Players }}
| [{{ $player.Alias }}](https://vlr.gg/player/{{ $player.PlayerID }}) **{{ $player.Stats.Agent | Title }}** | {{ $player.Stats.ACS }} | {{ $player.Stats.Kills }} | {{ $player.Stats.Deaths }} | {{ $player.Stats.Assists }}|
{{- end }}

| **{{(index .Teams 1).Name}}** | **ACS** | **K** | **D** | **A** |
|------------------|---------|-------|-------|-------|
{{- range $playerIdx, $player := (index .Teams 1).Players }}
| [{{ $player.Alias }}](https://vlr.gg/player/{{ $player.PlayerID }}) **{{ $player.Stats.Agent | Title }}** | {{ $player.Stats.ACS }} | {{ $player.Stats.Kills }} | {{ $player.Stats.Deaths }} | {{ $player.Stats.Assists }}|
{{- end }}

[Detailed {{ $map.Name | Title }} Statistics]({{ $map.Link }})
{{- end }}
{{- end }}
`

	titleFmt = "%s vs %s / %s - %s / Post-Match-Thread"
)

func (m *Match) ToMarkdown() (string, error) {
	funcMap := template.FuncMap{
		"Title": strings.Title,
		"AddOne": func(i int) int {
			return i + 1
		},
	}

	tmpl, err := template.New("toMarkdown").Funcs(funcMap).Parse(MatchMd)
	if err != nil {
		return "", err
	}
	result := bytes.Buffer{}
	err = tmpl.Execute(&result, m)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
