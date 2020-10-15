package server

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	database := strings.NewReader(`[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Kyle", "Wins": 22}]`)

	t.Run("/league from a reader", func(t *testing.T) {

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Kyle", 22},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Kyle")

		want := 22
		assertScore(t, got, want)
	})
}

func assertScore(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
