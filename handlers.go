package main

import (
	"net/http"

	"github.com/pborman/uuid"

	"github.com/unrolled/render"
)

func createMatchHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		guid := uuid.New()
		w.Header().Add("Location", "/matches/"+guid)

		formatter.JSON(w, http.StatusCreated, &newMatchResponse{ID: guid})
	}
}
