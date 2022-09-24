package main

import (
	"log"
	"net/http"
)

// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct{}

// GetPlayerScore retrieves scores for a given player.
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

// GetPlayerScore retrieves scores for a given player.
func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5001", server))
}

// Integration tests section
// TODO https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server#integration-tests
