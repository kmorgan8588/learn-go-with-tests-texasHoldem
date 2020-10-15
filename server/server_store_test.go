package server

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func createTempFile(t *testing.T, inititalData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Errorf("could not create temp file, %v", err)
	}

	tmpfile.Write([]byte(inititalData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func TestFileSystemStore(t *testing.T) {

	t.Run("/league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Kyle", "Wins": 22}]`)
		defer cleanDatabase()
		store := NewFileSystemPlayerStore(database)

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
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Kyle", "Wins": 22}]`)
		defer cleanDatabase()
		store := NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Kyle")

		want := 22
		assertScoreEquals(t, got, want)

	})

	t.Run("increase player wins", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Kyle", "Wins": 22}]`)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Kyle")

		got := store.GetPlayerScore("Kyle")
		want := 23
		assertScoreEquals(t, got, want)
	})

	t.Run("add wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Kyle", "Wins": 22}]`)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Tim")
		got := store.GetPlayerScore("Tim")
		want := 1
		assertScoreEquals(t, got, want)
	})
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
