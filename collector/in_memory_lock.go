package main

import (
	"sync"
)

type InMemoryLock struct {
	mu    sync.RWMutex
	locks map[int64]struct{}
}

func NewInMemoryLock() *InMemoryLock {
	return &InMemoryLock{
		locks: make(map[int64]struct{}),
	}
}
func (l *InMemoryLock) TryLock(userID int64) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, ok := l.locks[userID]; ok {
		return false
	}
	l.locks[userID] = struct{}{}
	return true
}

func (l *InMemoryLock) Unlock(userID int64) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.locks, userID)
}

func (l *InMemoryLock) IsLocked(userID int64) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()

	_, ok := l.locks[userID]
	return ok
}

func (l *InMemoryLock) GetAllLockedIDs() []int64 {
	l.mu.RLock()
	defer l.mu.RUnlock()

	ids := make([]int64, 0, len(l.locks))
	for id := range l.locks {
		ids = append(ids, id)
	}
	return ids
}
