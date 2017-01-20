package service

import (
	"encoding/json"
	"github.com/pborman/uuid"
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
		guid := uuid.New()
		w.Header().Add("Location", "/matches/"+guid)
		formatter.JSON(w, http.StatusCreated, &newMatchResponse{ID: guid, GridSize: newMatch.GridSize})
	}
}
