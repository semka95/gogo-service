package service

import (
	"github.com/pborman/uuid"

	"github.com/unrolled/render"
	"net/http"
)

func createMatchHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		guid := uuid.New()
		w.Header().Add("Location", "/matches/"+guid)
		formatter.JSON(w,
			http.StatusCreated,
			struct{ Test string }{"This is a test"})
	}
}
