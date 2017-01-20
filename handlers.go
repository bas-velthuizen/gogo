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
		err := json.Unmarshal(payload, &newMatchRequest)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse create match request.")
			return
		}

		if !newMatchRequest.isValid() {
			formatter.Text(w, http.StatusBadRequest, "Invalid new match request.")
			return
		}

		newMatch := gogo.NewMatch(newMatchRequest.GridSize, "Black", "White")
		repo.addMatch(newMatch)
		guid := uuid.New()
		w.Header().Add("Location", "/matches/"+guid)

		formatter.JSON(w, http.StatusCreated, &newMatchResponse{ID: guid, GridSize: newMatch.GridSize, PlayerBlack: newMatchRequest.PlayerBlack, PlayerWhite: newMatchRequest.PlayerWhite})
	}
}
