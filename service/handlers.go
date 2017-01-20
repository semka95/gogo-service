package service

import (
	"github.com/unrolled/render"
	"net/http"
)

func createMatchHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Location", "some value")
		formatter.JSON(w,
			http.StatusCreated,
			struct{ Test string }{"This is a test"})
	}
}
