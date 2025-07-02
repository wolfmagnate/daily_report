package main

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/wolfmagnate/daily_report/collector/infra"
)

type Dispatcher struct {
	userRepo    infra.UserRepository
	lockManager *InMemoryLock
	jobQueue    chan<- Job
	txProvider  infra.TxProvider
}

func NewDispatcher(repo infra.UserRepository, lock *InMemoryLock, queue chan<- Job, txProvider infra.TxProvider) *Dispatcher {
	return &Dispatcher{
		userRepo:    repo,
		lockManager: lock,
		jobQueue:    queue,
		txProvider:  txProvider,
	}
}

func (d *Dispatcher) dispatch() {
	queueCapacity := cap(d.jobQueue)
	currentQueueSize := len(d.jobQueue)
	availableSlots := queueCapacity - currentQueueSize
	jobsToFetch := availableSlots / 2

	if jobsToFetch <= 0 {
		return
	}

	lockedIDs := d.lockManager.GetAllLockedIDs()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := d.txProvider.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}
	defer tx.Rollback(ctx)

	candidates, err := d.userRepo.FindCandidates(ctx, tx, jobsToFetch, lockedIDs)
	if err != nil {
		return
	}

	if len(candidates) == 0 {
		return
	}

	dispatchedCount := 0
	for _, candidate := range candidates {
		if d.lockManager.TryLock(int64(candidate.ID)) {
			job := Job{targetUser: candidate}
			d.jobQueue <- job
			dispatchedCount++
		}
	}
}
