package service

import "github.com/cloudnativego/gogo-engine"

type newMatchResponse struct {
	ID          string `json:"id"`
	StartedAt   int64  `json:"started_at"`
	GridSize    int    `json:"gridsize"`
	PlayerWhite string `json:"playerWhite"`
	PlayerBlack string `json:"playerBlack"`
}

type player struct {
	Color string `json:"color"`
	Name  string `json:"name"`
}

type newMatchRequest struct {
	GridSize    int    `json:"gridsize"`
	PlayerWhite string `json:"playerWhite"`
	PlayerBlack string `json:"playerBlack"`
}

type matchRepository interface {
	addMatch(match gogo.Match) (err error)
	getMatches() []gogo.Match
}
