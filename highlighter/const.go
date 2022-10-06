package highlighter

import (
	"regexp"

	"google.golang.org/api/sheets/v4"
)

const (
	spreadsheetID = "1fqvgfdNYYOQylt4_mTA3MjaoysnDl-XPPnRLoTDGFKo"
	teamsRange    = "DATA!C57:D"
)

type Flair struct {
	Verified bool
	Teams    []string
	Text     string
	Type     string
}

var (
	teamsResp *sheets.ValueRange = nil

	flairRegex = regexp.MustCompile(`((:[\w-]+:)+) (.*) - (.*)((:[\w-]+:)*)`)
	iconRegex  = regexp.MustCompile(`:[\w-]+:`)

	prefixMap = map[string]string{
		"Riot Games": "Comments made by Riot Games employees:",
		"proplayer":  "Comments made by Pro Players:",
		"orgstaff":   "Comments made by Org Staff:",
		"other":      "Comments made by prominent members of the community:",
	}
)

const (
	flairTemplate = "$3 - $4"
)

// TODO: Might be useful to use this to automatically update flairs on reddit through the API
//
// type UserKinds struct {
// 	Type        string
// 	Prefix      string
// 	UserMap     userMap
// 	Ranges      Ranges
// 	SkipRetired bool
// }

// type Ranges struct {
// 	Sheet   string
// 	NameRow int
// 	RoleRow int
// }

//	userKinds = []UserKinds{
//		{
//			Type:    "rioter",
//			Prefix:  "Comments made by Riot Games employees:",
//			UserMap: make(userMap),
//			Ranges: Ranges{
//				Sheet:   "Riot Games!A3:I",
//				NameRow: 3,
//				RoleRow: 8,
//			},
//			SkipRetired: true,
//		},
//		{
//			Type:    "proplayer",
//			Prefix:  "Comments made by Pro Players:",
//			UserMap: make(userMap),
//			Ranges: Ranges{
//				Sheet:   "Pro Players!A3:L",
//				NameRow: 5,
//				RoleRow: 11,
//			},
//			SkipRetired: false,
//		},
//		{
//			Type:    "broadcasttalent",
//			Prefix:  "Comments made by Broadcast Talent:",
//			UserMap: make(userMap),
//			Ranges: Ranges{
//				Sheet:   "Broadcast Talent!A3:L",
//				NameRow: 5,
//				RoleRow: 11,
//			},
//			SkipRetired: false,
//		},
//		{
//			Type:    "orgstaff",
//			Prefix:  "Comments made by Org Staff:",
//			UserMap: make(userMap),
//			Ranges: Ranges{
//				Sheet:   "Pro Org Staff!A4:L",
//				NameRow: 5,
//				RoleRow: 11,
//			},
//			SkipRetired: false,
//		},
//		{
//			Type:    "officialaccounts",
//			Prefix:  "Comments made by official Org accounts:",
//			UserMap: make(userMap),
//			Ranges: Ranges{
//				Sheet:   "Esports Teams!A3:L",
//				NameRow: 5,
//				RoleRow: -1,
//			},
//			SkipRetired: false,
//		},
//	}
