package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryLock_TryLock(t *testing.T) {
	lock := NewInMemoryLock()

	// Test locking an unlocked ID
	t.Run("Lock an unlocked ID", func(t *testing.T) {
		success := lock.TryLock(1)
		assert.True(t, success, "Should be able to lock an unlocked ID")
	})

	// Test locking a locked ID
	t.Run("Lock a locked ID", func(t *testing.T) {
		lock.TryLock(2) // Pre-lock an ID
		success := lock.TryLock(2)
		assert.False(t, success, "Should not be able to lock an already locked ID")
	})
}

func TestInMemoryLock_Unlock(t *testing.T) {
	lock := NewInMemoryLock()

	lock.TryLock(1) // Lock an ID
	lock.Unlock(1)  // Unlock it

	success := lock.TryLock(1) // Try to lock it again
	assert.True(t, success, "Should be able to re-lock an unlocked ID")
}

func TestInMemoryLock_IsLocked(t *testing.T) {
	lock := NewInMemoryLock()

	// Test an unlocked ID
	assert.False(t, lock.IsLocked(1), "Should report an unlocked ID as not locked")

	// Test a locked ID
	lock.TryLock(2)
	assert.True(t, lock.IsLocked(2), "Should report a locked ID as locked")
}

func TestInMemoryLock_GetAllLockedIDs(t *testing.T) {
	lock := NewInMemoryLock()

	lock.TryLock(1)
	lock.TryLock(10)
	lock.TryLock(100)

	lockedIDs := lock.GetAllLockedIDs()
	assert.ElementsMatch(t, []int64{1, 10, 100}, lockedIDs, "Should return all locked IDs")
}

func TestInMemoryLock_Concurrency(t *testing.T) {
	lock := NewInMemoryLock()
	var wg sync.WaitGroup
	const numGoroutines = 100
	const userID = 1

	// This channel will count how many goroutines successfully locked the ID.
	successfulLocks := make(chan bool, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if lock.TryLock(userID) {
				successfulLocks <- true
				// Keep the lock for a short time to increase contention
				// time.Sleep(1 * time.Millisecond)
				lock.Unlock(userID)
			}
		}()
	}

	wg.Wait()
	close(successfulLocks)

	// In a concurrent scenario, it's hard to guarantee a specific number of locks,
	// but we can check that at least one succeeded and that the final state is unlocked.
	assert.GreaterOrEqual(t, len(successfulLocks), 1, "At least one goroutine should have acquired the lock")

	// After all operations, the lock should be released.
	assert.False(t, lock.IsLocked(userID), "Lock should be released after all goroutines are done")
}
