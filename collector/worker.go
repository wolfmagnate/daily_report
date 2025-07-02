package main

import (
	"context"
	"log"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/wolfmagnate/daily_report/collector/domain"
	"github.com/wolfmagnate/daily_report/collector/infra"
	"github.com/wolfmagnate/daily_report/collector/integration"
)

type Worker struct {
	id          int
	jobQueue    <-chan Job
	lockManager *InMemoryLock
	wg          *sync.WaitGroup
	txProvider  infra.TxProvider

	githubIntegrator integration.GithubIntegrator
	slackIntegrator  integration.SlackIntegrator

	githubActivityRepo infra.GithubActivityRepository
	slackActivityRepo  infra.SlackActivityRepository
}

type Job struct {
	targetUser *domain.User
}

func NewWorker(
	id int,
	jobQueue <-chan Job,
	lockManager *InMemoryLock,
	wg *sync.WaitGroup,
	txProvider infra.TxProvider,
	githubIntegrator integration.GithubIntegrator,
	slackIntegrator integration.SlackIntegrator,
	githubActivityRepo infra.GithubActivityRepository,
	slackActivityRepo infra.SlackActivityRepository,
) *Worker {
	return &Worker{
		id:               id,
		jobQueue:         jobQueue,
		lockManager:      lockManager,
		wg:               wg,
		txProvider:       txProvider,
		githubIntegrator: githubIntegrator,
		slackIntegrator:  slackIntegrator,
		githubActivityRepo: githubActivityRepo,
		slackActivityRepo: slackActivityRepo,
	}
}

func (w *Worker) start() {
	w.wg.Add(1)
	defer w.wg.Done()

	for job := range w.jobQueue {
		w.processJob(job)
	}
}

func (w *Worker) processJob(job Job) {
	userID := int64(job.targetUser.ID)
	if !w.lockManager.TryLock(userID) {
		log.Printf("User %d is already being processed. Skipping.", userID)
		return
	}
	defer w.lockManager.Unlock(userID)

	ctx := context.Background()
	tx, err := w.txProvider.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		log.Printf("Error starting transaction for user %d: %v", userID, err)
		return
	}
	defer tx.Rollback(ctx)

	githubActivities, isTokenUpdated, newToken, latestEventID, err := w.githubIntegrator.GetUserEvents(ctx, job.targetUser.GithubIntegration, job.targetUser.ID)
	if err != nil {
		log.Printf("Error getting GitHub events for user %d: %v", userID, err)
		return
	}

	if isTokenUpdated {
		err = w.githubActivityRepo.UpdateGithubIntegrationToken(ctx, tx, job.targetUser.ID, newToken)
		if err != nil {
			log.Printf("Error updating GitHub integration token for user %d: %v", userID, err)
			return
		}
	}

	err = w.githubActivityRepo.UpdateLastProcessedEventID(ctx, tx, job.targetUser.ID, latestEventID)
	if err != nil {
		log.Printf("Error updating last processed event ID for user %d: %v", userID, err)
		return
	}

	_, err = w.githubActivityRepo.CreateGithubActivities(ctx, tx, githubActivities)
	if err != nil {
		log.Printf("Error creating GitHub activities for user %d: %v", userID, err)
		return
	}

	slackActivities, isTokenUpdated, newToken, err := w.slackIntegrator.GetUserMessages(ctx, job.targetUser.SlackIntegration, job.targetUser.ID, job.targetUser.LastCollectedAt)
	if err != nil {
		log.Printf("Error getting Slack messages for user %d: %v", userID, err)
		return
	}

	if isTokenUpdated {
		err = w.slackActivityRepo.UpdateIntegrationToken(ctx, tx, job.targetUser.ID, newToken)
		if err != nil {
			log.Printf("Error updating Slack integration token for user %d: %v", userID, err)
			return
		}
	}

	_, err = w.slackActivityRepo.CreateSlackActivities(ctx, tx, slackActivities)
	if err != nil {
		log.Printf("Error creating Slack activities for user %d: %v", userID, err)
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Printf("Error committing transaction for user %d: %v", userID, err)
		return
	}

	log.Printf("Successfully processed user %d", userID)
}

