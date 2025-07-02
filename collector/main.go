package main

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/wolfmagnate/daily_report/collector/infra"
	"github.com/wolfmagnate/daily_report/collector/integration"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using environment variables")
	}

	dispatcher, workerPool, err := di()
	if err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
	}

	callDispatcher(dispatcher)
	workerPool.Run()
}

func di() (*Dispatcher, *WorkerPool, error) {
	pool, err := infra.NewPgxPool()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	txProvider := &infra.PgxTxProvider{Pool: pool}
	userRepo := infra.NewUserRepository()
	githubActivityRepo := infra.NewGithubActivityRepository()
	slackActivityRepo := infra.NewSlackActivityRepository()

	githubIntegrator := integration.NewGithubIntegrator()
	slackIntegrator := integration.NewSlackIntegrator()

	jobQueue := make(chan Job, 100)
	lockManager := NewInMemoryLock()

	dispatcher := NewDispatcher(userRepo, lockManager, jobQueue, txProvider)
	workerPool := NewWorkerPool(
		10,
		cap(jobQueue),
		jobQueue,
		lockManager,
		txProvider,
		githubIntegrator,
		slackIntegrator,
		githubActivityRepo,
		slackActivityRepo,
	)
	return dispatcher, workerPool, nil
}

func callDispatcher(dispatcher *Dispatcher) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		go dispatcher.dispatch()
		fmt.Println("Calling dispatcher...")
	}
}
