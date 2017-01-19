package main

import (
	"testing"

	"github.com/bas-velthuizen/gogo-engine"
)

func TestAddMatchShowsUpInRepository(t *testing.T) {
	match := gogo.NewMatch(19, "Black", "White")
	repo := newInMemoryRepository()
	err := repo.addMatch(match)
	if err != nil {
		t.Error("Got an error adding a match to repository, should not have.")
	}

	matches := repo.getMatches()
	if len(matches) != 1 {
		t.Errorf("Expected to have 1 match in the repository, got %d", len(matches))
	}

	if matches[0].PlayerBlack != "Black" {
		t.Errorf("Players's name should have been Black, got %s", matches[0].PlayerBlack)
	}

	if matches[0].PlayerWhite != "White" {
		t.Errorf("Players's name should have been White, got %s", matches[0].PlayerWhite)
	}
}

func TestNewRepositoryIsEmpty(t *testing.T) {
	repo := newInMemoryRepository()

	matches := repo.getMatches()
	if len(matches) != 0 {
		t.Errorf("Expected to have 0 matches in newly created repository, got %d", len(matches))
	}
}
