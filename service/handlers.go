package service

import (
	"github.com/pborman/uuid"

	"github.com/unrolled/render"
	"net/http"
)

func createMatchHandler(formatter *render.Render, repo matchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		newMatch := gogo.NewMatch(5)
		repo.addMatch(newMatch)
		guid := uuid.New()
		w.Header().Add("Location", "/matches/"+guid)
		formatter.JSON(w, http.StatusCreated, &newMatchResponse{ID: guid})
	}
}
