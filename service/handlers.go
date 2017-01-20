package service

import (
	"encoding/json"
	"io/ioutil"

	"github.com/unrolled/render"
	"net/http"
)

func createMatchHandler(formatter *render.Render, repo matchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		var newMatchRequest newMatchRequest
		json.Unmarshal(payload, &newMatchRequest)

		newMatch := gogo.NewMatch(newMatchRequest.GridSize)
		repo.addMatch(newMatch)
		w.Header().Add("Location", "/matches/"+newMatch.ID)
		formatter.JSON(w, http.StatusCreated, &newMatchResponse{ID: newMatch.ID, GridSize: newMatch.GridSize})
	}
}
