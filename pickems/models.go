package pickems

type RankResponse struct {
	Rank *struct {
		Global *struct {
			Absolute   *string `json:"absolute,omitempty"`
			Percentile *string `json:"percentile,omitempty"`
		} `json:"global,omitempty"`
		Local *struct {
			Absolute   *string `json:"absolute,omitempty"`
			Percentile *string `json:"percentile,omitempty"`
		} `json:"local,omitempty"`
	} `json:"rank"`
	Score *string `json:"score,omitempty"`
	Link  *string `json:"link,omitempty"`
}

type LeaderboardsResponse struct {
	Pickems []struct {
		RankResponse
		User *struct {
			RedditID *string `json:"reddit_id,omitempty" json:"reddit_id,omitempty"`
		} `json:"user,omitempty"`
	}
}
