package server

import "sync"

// InMemoryPlayerStore collects data about an empty player store
type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.Mutex
}

// NewInMemoryPlayerStore initialises an empty InMemoryPlayerStore
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{store: map[string]int{}}
}

// GetPlayerScore retrieves scores for a given player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

// RecordWin will record a player's win
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetLeague() (league []Player) {
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return
}
