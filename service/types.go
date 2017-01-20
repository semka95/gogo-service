package service

type newMatchResponse struct {
	ID        string   `json:"id"`
	StartedAt int64    `json:"started_at"`
	GridSize  int      `json:"gridsize"`
	Players   []player `json:"players"`
}

type player struct {
	Color string `json:"color"`
	Name  string `json:"name"`
}
