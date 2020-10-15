package server

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) (wins int) {
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}
	return
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type ReadSeeker interface {
	Reader
	Seeker
}

type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}
