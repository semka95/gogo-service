package service

import (
	"encoding/json"
	"io/ioutil"

	"github.com/cloudnativego/gogo-engine"
	"github.com/unrolled/render"
	"net/http"
)

func createMatchHandler(formatter *render.Render, repo matchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		var newMatchRequest newMatchRequest
		err := json.Unmarshal(payload, &newMatchRequest)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse create match request")
			return
		}
		if !newMatchRequest.isValid() {
			formatter.Text(w, http.StatusBadRequest, "Invalid new match request")
			return
		}

		newMatch := gogo.NewMatch(newMatchRequest.GridSize, newMatchRequest.PlayerBlack, newMatchRequest.PlayerWhite)
		repo.addMatch(newMatch)
		w.Header().Add("Location", "/matches/"+newMatch.ID)
		formatter.JSON(w, http.StatusCreated, &newMatchResponse{ID: newMatch.ID, GridSize: newMatch.GridSize, PlayerBlack: newMatchRequest.PlayerBlack, PlayerWhite: newMatchRequest.PlayerWhite})
	}
}
