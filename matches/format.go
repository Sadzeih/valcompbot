package matches

import (
	"bytes"
	"html/template"
	"strings"
)

const (
	MatchMd = `# [{{(index .Teams 0).Name | Trim }}](https://vlr.gg/team/{{(index .Teams 0).ID}}) {{(index .Teams 0).MapsWon}}-{{(index .Teams 1).MapsWon}} [{{(index .Teams 1).Name | Trim }}](https://vlr.gg/team/{{(index .Teams 1).ID}})

[vlr.gg]({{.Info.Link}})


{{ range .Maps }}
{{ if ne (index .Teams 0).IsWinner (index .Teams 1).IsWinner }}
**{{ .Name | Title }}**: {{ (index .Teams 0).RoundsWon }}-{{ (index .Teams 1).RoundsWon }}
{{- end }}
{{- end }}

---

**{{(index .Teams 0).Name | Trim }}** | [VLR](https://vlr.gg/team/{{(index .Teams 0).ID}})

**{{(index .Teams 1).Name | Trim }}** | [VLR](https://vlr.gg/team/{{(index .Teams 1).ID}})

---

{{- range $index, $map := .Maps }}
{{- if ne (index .Teams 0).IsWinner (index .Teams 1).IsWinner }}
# Map {{ $index | AddOne }}: {{ $map.Name | Title }}

{{- if eq (Add (index .Teams 0).RoundsWonAttack (index .Teams 1).RoundsWonDefense) 12 }}
| **Team** | **ATK** | **DEF** | **Total** |
|----------|---------|---------|-----------|
| **{{(index .Teams 0).Name | Trim }}** | {{ (index .Teams 0).RoundsWonAttack }} |  {{ (index .Teams 0).RoundsWonDefense }} | {{ (index .Teams 0).RoundsWon }} |
|| **DEF** | **ATK** ||
| **{{(index .Teams 1).Name | Trim }}** | {{ (index .Teams 1).RoundsWonDefense }} |  {{ (index .Teams 1).RoundsWonAttack }} | {{ (index .Teams 1).RoundsWon }} |
{{- else }}
| **Team** | **DEF** | **ATK** | **Total** |
|----------|---------|---------|-----------|
| **{{(index .Teams 0).Name | Trim }}** | {{ (index .Teams 0).RoundsWonDefense }} |  {{ (index .Teams 0).RoundsWonAttack }} | {{ (index .Teams 0).RoundsWon }} |
|| **ATK** | **DEF** ||
| **{{(index .Teams 1).Name | Trim }}** | {{ (index .Teams 1).RoundsWonAttack }} |  {{ (index .Teams 1).RoundsWonDefense }} | {{ (index .Teams 1).RoundsWon }} |
{{- end }}

#### Map Stats

| **{{(index .Teams 0).Name | Trim }}** | **ACS** | **K** | **D** | **A** |
|------------------|---------|-------|-------|-------|
{{- range $playerIdx, $player := (index .Teams 0).Players }}
| [{{ $player.Alias }}](https://vlr.gg/player/{{ $player.PlayerID }}) **{{ $player.Stats.Agent | Title }}** | {{ $player.Stats.ACS }} | {{ $player.Stats.Kills }} | {{ $player.Stats.Deaths }} | {{ $player.Stats.Assists }}|
{{- end }}
| **{{(index .Teams 1).Name | Trim }}** | **ACS** | **K** | **D** | **A** |
{{- range $playerIdx, $player := (index .Teams 1).Players }}
| [{{ $player.Alias }}](https://vlr.gg/player/{{ $player.PlayerID }}) **{{ $player.Stats.Agent | Title }}** | {{ $player.Stats.ACS }} | {{ $player.Stats.Kills }} | {{ $player.Stats.Deaths }} | {{ $player.Stats.Assists }}|
{{- end }}

[Detailed {{ $map.Name | Title }} Statistics]({{ $map.Link }})
{{- end }}
{{- end }}
`

	titleFmt = "%s vs %s / %s - %s / Post-Match Thread"
)

var (
	funcMap = template.FuncMap{
		"Title": func(s string) string {
			return strings.ToTitle(s)
		},
		"Trim": func(s string) string {
			return strings.Trim(s, "\t ")
		},
		"Add": func(a int, b int) int {
			return a + b
		},
		"AddOne": func(i int) int {
			return i + 1
		},
	}
	pmtTmpl = template.Must(template.New("pmtTmpl").Funcs(funcMap).Parse(MatchMd))
)

func (m *Match) ToMarkdown() (string, error) {
	result := bytes.Buffer{}
	err := pmtTmpl.Execute(&result, m)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
