package main

import (
	"net/http"

	"github.com/bas-velthuizen/gogo-engine"
	"github.com/pborman/uuid"

	"io/ioutil"

	"encoding/json"

	"github.com/unrolled/render"
)

func createMatchHandler(formatter *render.Render, repo matchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// newMatch := gogo.NewMatch(5, "Black", "White")
		payload, _ := ioutil.ReadAll(req.Body)
		var newMatchRequest newMatchRequest
		json.Unmarshal(payload, &newMatchRequest)

		newMatch := gogo.NewMatch(newMatchRequest.GridSize, "Black", "White")
		repo.addMatch(newMatch)
		guid := uuid.New()
		w.Header().Add("Location", "/matches/"+guid)

		formatter.JSON(w, http.StatusCreated, &newMatchResponse{ID: guid, GridSize: newMatch.GridSize, Players: newMatchRequest.Players})
	}
}
